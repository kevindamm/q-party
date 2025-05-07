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
// github:kevindamm/q-party/cmd/fetch/jarchive/final.go

package main

import (
	"errors"

	"github.com/kevindamm/q-party/schema"
	"golang.org/x/net/html"
)

// Full representation, as with the above, as well as each player's wager.
type JarchiveFinalChallenge struct {
	schema.Challenge
	Wagers    []schema.PlayerWager
	Responses []string
	Stumped   int // 0..|players|

	Correct []string
}

// There may not be a final jeopardy (or it may not have been entered yet).
func UnknownFinal() JarchiveFinalChallenge {
	return JarchiveFinalChallenge{
		schema.UnknownChallenge(),
		[]schema.PlayerWager{},
		[]string{},
		0,
		[]string{},
	}
}

func ParseFinalChallenge(div *html.Node) (*JarchiveFinalChallenge, error) {
	table := nextDescendantWithClass(div, "table", "")

	final := new(JarchiveFinalChallenge)
	final.Category = schema.CategoryName(innerText(
		nextDescendantWithClass(table, "td", "category_name")))
	final.Comments = innerText(
		nextDescendantWithClass(table, "td", "category_comments"))

	clue := nextDescendantWithClass(div, "td", "clue")
	clue_td := nextDescendantWithClass(clue, "td", "clue_text")
	final.Clue = innerText(clue_td)
	clue_td = nextSiblingWithClass(clue_td, "td", "clue_text")
	if clue_td == nil {
		return nil, errors.New("could not find final challenge response")
	}
	final.Correct = []string{
		innerText(nextDescendantWithClass(clue_td, "em", "correct_response"))}
	return final, nil
}

type JarchiveTiebreaker struct {
	schema.Challenge
	Contestants []schema.ContestantID
	Responses   []string
	Stumped     int // 0..|players|

	Correct []string // acceptable answer(s), typically one
}

// There may not be a final jeopardy (or it may not have been entered yet).
func UnknownTiebreaker() JarchiveTiebreaker {
	return JarchiveTiebreaker{
		schema.UnknownChallenge(),
		[]schema.ContestantID{},
		[]string{},
		0,
		[]string{},
	}
}

// A tiebreaker is formatted similarly to the final, but without additional wagers.
func ParseTiebreakerChallenge(div *html.Node) (*JarchiveTiebreaker, error) {
	table := nextDescendantWithClass(div, "table", "")
	tiebreaker := new(JarchiveTiebreaker)
	tiebreaker.Category = schema.CategoryName(innerText(
		nextDescendantWithClass(table, "td", "category_name")))
	tiebreaker.Comments = innerText(
		nextDescendantWithClass(table, "td", "category_comments"))

	clue := nextDescendantWithClass(div, "td", "clue")
	clue_td := nextDescendantWithClass(clue, "td", "clue_text")
	tiebreaker.Clue = innerText(clue_td)
	clue_td = nextSiblingWithClass(clue_td, "td", "clue_text")
	if clue_td == nil {
		return nil, errors.New("could not find tiebreaker challenge response")
	}
	tiebreaker.Correct = []string{innerText(
		nextDescendantWithClass(clue_td, "em", "correct_response"))}
	return tiebreaker, nil
}
