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
// github:kevindamm/q-party/episodes.go

package qparty

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type EpisodeMetadata struct {
	Show      ShowNumber `json:"show_number"`
	ShowTitle string     `json:"show_title"`

	EpisodeID `json:"episode_id,omitempty"`
	Aired     ShowDate `json:"aired,omitempty"`
	Taped     ShowDate `json:"taped,omitempty"`
}

type EpisodeStats struct {
	EpisodeMetadata `json:",inline"`

	SingleCount int          `json:"single_count"`
	DoubleCount int          `json:"double_count"`
	Stumpers    [][]Position `json:"triple_stumpers"`
}

// All details of the episode, including correct answers & the contestants' bios.
type FullEpisode struct {
	EpisodeMetadata `json:",inline"`
	Comments        string  `json:"comments,omitempty"`
	Media           []Media `json:"media,omitempty"`

	// Due to absence of archival evidence, not every episode has both single & double rounds.
	Single     *FullBoard     `json:"single,omitempty"`
	Double     *FullBoard     `json:"double,omitempty"`
	Final      *FullChallenge `json:"final,omitempty"`
	TieBreaker *FullChallenge `json:"tiebreaker,omitempty"`
}

// Unique numeric identifier for episodes in the archive.
// May be different than the sequential show number used in display.
type ShowNumber struct {
	Season SeasonID // empty string indicates regular-season play.
	Number uint     // cue:">0"
}

func (show *ShowNumber) UnmarshalJSON(json_bytes []byte) error {
	var intVal uint
	if err := json.Unmarshal(json_bytes, &intVal); err != nil {
		show.Number = intVal
	} else {
		var stringVal string
		if err := json.Unmarshal(json_bytes, &stringVal); err != nil {
			number, err := strconv.Atoi(stringVal)
			if err != nil {
				return err
			}
			if number < 1 {
				return fmt.Errorf("expected positive integer for show number")
			}
			show.Number = uint(number)
		}
	}
	return nil
}

// Show numbers are unique within regular seasons but a new sequence is created
// for non-regular seasons also.  Thus the JSON name uses both season and show
// number, even though this leads to some number-hyphen-number naming.
func (show ShowNumber) JSON() string {
	return fmt.Sprintf("%s-%d.json", string(show.Season), show.Number)
}

// Parses the numeric value from a string.
// Fatal error if the value cannot be converted into a number.
func MustParseShowNumber(numeric string) ShowNumber {
	id, err := strconv.Atoi(numeric)
	if err != nil {
		log.Fatalf("failed to parse Show Number from string '%s'\n%s", numeric, err)
	}
	if id < 1 {
		log.Fatalf("expected positive integer for show number, got '%s'", numeric)
	}
	return ShowNumber{"", uint(id)}
}

// Unique ID which J-Archive uses to identify its episodes.
type EpisodeID uint

// Only the HTML filename and not the JSON filename can be inferred from the ID.
func (id EpisodeID) HTML() string {
	return fmt.Sprintf("%d.html", uint(id))
}
