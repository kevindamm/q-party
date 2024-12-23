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
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	qparty "github.com/kevindamm/q-party"
	"github.com/kevindamm/q-party/encoding/html"
	"github.com/kevindamm/q-party/encoding/json"
	"github.com/kevindamm/q-party/ent"
	"github.com/kevindamm/q-party/service"
)

func main() {
	data_path := flag.String("data-path", "./.data",
		"path where converted and created games are written")
	db_path := flag.String("db-path", "", "path of the SQLite database file (don't write if empty)")
	init_db := flag.Bool("init-db", false,
		"initialize the local DB file with the initial metadata, using content in embedded JSON files")
	write_jsonl := flag.Bool("write-jsonl", true,
		"write the updates and new data to JSONL files in $data-path")
	//log_path := flag.String("log-path", "",
	//	"writes logging output to a file (empty string means not to log any output)")

	flag.Usage = cli_usage
	flag.Parse()

	var client *ent.Client
	if *db_path != "" {
		client = must_open_db(*db_path)
		if *init_db {
			must_drop_existing_tables(client)
			populate_tables_from_json(client, jsonFS)
		}
	} else {
		if *init_db {
			log.Fatal("must set -db-path if -init-db is true")
		}
	}

	if flag.NArg() < 1 {
		fmt.Println("ERROR: expected at least one command-line argument")
		fmt.Println()
		flag.Usage()
		return
	}

	if *data_path == "" {
		*data_path = "."
	}
	jarchive, err := read_json_data(*data_path)
	if err != nil {
		log.Fatal(err)
	}

	switch flag.Arg(0) {
	case "list":
		if flag.NArg() < 2 {
			list_seasons(jarchive)
			return
		}
		switch flag.Arg(1) {
		case "season", "seasons":
			if flag.NArg() == 2 {
				list_seasons(jarchive)
			}
			list_season_id(jarchive, flag.Arg(2))
		case "episode":
			if flag.NArg() == 2 {
				log.Fatal("expected third argument, an episode ID to list")
			}
			show_number, err := qparty.ParseShowNumber(flag.Arg(2))
			if err != nil {
				log.Fatal(err)
			}
			episode := list_episode_id(jarchive, show_number)
			fmt.Println(episode)
			if *write_jsonl {
				// TODO
				json.WriteEpisodeMetadata(*data_path, episode)
			}
		}

	case "play":
		if flag.NArg() < 2 {
			log.Fatal("expected a play version (season/show#, 'episode', 'category' or 'challenge')")
		}
		switch flag.Arg(1) {
		case "episode", "episodes":
			// TODO
		case "category", "categories":
			// TODO
		case "challenge", "challenges":
			// TODO
		default:
			// TODO attempt to parse arg-1 as season name and show # separated by '/'.
			show_number, err := qparty.ParseShowNumber(flag.Arg(1))
			if err != nil {
				log.Fatal(err)
			}
			// TODO load episode from show_number and play it
			_ = show_number
		}

	default:
		fmt.Println("ERROR: expected command (list | play | serve | build)")
		fmt.Println()
		flag.Usage()
	}
}

//go:embed json/*.json
var jsonFS embed.FS

func cli_usage() {
	fmt.Printf("%s [list|play]\n  where\n", path.Base(os.Args[0]))
	fmt.Println("    *list* lists season or episode info")
	fmt.Println("    *play* starts an interactive session")
	fmt.Println()
	fmt.Println("list takes arguments 'season' or 'episode' or no additional argument")
	fmt.Println("  with no argument, it lists all seasons and count of available episodes")
	fmt.Println("  with 'season' argument and season-ID it prints information about that season")
	fmt.Println("  with 'episode' argument and 'season/show' ID, prints episode information")
	fmt.Println()
	fmt.Println("play takes arguments [season/show ID] (playing a selected show)")
	fmt.Println("  or 'episode' or 'category' or 'challenge' (playing a random selection)")
	fmt.Println()
	fmt.Println("See README.md in the project's [source code](https://github.com/kevindamm/q-party)")
	fmt.Println()
	flag.PrintDefaults()
}

//

//

//

//

func WriteEpisodeMetadata(
	jarchive *service.JArchiveIndex,
	data_path string,
	output_format string) {

	var write_episode func(qparty.FullEpisode) error
	var write_metadata func(*service.JArchiveIndex) error

	season_episodes := make(map[qparty.SeasonID][]qparty.EpisodeID)
	for jeid, episode := range jarchive.Episodes {
		season_episodes[episode.Show.Season] = append(
			season_episodes[episode.Show.Season], jeid)
	}

	switch output_format {
	case "":
		write_episode = func(qparty.FullEpisode) error { return nil }
		write_metadata = func(*service.JArchiveIndex) error { return nil }
	case "json":
		write_metadata = json.WriteSeasonIndex(data_path)
		json_path := must_create_dir(data_path, "json")
		write_episode = json.WriteEpisode(json_path)
	case "sqlite":
		dbclient := must_open_db(data_path)
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

			episode := html.LoadEpisodeHTML(filepath)
			episode.EpisodeID = jeid
			episode.Show.Season = jsid

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

func must_open_db(sqlite_path string) *ent.Client {
	client, err := ent.Open("sqlite3", "file:"+sqlite_path+"?cache=shared&_fk=1")
	if err != nil {
		log.Fatal("failed to open DB\n", err)
	}
	return client
}

func must_drop_existing_tables(client *ent.Client) {
	// TODO

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

func list_seasons(jarchive *service.JArchiveIndex) {
	for _, season := range jarchive.Seasons {
		fmt.Println(season)
	}
}

func list_season_id(jarchive *service.JArchiveIndex, season_id string) {
	season, ok := jarchive.Seasons[qparty.SeasonID(season_id)]
	if !ok {
		log.Fatalf("season '%s' not a known Season ID", season_id)
	}

	fmt.Println("# %s (%s)", season.Title, season.SeasonID)
	for _, episode := range jarchive.EpisodesBySeason(season.SeasonID) {
		fmt.Println(episode)
	}
	fmt.Println()
}

func list_episode_id(jarchive *service.JArchiveIndex, show *qparty.ShowNumber) qparty.EpisodeMetadata {
	episode := jarchive.GetShowEpisode(*show)
	fmt.Println(episode)
	return episode
}

func play_episode(jarchive *service.JArchiveIndex, show *qparty.ShowNumber) {
	season, ok := jarchive.Seasons[qparty.SeasonID(show.Season)]
	if !ok {
		log.Fatalf("season '%s' not a known Season ID", show.Season)
	}

	// TODO
	_ = season

}

func play_random_episode(jarchive *service.JArchiveIndex) {
	log.Fatal("TBD (work in progress)")
}

func play_random_categories(jarchive *service.JArchiveIndex) {
	log.Fatal("TBD (work in progress)")
}

func play_random_challenges(jarchive *service.JArchiveIndex) {
	log.Fatal("TBD (work in progress)")
}
