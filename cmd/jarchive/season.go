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
// github:kevindamm/q-party/cmd/jarchive/season.go

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"time"
)

type JArchiveSeason struct {
	JSID  `json:"season"` // TODO json:"-"
	Name  string          `json:"name"`
	Aired ShowDateRange   `json:"aired"`
	Count int             `json:"count"`

	Episodes map[JEID]JArchiveEpisodeMetadata `json:"-"`
}

// Loads the season index and the metadata of each season in the index.
// Parameter [data_path] is the location of the seasons.jsonl, the same path
// as the `seasons` and `episodes` directory (.data/ relative to $CWD)
func LoadAllSeasons(data_path string) map[JSID]JArchiveSeason {
	jsonl_path := path.Join(data_path, "seasons.jsonl")
	seasons_dir := path.Join(data_path, "seasons")

	seasons := LoadSeasonsJSONL(jsonl_path)
	for jsid, season := range seasons {
		file_path := path.Join(seasons_dir, jsid.JSON())
		reader, err := os.Open(file_path)
		if err != nil {
			log.Fatalf("failed to open file path '%s'\n%s", file_path, err)
		}
		err = season.LoadSeasonMetadata(reader)
		if err != nil {
			log.Fatalf("failed to parse season %s index file", jsid)
		}
	}

	return seasons
}

func LoadSeasonsJSONL(jsonl_path string) map[JSID]JArchiveSeason {
	seasons_jsonl, err := os.Open(jsonl_path)
	if err != nil {
		return nil
	}
	scanner := bufio.NewScanner(seasons_jsonl)
	scanner.Split(bufio.ScanLines)

	seasons := make(map[JSID]JArchiveSeason, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		var season JArchiveSeason
		err := json.Unmarshal(line, &season)
		if err != nil {
			log.Fatal(err)
		}

		seasons[season.JSID] = season
	}
	return seasons
}

func (season *JArchiveSeason) LoadSeasonMetadata(reader io.Reader) error {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	reEpisodeLink := regexp.MustCompile(`"showgame\.php\?game_id=(\d+)"(.*)\n`)
	reTapedDate := regexp.MustCompile(`[tT]aped\s+(\d{4})-(\d{2})-(\d{2})`)
	reAiredDate := regexp.MustCompile(`[aA]ired.*(\d{4})-(\d{2})-(\d{2})`)
	reSeasonName := regexp.MustCompile(`<h2 class="season">(.*)</h2>`)

	match := reSeasonName.FindSubmatch(bytes)
	if len(match) > 0 {
		season.Name = string(match[1])
	}

	season.Episodes = make(map[JEID]JArchiveEpisodeMetadata)

	matches := reEpisodeLink.FindAllSubmatch(bytes, -1)
	for _, match := range matches {
		episode := JArchiveEpisodeMetadata{}
		episode.JEID = MustParseJEID(string(match[1]))
		taped := reTapedDate.FindSubmatch(match[2])
		if taped != nil {
			episode.Taped = parseTimeYYYYMMDD(taped[1], taped[2], taped[3])
		}
		aired := reAiredDate.FindSubmatch(match[2])
		if aired != nil {
			episode.Aired = parseTimeYYYYMMDD(aired[1], aired[2], aired[3])
		}
		season.Episodes[episode.JEID] = episode
	}

	return nil
}

func (season JArchiveSeason) FetchIndex(jsid JSID, filepath string) error {
	url := jsid.URL()
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

// Unique (sometimes numeric) identifier for seasons in the archive.
type JSID string

// Returns a non-zero value if this season is part of regular play,
// zero otherwise.
func (id JSID) RegularSeason() int {
	number, err := strconv.Atoi(string(id))
	if err != nil {
		return 0
	}
	return number
}

func (id JSID) HTML() string {
	return fmt.Sprintf("%s.html", id)
}

func (id JSID) JSON() string {
	return fmt.Sprintf("%s.json", id)
}
func (id JSID) URL() string {
	return fmt.Sprintf("https://j-archive.com/showseason.php?season=%s", id)
}
