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

	var write_episode func(qparty.FullEpisode, string) error
	var write_metadata func(qparty.SeasonIndex, string) error

	switch *output_format {
	case "":
		write_episode = NoOpEpisode
		write_metadata = NoOpMetadata
	case "json":
		must_create_dir(*data_path, "json")
		write_episode = WriteEpisodeJSON
		write_metadata = WriteSeasonIndexJSON
	case "sqlite":
		//post_process = output_sqlite
		// confirm database exists
		// write_episode = ...sqlite
		// write_metadata = ...sqlite
		write_episode = NoOpEpisode
		write_metadata = NoOpMetadata
	default:
		log.Print("unrecognized output format ", *output_format)
		flag.Usage()
		return
	}

	// Read season index (season.json)
	jarchive := qparty.SeasonIndex{
		Version:  []uint{1, 0},
		Episodes: make(map[qparty.EpisodeID]qparty.EpisodeStats, 10_000),
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
			episode.ShowNumber = qparty.ShowNumber(
				season.Prefix() + string(episode.ShowNumber))

			err = write_episode(*episode,
				path.Join(*data_path, "json", episode.ShowNumber.JSON()))
			if err != nil {
				log.Print("ERROR writing episode: ", err)
			}

			episode_stats := qparty.EpisodeStats{
				EpisodeMetadata: episode.EpisodeMetadata,
			}

			// Convert episode contestants to contestant IDs for metadata.
			episode_stats.ContestantIDs = make([]qparty.ContestantID, len(episode.Contestants))
			for i, contestant := range episode.Contestants {
				episode_stats.ContestantIDs[i].UCID = contestant.UCID
			}
			episode_stats.Stumpers = make([][]qparty.Position, 2)

			// Roll up stats (# challenges, # stumpers) for metadata.
			// TODO
			if episode.Single != nil && episode.Single.Columns != nil {
				for i, column := range episode.Single.Columns {
					for j, challenge := range column.Challenges {
						if challenge.ChallengeID != 0 {
							episode_stats.SingleCount += 1
							if challenge.TripleStumper {
								episode_stats.Stumpers[0] = append(episode_stats.Stumpers[0], qparty.Position{
									Column: uint(i),
									Index:  uint(j)})
							}
						}
					}
				}
			}
			if episode.Double != nil && episode.Double.Columns != nil {
				for i, column := range episode.Double.Columns {
					for j, challenge := range column.Challenges {
						if challenge.ChallengeID != 0 {
							episode_stats.DoubleCount += 1
							if challenge.TripleStumper {
								episode_stats.Stumpers[1] = append(episode_stats.Stumpers[1], qparty.Position{
									Column: uint(i),
									Index:  uint(j)})
							}
						}
					}
				}
			}

			// Index the episode metadata by its EpisodeID as that is all the client
			// has when listing a season index, before fetching the episode details.
			jarchive.Episodes[episode.EpisodeID] = episode_stats

			// Roll up the stats to the season-wide counts as well.
			season.ChallengesCount += episode_stats.SingleCount
			season.ChallengesCount += episode_stats.DoubleCount
			season.StumpersCount += len(episode_stats.Stumpers[0])
			season.StumpersCount += len(episode_stats.Stumpers[1])
			// TODO count final round triple-stumpers?
		}
	}

	write_metadata(jarchive, *data_path)
}

// Convenience wrapper around MkDirAll that returns the created path.
// If the directory cannot be created, it is treated as a fatal error.
func must_create_dir(at_path, child_path string) string {
	new_path := path.Join(at_path, child_path)
	err := os.MkdirAll(new_path, 0755)
	if err != nil {
		log.Fatal("failed to create directory", new_path)
	}
	return new_path
}

func WriteEpisodeJSON(episode qparty.FullEpisode, filepath string) error {
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

func NoOpEpisode(qparty.FullEpisode, string) error  { return nil }
func NoOpMetadata(qparty.SeasonIndex, string) error { return nil }
