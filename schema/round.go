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
// github:kevindamm/q-party/schema/rounds.go

package schema

import (
	_ "embed"
	"fmt"
)

// go:embed round.cue
var schemaRounds string

type RoundID struct {
	Episode MatchNumber `json:"episode,omitempty"`
	Round   RoundEnum   `json:"round,omitempty"`
}

// An enum-like value for the different rounds.
type RoundEnum int

const (
	ROUND_UNKNOWN RoundEnum = iota
	ROUND_SINGLE
	ROUND_DOUBLE
	ROUND_FINAL
	ROUND_TIEBREAKER
	PRINTED_MEDIA
	MaxRoundEnum
)

func (round RoundID) String() string {
	round_name := round_names[0]
	if round.Round < MaxRoundEnum {
		round_name = round_names[round.Round]
	}
	return fmt.Sprintf("#%05d: %s", round.Episode, round_name)
}

func (round RoundID) RoundName() string {
	if round.Round < 0 || round.Round >= MaxRoundEnum {
		return round_names[0]
	}
	return round_names[round.Round]
}

var round_names = [6]string{
	"[UNKNOWN]",
	"Single!",
	"Double!",
	"Final!",
	"Tiebreaker!!",
	"[printed media]"}

type Board struct {
	RoundID `json:",inline"`
	Columns []CategoryMetadata `json:"columns"`
	Missing []BoardPosition    `json:"missing,omitempty"`
}

type BoardState struct {
	Board `json:",inline"`

	Layout  BoardLayout        `json:"cat_bitmap"`
	History []SelectionOutcome `json:"history"`
}

type BoardLayout []byte

type BoardPosition struct {
	Column uint `json:"column"`
	Index  uint `json:"index"`
}

type BoardSelection struct {
	ChallengeMetadata `json:",inline"`
	BoardPosition     `json:",inline"`
}

type SelectionOutcome struct {
	BoardSelection `json:",inline"`
	Correct        bool  `json:"correct"`
	Delta          Value `json:"delta"`
}
