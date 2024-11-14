// Copyright (c) 2024 Kevin Damm
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// github:kevindamm/q-party/cmd/jarchive/main.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	qparty "github.com/kevindamm/q-party"
	"github.com/kevindamm/q-party/cmd/jarchive/html-parser"
)

func main() {
	data_path := flag.String("data", ".data",
		"path where converted and created games are written")
	output_format := flag.String("out", "json",
		"(json|sqlite) encoding of the converted season or episode representation")

	flag.Usage = func() {
		fmt.Printf("%s command id# [flags]\n", path.Base(os.Args[0]))
		fmt.Println("  where")
		fmt.Println("    command is either 'season' or 'episode'")
		fmt.Println("    id# is the unique ID for the season or episode")
		fmt.Println()
		fmt.Println("Fetches the season or episode if absent, then converts it.")
		fmt.Println()
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	var post_process func(qparty.SeasonIndex, string) error
	switch *output_format {
	case "json":
		post_process = WriteSeasonIndexJSON
	case "sqlite":
		post_process = output_sqlite
	default:
		log.Print("unrecognized output format ", *output_format)
		flag.Usage()
		return
	}

	// Read season index (season.json)
	jarchive := qparty.SeasonIndex{
		Version:  []uint{1, 0},
		Episodes: make(map[qparty.ShowNumber]qparty.EpisodeMetadata, 10_000),
	}
	seasons := html.MustLoadAllSeasons(*data_path)

	log.Print("loaded ", len(seasons.Seasons), " seasons")

	// Read per-season episode list (seasons/[seasonid].json)
	jarchive.Seasons = make(map[qparty.SeasonID]qparty.SeasonMetadata, len(seasons.Seasons)+1)
	for jsid, season := range seasons.Seasons {
		key := qparty.SeasonID(jsid)
		number := key.RegularSeason()
		if number > 0 {
			key = qparty.SeasonID(fmt.Sprintf("%d", number))
		}
		season.SeasonID = qparty.SeasonID(season.Season)
		season.Season = ""
		season.Title = season.Name
		season.Name = ""
		season.EpisodesCount = season.Count
		season.Count = 0
		jarchive.Seasons[key] = season
	}

	// Read per-season episode list (seasons/[seasonid].json)
	for jsid, season := range jarchive.Seasons {
		var season_episodes SeasonEpisodes
		season_json := path.Join(*data_path, "seasons", jsid.JSON())
		bytes, err := os.ReadFile(season_json)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(bytes, &season_episodes)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Converting episodes from season [", jsid, "],", season.Name)
		log.Print("~ - ~ - ~ - ~ - ~ - ~ - ~ - ~ - ~\n\n")

		for jeid, episode_dates := range season_episodes.Episodes {
			filepath := path.Join(*data_path, "episodes", jeid.HTML())
			if _, err := os.Stat(filepath); os.IsNotExist(err) {
				log.Print("HTML not found (fetch?) ", jeid.HTML())
				continue
			}

			// Parse the episode's HTML to get the show and challenge details.
			reader, err := os.Open(filepath)
			if err != nil {
				// Unlikely, we know the file exists at this point, but state changes...
				log.Print("ERROR: ", err)
				continue
			}
			defer reader.Close()

			var metadata qparty.EpisodeMetadata
			episode := html.ParseEpisode(jeid, reader)

			// Convert the episode metadata to qparty.* type and include added fields.
			metadata = episode.EpisodeMetadata.EpisodeMetadata
			if !episode_dates.Aired.IsZero() {
				metadata.Aired = episode_dates.Aired
			}
			if !episode_dates.Taped.IsZero() {
				metadata.Taped = episode_dates.Aired
			}
			metadata.SeasonID = jsid
			for i := range 3 {
				metadata.ContestantIDs[i] = episode.Contestants[i].ContestantID
			}
			jarchive.Episodes[episode.ShowNumber] = metadata

			qpepisode := qparty.Episode{
				EpisodeMetadata: metadata,
				Comments:        episode.Comments,
				Media:           episode.Media,
				Single:          convert_board(episode.Single),
				Double:          convert_board(episode.Double),
				Final:           &episode.Final.Challenge,
				TieBreaker:      &episode.TieBreaker.Challenge,
			}

			// Also write the converted episode details to .json format
			// (these will be read again if the output format is sqlite).
			err = WriteEpisodeJSON(qpepisode,
				path.Join(*data_path, "episodes", episode.ShowNumber.JSON()))
			if err != nil {
				log.Print("ERROR writing episode: ", err)
			}
		}
	}

	post_process(jarchive, *data_path)
}

func WriteEpisodeJSON(episode qparty.Episode, filepath string) error {
	writer, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file %s\n%s", filepath, err)
	}
	defer writer.Close()

	bytes, err := json.MarshalIndent(episode, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal index into json\n%s", err)
	}

	_, err = writer.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write bytes\n%s", err)
	}
	return nil
}

func convert_board(from [6]html.CategoryChallenges) *qparty.HostBoard {
	board := new(qparty.HostBoard)

	board.Columns = make([]qparty.HostCategory, 6)
	for i, category := range from {
		board.Columns[i].Title = string(category.JArchiveCategory)
		board.Columns[i].Comments = category.Commentary
		board.Columns[i].Challenges = convert_challenges(category.Challenges)
	}

	return board
}

func convert_challenges(from []html.JArchiveChallenge) []qparty.HostChallenge {
	challenges := make([]qparty.HostChallenge, len(from))
	for i, challenge := range from {
		challenges[i].Challenge = challenge.Challenge
		challenges[i].Correct = challenge.Correct
	}
	return challenges
}

//
//
//
//

type SeasonEpisodes struct {
	qparty.SeasonID `json:"season"`
	Name            string                     `json:"name"`
	Count           int                        `json:"count"`
	Episodes        map[html.JEID]EpisodeDates `json:"episodes"`
}

type EpisodeDates struct {
	html.JEID `json:"-"`
	Aired     qparty.ShowDate `json:"aired"`
	Taped     qparty.ShowDate `json:"taped"`
}

//
//
//
//

func WriteSeasonIndexJSON(jarchive qparty.SeasonIndex, data_path string) error {
	filepath := path.Join(data_path, "jarchive.json")
	writer, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file %s\n%s", filepath, err)
	}
	defer writer.Close()

	bytes, err := json.MarshalIndent(jarchive, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal index into json\n%s", err)
	}

	_, err = writer.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write bytes\n%s", err)
	}
	return nil
}

func output_sqlite(jarchive qparty.SeasonIndex, data_path string) error {
	log.Fatal("TODO write to database")
	return nil
}
