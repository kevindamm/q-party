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
// github:kevindamm/q-party/json/seasons.go

package json

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path"
	"strconv"
)

type SeasonIndex struct {
	Version  []int                          `json:"version,omitempty"`
	Seasons  map[SeasonID]SeasonMetadata    `json:"seasons"`
	Episodes map[ShowNumber]EpisodeMetadata `json:"episodes"`
}

func LoadSeasonIndex(data_path string) SeasonIndex {
	season_index := SeasonIndex{}

	// TODO load all seasons from one file.
	season_index.Seasons = LoadSeasonsJSONL(path.Join(data_path, "seasons.jsonl"))

	return season_index
}

type SeasonMetadata struct {
	SeasonID `json:"id"`
	Name     string        `json:"name,omitempty"`
	Aired    ShowDateRange `json:"aired"`

	EpisodesCount   int `json:"episode_count"`
	ChallengesCount int `json:"challenge_count"`
	StumpersCount   int `json:"tripstump_count"`
}

// Unique (sometimes numeric) identifier for seasons in the archive.
type SeasonID string

// Returns a non-zero value if this season is part of regular play,
// zero otherwise (e.g. championship series, themed series).
func (id SeasonID) RegularSeason() int {
	number, err := strconv.Atoi(string(id))
	if err != nil {
		return 0
	}
	return number
}

func LoadSeasonsJSONL(jsonl_path string) map[SeasonID]SeasonMetadata {
	seasons_jsonl, err := os.Open(jsonl_path)
	if err != nil {
		return nil
	}
	scanner := bufio.NewScanner(seasons_jsonl)
	scanner.Split(bufio.ScanLines)

	seasons := make(map[SeasonID]SeasonMetadata, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		var season SeasonMetadata
		err = json.Unmarshal(line, &season)
		if err != nil {
			log.Fatal(err)
		}

		seasons[season.SeasonID] = season
	}
	return seasons
}
