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
// github:kevindamm/q-party/cmd/jarchive/util.go

package main

// Utility functions for main() functionality.

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	qparty "github.com/kevindamm/q-party"
	"github.com/kevindamm/q-party/ent"
	"github.com/kevindamm/q-party/service"
)

// Reads the JArchiveIndex contents of local file system path (not embedded).
func read_json_data(base_path string) (*service.JArchiveIndex, error) {
	jarchive := new(service.JArchiveIndex)

	// jarchive (full index in one .json file)
	load_file(path.Join(base_path, "jarchive.json"), func(bytes []byte) error {
		err := json.Unmarshal(bytes, jarchive)
		if err != nil {
			return err
		}
		return nil
	})

	// seasons
	load_file(path.Join(base_path, "json/seasons.jsonl"), func(bytes []byte) error {
		for season := range scan_json_lines[qparty.SeasonMetadata](bytes) {
			jarchive.Seasons[season.SeasonID] = *season
			fmt.Printf("season %s (%s)", season.Name, season.SeasonID)
		}
		return nil
	})

	// episodes
	load_file(path.Join(base_path, "json/episodes.jsonl"), func(bytes []byte) error {
		for episode := range scan_json_lines[qparty.EpisodeMetadata](bytes) {
			jarchive.Episodes[episode.EpisodeID] = qparty.EpisodeStats{EpisodeMetadata: *episode}
			fmt.Printf("show %s episode [%s] (jeid: %d)", episode.Show, episode.ShowTitle, episode.EpisodeID)
		}
		return nil
	})

	// categories
	load_file(path.Join(base_path, "json/categories.jsonl"), func(bytes []byte) error {
		for category := range scan_json_lines[qparty.CategoryMetadata](bytes) {
			cat_meta, found := jarchive.Categories[category.Title]
			if !found {
				jarchive.Categories[category.Title] = *category
			} else {
				cat_meta.Episodes = append(cat_meta.Episodes, category.Episodes...)
			}
			fmt.Printf("category %s", category.Title)
		}
		return nil
	})

	return jarchive, nil
}

// Load a file from local filesystem and process its bytes with [procfn]
// which as used above has a closure around a common structure that receives
// all updates found in subsequent files being loaded similarly.
func load_file(filename string, procfn func(bytes []byte) error) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}

	if bytes, err := os.ReadFile(filename); err != nil {
		return err
	} else {
		if err := procfn(bytes); err != nil {
			return err
		}
	}
	return nil
}

// Utility for sending each line (as []byte) from a larger byte slice.
// Skips empty lines.
func scan_json_lines[T any](byteslice []byte) <-chan *T {
	json_chan := make(chan *T)
	go func() {
		defer close(json_chan)
		line_count := 0
		scanner := bufio.NewScanner(bytes.NewReader(byteslice))
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			line := scanner.Bytes()
			line_count++
			if len(line) == 0 {
				continue
			}
			var jsondata *T = new(T)
			err := json.Unmarshal(line, jsondata)
			if err != nil {
				log.Fatalf("ERROR failed to unmarshal data at line %d:\n%s\n",
					line_count, line)
				continue
			}
			json_chan <- jsondata
		}
	}()

	return json_chan
}

// Create database tables and populate with the minimal index from embedded FS.
func populate_tables_from_json(client *ent.Client, jsonFiles embed.FS) error {
	// TODO
	_ = client
	_ = jsonFiles
	return nil
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
