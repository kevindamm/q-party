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
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func ConvertAllEpisodes(seasons_path string, episodes_path string, out_path string) {
	seasons := make([]JArchiveSeason, 0, 50)
	err := json.Unmarshal([]byte(all_seasons), &seasons)
	if err != nil {
		log.Fatalf("failed to unmarshal season metadata; %s", err)
	}

	err = os.MkdirAll(out_path, 0755)
	if err != nil {
		log.Fatalf("failed to create directory for converted episodes %s\n%s", out_path, err)
	}

	for _, season := range seasons {
		episodes, err := season.FindEpisodeIDs(seasons_path)
		if err != nil {
			log.Fatalf("failed to parse season %s index file", season.Season)
		}
		for _, episode := range episodes {
			ep_path := path.Join(out_path, fmt.Sprintf("%s.html", episode))
			reader, err := os.Open(ep_path)
			if err != nil {
				log.Print("could not open episode ", episode, err)
				continue
			}

			filename := fmt.Sprintf("%s.json", episode)
			filepath := path.Join(out_path, filename)
			writer, err := os.Create(filepath)
			if err != nil {
				log.Print("could not create json file for episode ", episode, err)
				continue
			}

			err = ConvertEpisode(episode, reader, writer)
			if err != nil {
				log.Print("could not convert episode ", episode, err)
				continue
			}
		}
	}
}

func ConvertEpisode(ep_id string, reader io.Reader, writer io.Writer) error {
	jgame := ParseEpisode(ep_id, reader)

	jgame_json, err := json.MarshalIndent(jgame, "", "  ")
	if err != nil {
		return err
	}
	_, err = writer.Write(jgame_json)

	return err
}
