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
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/kevindamm/q-party/ent"
)

func main() {
	data_path := flag.String("data", "../.data",
		"path where converted and created games are written")
	flag.Usage = func() {
		fmt.Printf("%s command episode# [flags]\n", path.Base(os.Args[0]))
		fmt.Println("  where")
		fmt.Println("    command is either 'fetch' or 'convert'")
		fmt.Println("    episode# is the index ID for the episode")
		fmt.Println()
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	//? type identical to cue:EpisodeMetadata?
	type EpisodeMetadata struct {
		JSID       `json:"season"`
		JEID       `json:"episode"`
		ShowNumber int `json:"show_number"`

		JArchiveEpisodeMetadata
	}
	//? type identical to cue:AllSeasonsMetadata?
	type AllSeasonsMetadata struct {
		Version  []uint                   `json:"version"`
		Seasons  map[JSID]JArchiveSeason  `json:"seasons"`
		Episodes map[JEID]EpisodeMetadata `json:"episodes"`
	}

	metadata := AllSeasonsMetadata{
		Version:  []uint{0, 9},
		Seasons:  LoadAllSeasons(*data_path),
		Episodes: make(map[JEID]EpisodeMetadata)}

	log.Print("loaded", len(metadata.Seasons), "seasons")
	filepath := path.Join(*data_path, "seasons.json")
	log.Print("writing all seasons to a single file", filepath)

	writer, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	for jsid, season := range metadata.Seasons {
		metadata.Seasons[jsid] = season
		for jeid, jarchive_meta := range season.Episodes {
			episode, err := LoadEpisode(path.Join(*data_path, jeid.HTML()), jarchive_meta)
			if err != nil {
				log.Print("ERROR:", err)
				continue
			}
			ep_meta := EpisodeMetadata{
				jsid, jeid, episode.ShowNumber, jarchive_meta}
			metadata.Episodes[jeid] = ep_meta
		}
	}

	bytes, err := json.Marshal(metadata)
	if err != nil {
		log.Fatal("failed to marshal seasons to JSON bytes")
	}
	nbytes, err := writer.Write(bytes)
	if err != nil {
		log.Fatal("failed to write", filepath, "\n", err)
	} else {
		log.Printf("Wrote seasons.json, %d bytes", nbytes)
	}
}

// TODO season util

func LegacyMain(data_path *string) {
	switch flag.Arg(0) {

	case "fetch":
		jeid := MustParseJEID(flag.Arg(1))
		episodes_path := create_dir(*data_path, "episodes")
		filepath := path.Join(episodes_path, jeid.HTML())

		err := FetchEpisode(jeid, filepath)
		if err != nil {
			log.Fatal(err)
		}

	case "convert":
		seasons := LoadAllSeasons(*data_path)

		/// TODO DELETE
		var sqlclient ent.Client
		if flag.Arg(1) == "*" {
			ConvertAllEpisodes(*data_path, seasons, sqlclient)
			break
		}
		/// TODO DELETE

		jeid := MustParseJEID(flag.Arg(1))
		var metadata *JArchiveEpisodeMetadata
		for _, season := range seasons {
			for id, episode := range season.Episodes {
				if id == jeid {
					metadata = &episode
				}
			}
		}

		filepath := path.Join(*data_path, "episodes", jeid.HTML())
		if _, err := os.Stat(filepath); err == os.ErrNotExist {
			log.Fatal("episode", jeid, "HTML does not exist", filepath,
				"\n", err)
		}

		err := ConvertEpisode(jeid, *metadata, *data_path, sqlclient)
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

func FetchEpisode(episode JEID, filepath string) error {
	url := episode.URL()
	log.Print("Fetching ", url, "  -> ", filepath)

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	return nil
}
