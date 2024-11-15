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

	// if flag.NArg() < 1 {
	// 	flag.Usage()
	// 	return
	// }

	var post_process func(qparty.SeasonIndex, string) error
	switch *output_format {
	case "json":
		post_process = WriteSeasonIndexJSON
		// write_episode = ...json
		// write_metadata = ...json
	case "sqlite":
		post_process = output_sqlite
		// write_episode = ...sqlite
		// write_metadata = ...sqlite
	default:
		log.Print("unrecognized output format ", *output_format)
		flag.Usage()
		return
	}

	// Read season index (season.json)
	jarchive := qparty.SeasonIndex{
		Version:  []uint{1, 0},
		Episodes: make(map[qparty.EpisodeID]qparty.EpisodeMetadata, 10_000),
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

		log.Println()
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

			//log.Print("episode id ", jeid)
			episode := html.ParseEpisode(reader)
			episode.EpisodeID = jeid
			episode.SeasonID = jsid
			if !episode_dates.Aired.IsZero() {
				episode.Aired = episode_dates.Aired
			}
			if !episode_dates.Taped.IsZero() {
				episode.Taped = episode_dates.Aired
			}

			// Also write the converted episode details to .json format
			// (these will be read again if the output format is sqlite).
			err = WriteEpisodeJSON(episode,
				path.Join(*data_path, "json", episode.ShowNumber.JSON()))
			if err != nil {
				log.Print("ERROR writing episode: ", err)
			}

			// Convert contestants to contestant IDs before storing in metadata
			episode.ContestantIDs = make([]qparty.ContestantID, len(episode.Contestants))
			for i, contestant := range episode.Contestants {
				episode.ContestantIDs[i].UCID = contestant.UCID
			}
			episode.Contestants = []qparty.Contestant{}

			// Index the episode metadata by its EpisodeID as that is all the client
			// has when listing a season index, before fetching the episode details.
			jarchive.Episodes[episode.EpisodeID] = episode.EpisodeMetadata

		}
	}

	post_process(jarchive, *data_path)
}

func WriteEpisodeJSON(episode *qparty.FullEpisode, filepath string) error {
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

//
//
//
//

type SeasonEpisodes struct {
	qparty.SeasonID `json:"season"`
	Name            string                            `json:"name"`
	Count           int                               `json:"count"`
	Episodes        map[qparty.EpisodeID]EpisodeDates `json:"episodes"`
}

type EpisodeDates struct {
	qparty.EpisodeID `json:"-"`
	Aired            qparty.ShowDate `json:"aired"`
	Taped            qparty.ShowDate `json:"taped"`
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
