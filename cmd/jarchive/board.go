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
// github:kevindamm/q-party/cmd/jarchive/board.go

package main

import (
	"log"

	"golang.org/x/net/html"
)

type JArchiveBoard struct {
	ShowDate `json:"-"`
	Round    EpisodeRound          `json:"round"`
	Columns  [6]CategoryChallenges `json:"columns"`
	Wagering []BoardPosition       `json:"wag,omitempty"`
}

func NewBoard(date ShowDate, round EpisodeRound) *JArchiveBoard {
	if round != ROUND_SINGLE_JEOPARDY && round != ROUND_DOUBLE_JEOPARDY {
		panic("attempting to create a new board with invalid round " + round.String())
	}

	single_double := 1
	if round == ROUND_DOUBLE_JEOPARDY {
		single_double = 2
	}

	board := new(JArchiveBoard)
	board.Round = round
	board.Wagering = make([]BoardPosition, 0, single_double)
	return board
}

func (board JArchiveBoard) WageringChallenges() []JArchiveChallenge {
	challenges := make([]JArchiveChallenge, len(board.Wagering))
	for i, position := range board.Wagering {
		challenges[i] = board.Columns[position.Column].Challenges[position.Index]
	}
	return challenges
}

type BoardPosition struct {
	Column int `cue:">=0 & <6"`
	Index  int `cue:">=0 & <5"`
}

func (board *JArchiveBoard) parseBoard(root *html.Node) {
	round_table := nextDescendantWithClass(root, "table", "round")
	category_tr := nextDescendantWithClass(round_table, "tr", "")
	category_tds := childrenWithClass(category_tr, "td", "category")
	if len(category_tds) != 6 {
		log.Fatal("expected 6 category entries, found", len(category_tds))
	}
	for i, td := range category_tds {
		err := board.Columns[i].parseCategoryHeader(td)
		if err != nil {
			log.Fatal("failed to parse category header (name and comments)")
		}
		board.Columns[i].Round = board.Round
	}

	for row := range 5 {
		category_tr = nextSiblingWithClass(category_tr, "tr", "")
		clue_tds := childrenWithClass(category_tr, "td", "clue")
		if len(clue_tds) != 6 {
			log.Fatal("expected 6 clue entries, found", len(clue_tds))
		}
		for i, clue_td := range clue_tds {
			err := board.Columns[i].parseCategoryChallenge(clue_td)
			if err != nil {
				log.Fatal("failed to parse clue entry", err)
			}
			board.Columns[i].Challenges[row].ShowDate = board.ShowDate
		}
	}
}

// enum representation for
type EpisodeRound int

const (
	ROUND_UNKNOWN EpisodeRound = iota
	ROUND_SINGLE_JEOPARDY
	ROUND_DOUBLE_JEOPARDY
	ROUND_FINAL_JEOPARDY
	ROUND_TIE_BREAKER
	ROUND_PRINTED_MEDIA
)

var round_strings = map[EpisodeRound]string{
	ROUND_UNKNOWN:         "[UNKNOWN]",
	ROUND_SINGLE_JEOPARDY: "Jeopardy!",
	ROUND_DOUBLE_JEOPARDY: "Double Jeopardy!",
	ROUND_FINAL_JEOPARDY:  "Final Jeopardy!",
	ROUND_TIE_BREAKER:     "Tiebreaker",
	ROUND_PRINTED_MEDIA:   "[printed media]",
}

func (round EpisodeRound) String() string {
	printed := round_strings[round]
	if printed == "" {
		printed = round_strings[ROUND_UNKNOWN]
	}
	return printed
}

func ParseString(round string) EpisodeRound {
	for k, v := range round_strings {
		if v == round {
			return k
		}
	}
	return ROUND_UNKNOWN
}
