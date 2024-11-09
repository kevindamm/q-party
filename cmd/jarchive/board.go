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

	"github.com/kevindamm/q-party/json"
	"golang.org/x/net/html"
)

func parseBoard(root *html.Node) [6]CategoryChallenges {
	categories := [6]CategoryChallenges{}
	round_table := nextDescendantWithClass(root, "table", "round")
	category_tr := nextDescendantWithClass(round_table, "tr", "")
	category_tds := childrenWithClass(category_tr, "td", "category")
	if len(category_tds) != 6 {
		log.Fatal("expected 6 category entries, found", len(category_tds))
	}
	for i, td := range category_tds {
		err := parseCategoryHeader(td, &categories[i])
		if err != nil {
			log.Fatal("failed to parse category header (name and comments)")
		}
	}

	for range 5 {
		category_tr = nextSiblingWithClass(category_tr, "tr", "")
		clue_tds := childrenWithClass(category_tr, "td", "clue")
		if len(clue_tds) != 6 {
			log.Fatal("expected 6 clue entries, found", len(clue_tds))
		}
		for i, clue_td := range clue_tds {
			err := parseCategoryChallenge(clue_td, &categories[i])
			if err != nil {
				log.Fatal("failed to parse clue entry", err)
			}
		}
	}
	return categories
}

var round_strings = map[json.EpisodeRound]string{
	json.ROUND_UNKNOWN:    "[UNKNOWN]",
	json.ROUND_SINGLE:     "Jeopardy!",
	json.ROUND_DOUBLE:     "Double Jeopardy!",
	json.ROUND_FINAL:      "Final Jeopardy!",
	json.ROUND_TIEBREAKER: "Tiebreaker",
	json.PRINTED_MEDIA:    "[printed media]",
}

func ParseString(round string) json.EpisodeRound {
	if round == "" {
		return json.ROUND_UNKNOWN
	}
	for k, v := range round_strings {
		if v == round {
			return k
		}
	}
	return json.ROUND_UNKNOWN
}
