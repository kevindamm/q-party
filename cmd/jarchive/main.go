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
	"github.com/kevindamm/q-party/ent"
	"github.com/kevindamm/q-party/service"
)

func main() {
	data_path := flag.String("data", ".data",
		"path where converted and created games are written")
	//	sqlite_path := flag.String("db", "jarchive.sqlite",
	//		"path of the file (within `data` path) that represents the sqlite3 database")
	output_format := flag.String("out", "",
		"(json|sqlite) encoding of the converted season or episode representation")

	flag.Usage = func() {
		fmt.Printf("%s command id# [flags]\n", path.Base(os.Args[0]))
		fmt.Println("  where")
		fmt.Println("    command is either 'season' or 'episode'")
		fmt.Println("    id# is the unique ID for the season or episode")
		fmt.Println()
		fmt.Println("Fetches the season or episode if absent, then converts it " +
			"and stores extracted data into the jarchive.sqlite database.")
		fmt.Println()
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}
	jarchive_path := path.Join(*data_path, "jarchive.json")
	jarchive_bytes, err := os.ReadFile(jarchive_path)
	if err != nil {
		log.Fatal(err)
	}

	var jarchive service.JArchiveIndex
	err = json.Unmarshal(jarchive_bytes, &jarchive)
	if err != nil {
		log.Print("failed to unmarshal JArchive JSON")
		log.Fatal(err)
	}

	switch flag.Arg(0) {
	case "season":
		if flag.NArg() < 2 {
			log.Print("ERROR: expected season ID after `season` command")
		}
		LogSeasonDetails(flag.Arg(1))

	case "episode":
		if flag.NArg() < 2 {
			log.Print("ERROR: expected episode ID after `episode` command")
		}
		LogEpisodeDetails(flag.Arg(1))

	case "all":
		WriteEpisodeMetadata(&jarchive, *data_path, *output_format)
	}
}

func WriteEpisodeMetadata(
	jarchive *service.JArchiveIndex,
	data_path string,
	output_format string) {

	var write_episode func(qparty.FullEpisode) error
	var write_metadata func(*service.JArchiveIndex) error

	season_episodes := make(map[qparty.SeasonID][]qparty.EpisodeID)
	for jeid, episode := range jarchive.Episodes {
		season_episodes[episode.SeasonID] = append(
			season_episodes[episode.SeasonID], jeid)
	}

	switch output_format {
	case "":
		write_episode = NoOpEpisode
		//write_metadata = NoOpMetadata
		write_metadata = WriteSeasonIndexJSON(data_path)
	case "json":
		write_metadata = WriteSeasonIndexJSON(data_path)
		json_path := must_create_dir(data_path, "json")
		write_episode = WriteEpisodeJSON(json_path)
	case "sqlite":
		sqlite_path := fmt.Sprintf("%s/jarchive.sqlite", data_path)
		dbclient := must_open_db(sqlite_path)
		write_metadata = WriteMetadataDB(dbclient)
		write_episode = WriteEpisodeDB(dbclient)
	default:
		log.Print("unrecognized output format ", output_format)
		flag.Usage()
		return
	}

	// Read per-season episode list (seasons/[seasonid].json)
	for jsid, season := range jarchive.Seasons {
		for _, jeid := range season_episodes[jsid] {
			filepath := path.Join(data_path, "episodes", jeid.HTML())
			if _, err := os.Stat(filepath); os.IsNotExist(err) {
				log.Print("HTML not found (fetch?) ", jeid.HTML())
				continue
			}

			episode := LoadEpisode(filepath)
			episode.EpisodeID = jeid
			episode.SeasonID = jsid

			err := write_episode(*episode)
			if err != nil {
				log.Print("ERROR writing episode: ", err)
			}

			episode_stats := jarchive.Episodes[jeid]
			episode_stats.EpisodeMetadata = episode.EpisodeMetadata
			episode_stats.Stumpers = make([][]qparty.Position, 2)

			// Stats (count challenges, triple-stumpers) for the first round.
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
			// Stats (count challenges, triple-stumpers) for the second round.
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
			// TODO also count correct/incorrect responses from contestants (proxy for difficulty)

			jarchive.Episodes[episode.EpisodeID] = episode_stats

			// Aggregate the stats to the season-wide counts as well.
			season.ChallengesCount += episode_stats.SingleCount
			season.ChallengesCount += episode_stats.DoubleCount
			season.StumpersCount += len(episode_stats.Stumpers[0])
			season.StumpersCount += len(episode_stats.Stumpers[1])
			jarchive.Seasons[jsid] = season
		}
	}

	err := write_metadata(jarchive)
	if err != nil {
		log.Fatal("failed to write JArchive index\n", err)
	}
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

func NoOpEpisode(qparty.FullEpisode) error      { return nil }
func NoOpMetadata(*service.JArchiveIndex) error { return nil }

func LoadEpisode(html_path string) *qparty.FullEpisode {
	// Parse the episode's HTML to get the show and challenge details.
	reader, err := os.Open(html_path)
	if err != nil {
		// Unlikely, we know the file exists at this point, but state changes...
		log.Print("ERROR: ", err)
		return nil
	}
	defer reader.Close()

	episode := html.ParseEpisode(reader)
	return episode
}

func WriteEpisodeJSON(json_path string) func(qparty.FullEpisode) error {
	return func(episode qparty.FullEpisode) error {
		filepath := path.Join(json_path, episode.ShowNumber.JSON(episode.SeasonID))
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
}

func WriteSeasonIndexJSON(data_path string) func(*service.JArchiveIndex) error {
	return func(jarchive *service.JArchiveIndex) error {
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
}

func must_open_db(sqlite_path string) *ent.Client {
	client, err := ent.Open("sqlite3", "file:"+sqlite_path+"?cache=shared&_fk=1")
	if err != nil {
		log.Fatal("failed to open DB\n", err)
	}
	return client
}

func WriteEpisodeDB(dbclient *ent.Client) func(qparty.FullEpisode) error {
	log.Fatal("TODO (NYI)")
	return func(qparty.FullEpisode) error {
		return nil
	}
}

func WriteMetadataDB(dbclient *ent.Client) func(*service.JArchiveIndex) error {
	log.Fatal("TODO (NYI)")
	return func(*service.JArchiveIndex) error {
		return nil
	}
}

func LogSeasonDetails(season string) {

}

func LogEpisodeDetails(season string) {

}
