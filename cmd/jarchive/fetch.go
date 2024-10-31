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
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func FetchAllSeasons(all_seasons []JArchiveSeason, out_path string) error {
	season_path := path.Join(out_path, "seasons")
	episodes_path := path.Join(out_path, "episodes")
	os.MkdirAll(season_path, 0755)
	os.MkdirAll(episodes_path, 0755)

	for _, season := range all_seasons {
		fmt.Printf("season: %s (%s)\n", season.Season, season.Name)
		filepath := path.Join(season_path, fmt.Sprintf("%s.html", season.Season))
		if _, err := os.Stat(filepath); err == os.ErrNotExist {
			err = season.FetchIndex(filepath)
			if err != nil {
				return err
			}
		}

		episode_ids, err := season.FindEpisodeIDs(season_path)
		if err != nil {
			return err
		}
		for _, episode_id := range episode_ids {
			filepath = path.Join(episodes_path, fmt.Sprintf("%s.html", episode_id))
			if _, err = os.Stat(filepath); err != os.ErrNotExist {
				continue
			}
			err = FetchEpisode(episode_id, filepath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func FetchEpisode(episode string, filepath string) error {
	episode_id, err := strconv.Atoi(episode)
	if err != nil {
		return fmt.Errorf("failed to convert episode id '%s'\n%s", episode, err)
	}
	url := episode_url(episode_id)
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
