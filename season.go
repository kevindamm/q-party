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

package qparty

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type SeasonIndex struct {
	Version  []uint                         `json:"version,omitempty"`
	Seasons  map[SeasonID]SeasonMetadata    `json:"seasons"`
	Episodes map[ShowNumber]EpisodeMetadata `json:"episodes"`
}

type SeasonMetadata struct {
	SeasonID `json:"id"`
	Season   string        `json:"season,omitempty"` // deprecated
	Name     string        `json:"name,omitempty"`   // deprecated
	Title    string        `json:"title"`
	Aired    ShowDateRange `json:"aired"`

	Count           int `json:"count,omitempty"` // deprecated
	EpisodesCount   int `json:"episode_count"`
	ChallengesCount int `json:"challenge_count"`
	StumpersCount   int `json:"tripstump_count"`
}

type EpisodeMap map[uint]EpisodeMetadata

func (all_seasons SeasonIndex) WriteSeasonIndexJSON(json_path string) error {
	writer, err := os.Create(json_path)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(all_seasons)
	if err != nil {
		return fmt.Errorf("failed to marshal seasons to JSON bytes\n%s", err)
	}
	nbytes, err := writer.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write%s\n%s", json_path, err)
	} else {
		log.Printf("Wrote seasons.json, %d bytes", nbytes)
	}

	return nil
}

func LoadSeasonsJSON(seasons_path string) *SeasonIndex {
	season_index := new(SeasonIndex)
	season_index.Seasons = make(map[SeasonID]SeasonMetadata)
	if _, err := os.Stat(seasons_path); os.IsNotExist(err) {
		log.Fatal("file not found ", seasons_path)
	}
	seasons_json, err := os.ReadFile(seasons_path)
	if err != nil {
		log.Println("failed to open seasons.json path", seasons_path)
		log.Fatal(err)
	}
	err = json.Unmarshal(seasons_json, season_index)
	if err != nil {
		log.Fatal("failed to decode seasons metadata ", seasons_path)
	}

	return season_index
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

func (season SeasonID) JSON() string {
	return fmt.Sprintf("%s.json", season)
}
