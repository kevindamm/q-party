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
// github:kevindamm/q-party/cmd/jarchive/converter.go

package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path"

	"github.com/kevindamm/q-party/ent"
)

func ConvertAllEpisodes(data_path string, sqlclient ent.Client) error {
	seasons := LoadAllSeasons(data_path)
	episodes_path := path.Join(data_path, "episodes")
	err := os.MkdirAll(episodes_path, 0755)
	if err != nil {
		log.Fatal("failed to create directory for converted episodes", data_path,
			"\n", err)
	}

	for jsid, season := range seasons {
		log.Println("Converting episodes from season", jsid, season.Name)
		log.Println("~ ~ ~ ~ ~ ~ ~ ~ ~ ~")
		for jeid, episode := range season.Episodes {
			reader, err := os.Open(path.Join(episodes_path, jeid.HTML()))
			if err != nil {
				log.Print("could not open episode", episode, err)
				continue
			}

			writer, err := os.Create(path.Join(data_path, jeid.JSON()))
			if err != nil {
				log.Print("could not create json file for episode", jeid, err)
				continue
			}

			err = ConvertEpisode(episode.JEID, reader, writer)
			if err != nil {
				log.Print("could not convert episode", episode, err)
				continue
			}
		}
	}

	return nil
}

func ConvertEpisode(jeid JEID, reader io.Reader, writer io.Writer) error {
	episode := ParseEpisode(jeid, reader)
	episode_json, err := json.MarshalIndent(episode, "", "  ")
	if err != nil {
		return err
	}
	nbytes, err := writer.Write(episode_json)
	log.Println("writing episode,", nbytes, "bytes written")
	return err
}
