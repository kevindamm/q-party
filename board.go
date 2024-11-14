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
// github:kevindamm/q-party/board.go

package qparty

// The basic board definition (contains only category metadata).
type Board struct {
	BoardID `json:",inline"`
	Columns []Category `json:"columns"`
	Missing []Position `json:"-,omitempty"`
}

// Extends the board definition with a history of challenge selections.
type BoardState struct {
	Board   `json:",inline"`
	History []Selection `json:"history,omitempty"`
}

// The host's view of a board includes the details, including correct response.
type HostBoard struct {
	BoardID `json:",inline"`
	Columns []HostCategory `json:"columns"`
}

// A board is identified by its episode and whether it's single or double round.
type BoardID struct {
	Episode ShowNumber   `json:"episode,omitempty"`
	Round   EpisodeRound `json:"round,omitempty"` // cue:"<len(round_names)"
}

func (board BoardID) RoundName() string {
	return round_names[board.Round]
}

// A board position, located by the column and (descending) index in the column.
type Position struct {
	Column uint `json:"column" cue:"<6"`
	Index  uint `json:"index" cue:"<5"`
}

// A contestant's selection (including possibly the wager value).
type Selection struct {
	ContestantIndex   uint `json:"player"` // \in { 0, 1, 2 }
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
