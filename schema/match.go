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
)

// go:embed match.cue
var schemaMatches string

// A match identifier refers to the unique identifier of the ?-Party database.
// These are not universally unique, they are only certain to be locally unique.
type MatchNumber uint64

// Shows are numbered sequentially based on air date.
type MatchID struct {
	MatchNumber `json:"match"`
	ShowTitle   string     `json:"show_title,omitempty"`
	SeasonSlug  SeasonSlug `json:"season,omitempty"`
}

// Convenience function for constructing the above, when not querying from DB.
func NewMatchID(match_number int) MatchID {
	return MatchID{
		MatchNumber: MatchNumber(uint64(match_number)),
	}
}

// Data about the episode that is unrelated to the match, rounds or challenges.
type MatchMetadata struct {
	MatchID `json:",inline"`

	AiredDate *ShowDate `json:"aired,omitempty"`
	TapedDate *ShowDate `json:"taped,omitempty"`

	Contestants []ContestantID `json:"contestants,omitempty"`
	Media       []MediaRef     `json:"media,omitempty"`
	Comments    string         `json:"comments,omitempty"`
}

type MatchIndex map[MatchNumber]*MatchMetadata

// Will add (and possibly overwrite) the values from metadata into the mapping.
func (episodes MatchIndex) Update(metadata MatchMetadata) {
	match := metadata.MatchNumber
	episode, exists := episodes[match]
	if !exists {
		episodes[match] = &metadata
		return
	}

	if len(metadata.SeasonSlug) > 0 {
		episode.SeasonSlug = metadata.SeasonSlug
	}
	if len(metadata.ShowTitle) > 0 {
		episode.ShowTitle = metadata.ShowTitle
	}

	if metadata.AiredDate.String() != "" {
		episode.AiredDate = metadata.AiredDate
	}
	if metadata.TapedDate.String() != "" {
		episode.TapedDate = metadata.AiredDate
	}

	// Sequence types are added to, not replaced with.  TODO merge lists
	if len(metadata.Contestants) > 0 {
		episode.Contestants = append(episode.Contestants, metadata.Contestants...)
	}
	if len(metadata.Media) > 0 {
		episode.Media = append(episode.Media, metadata.Media...)
	}

	if len(metadata.Comments) > 0 {
		if len(episode.Comments) == 0 {
			episode.Comments = metadata.Comments
		} else {
			episode.Comments = episode.Comments + "\n\n" + metadata.Comments
		}
	}
}

type BoardLayout struct {
	CategoryBitmaps []uint `json:"cat_bitmap"`
}

type MatchStats struct {
	MatchMetadata `json:",inline"`
	SingleCount   int `json:"single_count,omitempty"`
	DoubleCount   int `json:"double_count,omitempty"`

	TripleStumpers []BoardPosition `json:"triple_stumpers,omitempty"`
}
