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

package json

type EpisodeID struct {
	JEID       `json:"jeid" cue:">0"`
	ShowNumber int `json:"show_number,omitempty" cue:">=0"`
}

type JEID int

type EpisodeMetadata struct {
	EpisodeID `json:",inline"`
	SeasonID  `json:"season"`
	Aired     ShowDate `json:"aired,omitempty"`

	ContestantIDs [3]int  `json:"contestant_ids"`
	Comments      string  `json:"comments,omitempty"`
	Media         []Media `json:"media,omitempty"`

	SingleClues    int `json:"single_count"`
	DoubleClues    int `json:"double_count"`
	TripleStumpers int `json:"triple_stumpers"`
}

type ShowDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}
