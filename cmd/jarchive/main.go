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
	"bufio"
	gojson "encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/kevindamm/q-party/cmd/jarchive/html"
	"github.com/kevindamm/q-party/ent"
	"github.com/kevindamm/q-party/json"
)

func main() {
	data_path := flag.String("data", ".data",
		"path where converted and created games are written")
	flag.Usage = func() {
		fmt.Printf("%s command id# [flags]\n", path.Base(os.Args[0]))
		fmt.Println("  where")
		fmt.Println("    command is either 'season' or 'episode'")
		fmt.Println("    id# is the unique ID for the season or episode")
		fmt.Println()
		flag.PrintDefaults()
	}
	flag.Parse()

	//	if flag.NArg() < 2 {
	//		flag.Usage()
	//		return
	//	}

	jarchive := json.SeasonIndex{
		Version:  []uint{1, 0},
		Episodes: make(map[json.ShowNumber]json.EpisodeMetadata, 10_000),
	}
	seasons := html.MustLoadAllSeasons(*data_path)

	log.Print("loaded ", len(seasons.Seasons), " seasons")
	jarchive.Seasons = make(map[json.SeasonID]json.SeasonMetadata, len(seasons.Seasons)+1)
	for k, v := range seasons.Seasons {
		key := json.SeasonID(k)
		number := key.RegularSeason()
		if number > 0 {
			key = json.SeasonID(fmt.Sprintf("%d", number))
		}
		modernize_season(&v)
		jarchive.Seasons[key] = v
	}

	for jsid := range jarchive.Seasons {
		var season_episodes SeasonEpisodes
		season_json := path.Join(*data_path, "seasons", jsid.JSON())
		bytes, err := os.ReadFile(season_json)
		if err != nil {
			log.Fatal(err)
		}
		err = gojson.Unmarshal(bytes, &season_episodes)
		if err != nil {
			log.Fatal(err)
		}

		for jeid, episode_dates := range season_episodes.Episodes {
			episode, err := LoadEpisodeMetadataForSeason(
				path.Join(*data_path, "episodes", jeid.HTML()), jeid, jsid)
			if os.IsNotExist(err) {
				continue
			}
			if err != nil {
				log.Print("ERROR:", err)
				continue
			}

			if !episode_dates.Aired.IsZero() {
				episode.Aired = episode_dates.Aired
			}
			if !episode_dates.Taped.IsZero() {
				episode.Taped = episode_dates.Taped
			}
			jarchive.Episodes[episode.ShowNumber] = episode
		}
	}

	filepath := path.Join(*data_path, "jarchive.json")
	log.Print("writing all seasons to a single file ", filepath)

	writer, err := os.Create(filepath)
	if err != nil {
		log.Fatal("failed to create file ", filepath, "\n", err)
	}
	defer writer.Close()

	bytes, err := gojson.MarshalIndent(jarchive, "", "  ")
	if err != nil {
		log.Fatal("failed to marshal index into json\n", err)
	}

	_, err = writer.Write(bytes)
	if err != nil {
		log.Fatal("failed to write bytes\n", err)
	}
}

func LoadSeasonsJSONL(jsonl_path string) map[json.SeasonID]json.SeasonMetadata {
	seasons_jsonl, err := os.Open(jsonl_path)
	if err != nil {
		return nil
	}
	scanner := bufio.NewScanner(seasons_jsonl)
	scanner.Split(bufio.ScanLines)

	seasons := make(map[json.SeasonID]json.SeasonMetadata)
	for scanner.Scan() {
		line := scanner.Bytes()
		var season json.SeasonMetadata
		err = gojson.Unmarshal(line, &season)
		if err != nil {
			log.Fatal(err)
		}
		modernize_season(&season)
		seasons[season.SeasonID] = season
	}
	return seasons
}

func LoadEpisodeMetadataForSeason(filepath string, jeid html.JEID, jsid json.SeasonID) (json.EpisodeMetadata, error) {
	var metadata json.EpisodeMetadata
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return metadata, err
	}

	reader, err := os.Open(filepath)
	if err != nil {
		return metadata, err
	}
	defer reader.Close()

	episode := html.ParseEpisode(jeid, reader)
	metadata = episode.EpisodeMetadata.EpisodeMetadata
	metadata.SeasonID = jsid
	for i := range 3 {
		metadata.ContestantIDs[i] = episode.Contestants[i].ContestantID
	}
	return metadata, nil
}

//
//
//
//

type SeasonEpisodes struct {
	json.SeasonID `json:"season"`
	Name          string                     `json:"name"`
	Count         int                        `json:"count"`
	Episodes      map[html.JEID]EpisodeDates `json:"episodes"`
}

type EpisodeDates struct {
	html.JEID `json:"-"`
	Aired     json.ShowDate `json:"aired"`
	Taped     json.ShowDate `json:"taped"`
}

//
//
//
//

func LegacyMain(data_path *string) {
	switch flag.Arg(0) {

	case "fetch":
		jeid := html.JEID(json.MustParseShowNumber(flag.Arg(1)))
		episodes_path := create_dir(*data_path, "episodes")
		filepath := path.Join(episodes_path, jeid.HTML())

		err := html.FetchEpisode(jeid, filepath)
		if err != nil {
			log.Fatal(err)
		}

	case "convert":
		seasons := json.LoadSeasonsJSON(*data_path)

		/// TODO DELETE
		var sqlclient ent.Client
		if flag.Arg(1) == "*" {
			ConvertAllEpisodes(*data_path, seasons, sqlclient)
			break
		}
		/// TODO DELETE

		jeid := json.MustParseShowNumber(flag.Arg(1))
		metadata := seasons.Episodes[jeid]

		filepath := path.Join(*data_path, "episodes", jeid.HTML())
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			log.Fatal("episode", jeid, "HTML does not exist", filepath,
				"\n", err)
		}

		err := ConvertEpisode(html.JEID(jeid), metadata, *data_path, sqlclient)
		if err != nil {
			log.Fatal(err)
		}

	default:
		flag.Usage()
	}
}

// MkdirAll to ensure path exists and return the joined path.
// No file change if the path already existed.
func create_dir(at_path, child_path string) string {
	new_path := path.Join(at_path, child_path)
	err := os.MkdirAll(new_path, 0755)
	if err != nil {
		log.Fatal("failed to create directory", new_path)
	}
	return new_path
}
