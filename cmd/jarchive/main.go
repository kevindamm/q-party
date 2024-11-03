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
)

func main() {
	data_path := flag.String("data", "./.data",
		"path where converted and created games are written")
	flag.Usage = func() {
		fmt.Printf("%s command episode# [flags]\n", os.Args[0])
		fmt.Println("  where")
		fmt.Println("    command is either 'fetch' or 'season' or 'convert'")
		fmt.Println("    episode# is the index ID for the episode")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
	}

	switch os.Args[1] {

	case "fetch":
		episodes_path := path.Join(*data_path, "episodes")
		episode_id := os.Args[2]
		jeid := MustParseJEID(episode_id)
		filepath := path.Join(episodes_path, fmt.Sprintf("%d.html", jeid))

		err := FetchEpisode(jeid, filepath)
		if err != nil {
			log.Fatal(err)
		}

	case "season":
		jsid := JSID(os.Args[2])
		season := JArchiveSeason{Season: jsid}

		seasons_path := path.Join(*data_path, "seasons")
		err := os.MkdirAll(seasons_path, 0755)
		if err != nil {
			log.Fatalf("failed to make a directory for writing Season metadata")
		}
		filepath := path.Join(seasons_path, fmt.Sprintf("%s.html", jsid))

		reader, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		err = season.LoadSeasonMetadata(reader)
		if err != nil {
			log.Fatalf("failed to load season '%s' metadata\n%s", jsid, err)
		}

		bytes, err := json.Marshal(season)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bytes))

	case "convert":
		jeid := MustParseJEID(os.Args[2])
		ep_path := path.Join(*data_path, "episodes")
		filename := fmt.Sprintf("%d.html", jeid)

		reader, err := os.Open(path.Join(ep_path, filename))
		if err != nil {
			log.Fatalf("could not open '%s'\n%s", ep_path, err)
		}
		defer reader.Close()

		filename = filename[:len(filename)-4] + "json"
		filepath := path.Join(*data_path, "episodes", filename)
		writer, err := os.Create(filepath)
		if err != nil {
			log.Fatalf("could not create json file for episode %s\n%s", filename, err)
		}
		defer writer.Close()

		err = ConvertEpisode(jeid, reader, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}

	default:
		flag.Usage()
	}
}
