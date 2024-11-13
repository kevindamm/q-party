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
// github:kevindamm/q-party/json/board.go

package qparty

type Board struct {
	ShowNumber `json:"episode"`
	Round      EpisodeRound `json:"round" cue:"<len(round_names)"`

	Columns []Category  `json:"columns"`
	Missing []Position  `json:"missing,omitempty"`
	History []Selection `json:"history,omitempty"`
}

func (board Board) RoundName() string {
	return round_names[board.Round]
}

type Position struct {
	Column uint `json:"column" cue:"<6"`
	Index  uint `json:"index" cue:"<5"`
}

type Selection struct {
	Position          `json:",inline"`
	ChallengeMetadata `json:",inline"`
}

// An enum-like value for the different rounds.
type EpisodeRound uint

func (round EpisodeRound) String() string {
	if int(round) >= len(round_names) {
		return round_names[0]
	}
	return round_names[round]
}

const (
	ROUND_UNKNOWN EpisodeRound = iota
	ROUND_SINGLE
	ROUND_DOUBLE
	ROUND_FINAL
	ROUND_TIEBREAKER
	PRINTED_MEDIA
)

var round_names = [6]string{
	"[UNKNOWN]",
	"Single!",
	"Double!",
	"Final!",
	"Tiebreaker!!",
	"[printed media]"}
