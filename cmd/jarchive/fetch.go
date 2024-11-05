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
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func FetchAllSeasons(all_seasons map[JSID]JArchiveSeason, out_path string) error {
	seasons_path := path.Join(out_path, "seasons")
	episodes_path := path.Join(out_path, "episodes")
	os.MkdirAll(seasons_path, 0755)
	os.MkdirAll(episodes_path, 0755)

	for jsid, season := range all_seasons {
		log.Println("fetching season", jsid, season.Name)
		filepath := path.Join(seasons_path, jsid.HTML())
		if _, err := os.Stat(filepath); err == os.ErrNotExist {
			err = season.FetchIndex(jsid, filepath)
			if err != nil {
				return err
			}
		}

		reader, err := os.Open(path.Join(seasons_path, jsid.HTML()))
		if err != nil {
			return err
		}
		err = season.LoadSeasonMetadata(reader)
		if err != nil {
			return err
		}

		for jeid := range season.Episodes {
			filepath = path.Join(episodes_path, jeid.HTML())
			if _, err = os.Stat(filepath); err != os.ErrNotExist {
				log.Println("episode", jeid.HTML(), "does not exist")
				continue
			}
			err = FetchEpisode(jeid, filepath)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
