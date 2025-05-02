// Copyright (c) 2025 Kevin Damm
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
// github:kevindamm/q-party/schema/matches.go

package schema

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

// go:embed match.cue
var schemaMatches string

// A match identifier refers to the unique identifier of the ?-Party database.
// These are not universally unique, they are only certain to be locally unique.
type MatchNumber uint64

// Shows are numbered sequentially based on air date.
type MatchID struct {
	Match     MatchNumber `json:"match"`
	ShowTitle string      `json:"show_title,omitempty"`

	SeasonSlug SeasonSlug `json:"season,omitempty"`
}

// Convenience function for constructing the above, when not querying from DB.
func NewMatchID(match_number int) MatchID {
	return MatchID{
		Match: MatchNumber(uint64(match_number)),
	}
}

type EpisodeMetadata struct {
	MatchID   `json:",inline"`
	AiredDate *ShowDate `json:"aired,omitempty"`
	TapedDate *ShowDate `json:"taped,omitempty"`

	Contestants []ContestantID `json:"contestants,omitempty"`
	Media       []MediaRef     `json:"media,omitempty"`
	Comments    string         `json:"comments,omitempty"`
}

type EpisodeIndex map[MatchNumber]*EpisodeMetadata

// Will add (and possibly overwrite) the values from metadata into the mapping.
func (episodes EpisodeIndex) Update(metadata EpisodeMetadata) {
	match := metadata.Match
	existing, ok := episodes[match]
	if !ok {
		episodes[match] = &metadata
		return
	}

	if len(metadata.SeasonSlug) > 0 {
		existing.SeasonSlug = metadata.SeasonSlug
	}
	if len(metadata.ShowTitle) > 0 {
		existing.ShowTitle = metadata.ShowTitle
	}

	if metadata.AiredDate.String() != "" {
		existing.AiredDate = metadata.AiredDate
	}
	if metadata.TapedDate.String() != "" {
		existing.TapedDate = metadata.AiredDate
	}

	// Sequence types are added to, not replaced with.
	if len(metadata.Contestants) > 0 {
		existing.Contestants = append(existing.Contestants, metadata.Contestants...)
	}
	if len(metadata.Media) > 0 {
		existing.Media = append(existing.Media, metadata.Media...)
	}
	if len(metadata.Comments) > 0 {
		if len(existing.Comments) == 0 {
			existing.Comments = metadata.Comments
		} else {
			existing.Comments = existing.Comments + "\n\n" + metadata.Comments
		}
	}
}

type BoardLayout struct {
	CategoryBitmaps []uint `json:"cat_bitmap"`
}

type EpisodeStats struct {
	SingleCount int `json:"single_count,omitempty"`
	DoubleCount int `json:"double_count,omitempty"`

	TripleStumpers []BoardPosition `json:"triple_stumpers,omitempty"`
}

type ShowDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// formatted as YYYY-MM-DD
func ParseShowDate(image string) *ShowDate {
	reShowDate := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	matches := reShowDate.FindStringSubmatch(image)
	// We can ignore the error here because we know these are all digit patterns.
	year, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[3])

	return &ShowDate{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func (sd ShowDate) String() string {
	if sd.Year+sd.Month+sd.Day == 0 {
		return ""
	}
	return fmt.Sprintf("%04d/%02d/%02d", sd.Year, sd.Month, sd.Day)
}

// Returns 0 if `this` and `other` are equal;
// +1 if this is later than other, and -1 if before.
func (this ShowDate) Compare(other *ShowDate) int {
	if this.Year < other.Year {
		return -1
	}
	if this.Year > other.Year {
		return +1
	}
	// (this.Year == other.Year)
	if this.Month < other.Month {
		return -1
	}
	if this.Month > other.Month {
		return +1
	}
	// (this.Month == other.Month)
	if this.Day < other.Day {
		return -1
	}
	if this.Day > other.Day {
		return +1
	}
	// The dates are equal.
	return 0
}

type ShowDateRange struct {
	From  *ShowDate `json:"from,omitempty"`
	Until *ShowDate `json:"until,omitempty"`
}

func (scope ShowDateRange) Contains(date ShowDate) bool {
	return (                         // including endpoints,
	date.Compare(scope.From) >= 0 && // after beginning and
		date.Compare(scope.Until) <= 0) // before ending
}
