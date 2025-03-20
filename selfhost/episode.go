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
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
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

// Parses the show number from its string representation.
func ParseShowNumber(show_id string) (*ShowNumber, error) {
	parts := strings.Split(show_id, "/")
	if len(parts) != 2 {
		return nil, errors.New("expected format (season '/' number)")
	}
	// TODO verify season name
	// TODO verify episode in season range
	show := ShowNumber{
		Season: SeasonID(parts[0]),
		Number: must_parse_uint(parts[1])}

	return &show, nil
}

// Show numbers are unique within regular seasons but a new sequence is created
// for non-regular seasons also.  Thus the JSON name uses both season and show
// number, even though this leads to some number-hyphen-number naming.
func (show ShowNumber) JSON() string {
	return fmt.Sprintf("%s-%d.json", string(show.Season), show.Number)
}

// The display string representation of this show number.
func (show ShowNumber) String() string {
	return fmt.Sprintf("%s/%d", show.Season.prefix(), show.Number)
}

// Parses the numeric value from a string.
// Fatal error if the value cannot be converted into a number.
func must_parse_uint(numeric string) uint {
	id, err := strconv.Atoi(numeric)
	if err != nil || id < 1 {
		log.Fatalf("expected positive integer, got '%s'", numeric)
	}
	return uint(id)
}

// Unique ID which J-Archive uses to identify its episodes.
type EpisodeID uint

// Only the HTML filename and not the JSON filename can be inferred from the ID.
func (id EpisodeID) HTML() string {
	return fmt.Sprintf("%d.html", uint(id))
}

type ShowDateRange struct {
	From  ShowDate `json:"from"`
	Until ShowDate `json:"until,omitempty"`
}

type ShowDate struct {
	Year  int `json:"year"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

func (date ShowDate) IsZero() bool {
	return date.Year == 0 && date.Month == 0 && date.Day == 0
}

func (date ShowDate) Equal(other ShowDate) bool {
	return other.Year == date.Year &&
		other.Month == date.Month &&
		other.Day == date.Day
}

func (date ShowDate) ToTime() time.Time {
	return time.Date(
		date.Year, time.Month(date.Month), date.Day,
		23, 7, 42, 0, time.UTC)
}

func (date ShowDate) String() string {
	if date.IsZero() {
		return ""
	}
	if date.Month != 0 && date.Day != 0 {
		return fmt.Sprintf("%d/%02d/%02d",
			date.Year, date.Month, date.Day)
	}

	month := "??"
	if date.Month != 0 {
		month = fmt.Sprintf("%02d", date.Month)
	}
	day := "??"
	if date.Day != 0 {
		day = fmt.Sprintf("%02d", date.Day)
	}
	return fmt.Sprintf("%d/%s/%s", date.Year, month, day)
}

func (date ShowDate) MarshalText() ([]byte, error) {
	return []byte(date.String()), nil
}

func (date *ShowDate) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*date = ShowDate{0, 0, 0}
		return nil
	}
	if len(text) < len("YYYY") || len(text) > len("YYYY_MM_DD") {
		return fmt.Errorf("incorrect format for aired date '%s', use YYYY/MM/DD having at least the year", text)
	}

	year, err := strconv.Atoi(string(text[:4]))
	if err != nil {
		return err
	}
	var month int
	if len(text) > 4 {
		month, err = strconv.Atoi(string(text[5:7]))
		if err != nil {
			month = 0
		}
	}
	var day int
	if len(text) > 6 {
		day, err = strconv.Atoi(string(text[8:]))
		if err != nil {
			day = 0
		}
	}

	*date = ShowDate{year, month, day}
	return nil
}
