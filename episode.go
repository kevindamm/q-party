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
// github:kevindamm/q-party/json/episodes.go

package qparty

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

type EpisodeMetadata struct {
	ShowNumber `json:"show_number"` // cue:">0"
	SeasonID   `json:"season,omitempty"`
	Aired      ShowDate `json:"aired,omitempty"`
	Taped      ShowDate `json:"taped,omitempty"`

	ShowTitle     string          `json:"show_title"`
	ContestantIDs [3]ContestantID `json:"contestant_ids,omitempty"`
}

type EpisodeStats struct {
	EpisodeMetadata `json:",inline"`

	SingleClues int          `json:"single_count"`
	DoubleClues int          `json:"double_count"`
	Stumpers    [][]Position `json:"triple_stumpers"`
}

type Episode struct {
	EpisodeMetadata `json:",inline"`
	Comments        string  `json:"comments,omitempty"`
	Media           []Media `json:"media,omitempty"`

	Single     *Board     `json:"single,omitempty"`
	Double     *Board     `json:"double,omitempty"`
	Final      *Challenge `json:"final,omitempty"`
	TieBreaker *Challenge `json:"tiebreaker,omitempty"`
}

// Unique numeric identifier for episodes in the archive.
// May be different than the sequential show number used in display.
type ShowNumber uint

func (id ShowNumber) String() string {
	return fmt.Sprintf("%d", uint(id))
}

func (id ShowNumber) HTML() string {
	return fmt.Sprintf("episodes/%d.html", id)
}

func (id ShowNumber) JSON() string {
	return fmt.Sprintf("episodes/%d.json", id)
}

// Parses the numeric value from a string.
// Fatal error if the value cannot be converted into a number.
func MustParseShowNumber(numeric string) ShowNumber {
	id, err := strconv.Atoi(numeric)
	if err != nil {
		log.Fatalf("failed to parse JEID from string '%s'\n%s", numeric, err)
	}
	return ShowNumber(id)
}

func (episode EpisodeMetadata) WriteJSON(dir string) error {
	episode_json, err := json.MarshalIndent(episode, "", "  ")
	if err != nil {
		return err
	}

	filepath := path.Join(dir, episode.ShowNumber.JSON())
	log.Println("converting episode", filepath)
	writer, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer writer.Close()

	nbytes, err := writer.Write(episode_json)
	if err != nil {
		return err
	}

	log.Println(nbytes, "bytes written")
	return nil
}
