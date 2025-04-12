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
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	data_path := flag.String("data-path", "./.data",
		"path where converted and created games are written")
	db_file := flag.String("db", "",
		"path (relative to data_path) where the sqlite3 database can be found")
	log_path := flag.String("log-path", "",
		"writes logging output to a file (empty string means not to log any output)")

	flag.Usage = cli_usage
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("ERROR: expected at least one command-line argument")
		fmt.Println()
		flag.Usage()
		return
	}

	var debug *log.Logger = log.Default()
	if *log_path != "" {
		logfile, err := os.Create(*log_path)
		if err != nil {
			log.Fatal(err)
		}
		debug = log.New(logfile, "", 0)
	}

	if *data_path == "" {
		*data_path = "."
	}
	jarchive, err := LoadLocalIndex(*data_path) // loads jarchive.jsonl and supporting seasons, episodes
	if err != nil {
		debug.Fatal(err)
	}

	if flag.NArg() < 1 {
		jarchive.GetSeasonList()
		return
	}

	//	switch flag.Arg(1) {
	//	case "season", "seasons":
	//		if flag.NArg() == 2 {
	//			FetchSeasonIndex(jarchive)
	//		}
	//		FetchSeason(jarchive, flag.Arg(2))
	//	case "episode", "episodes":
	//		if flag.NArg() == 2 {
	//			debug.Fatal("expected third argument, an episode ID to list")
	//		}
	//		episode_id := MustParseEpisodeID(flag.Arg(2))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		episode := FetchEpisode(jarchive, episode_id)
	//		fmt.Println(episode)
	//		// TODO write episode metadata to database, write episode
	//	}
	//
	if len(*db_file) == 0 {
		// upsert DB tables with index contents and fetched episode
	}
}

func cli_usage() {
	fmt.Printf("%s [season|episode] [id]\n  where\n", path.Base(os.Args[0]))
	fmt.Println("    *season* fetches season list, or episode info for season with id")
	fmt.Println("    *episode* fetches episode data, including rounds & categories & challenges")
	fmt.Println()
	fmt.Println("See README.md in the project's [source code](https://github.com/kevindamm/q-party/selfhost)")
	fmt.Println()
	flag.PrintDefaults()
}

func ListSeasons(jarchive JarchiveIndex) {
	for slug := range jarchive.GetSeasonList() {
		fmt.Println(slug)
	}
}

/*
//

//

//

//

func WriteEpisodeDB(dbclient *any) func(qparty.FullEpisode) error {
	log.Fatal("TODO (NYI)")
	return func(qparty.FullEpisode) error {
		return nil
	}
}

func WriteMetadataDB(dbclient *any) func(*service.JArchiveIndex) error {
	log.Fatal("TODO (NYI)")
	return func(*service.JArchiveIndex) error {
		return nil
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
*/
