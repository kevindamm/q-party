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
// github:kevindamm/q-party/selfhost/cmd/fetch/jarchive/boards.go

package main

import (
	"log"

	"github.com/kevindamm/q-party/schema"
	"golang.org/x/net/html"
)

type JarchiveBoard struct {
	schema.MatchNumber
	schema.RoundEnum
	Columns []JarchiveCategory
}

func ParseBoard(root *html.Node) *JarchiveBoard {
	board := new(JarchiveBoard)
	board.Columns = make([]JarchiveCategory, 6)
	round_table := nextDescendantWithClass(root, "table", "round")
	category_tr := nextDescendantWithClass(round_table, "tr", "")
	category_tds := childrenWithClass(category_tr, "td", "category")
	if len(category_tds) != 6 {
		log.Fatal("expected 6 category entries, found", len(category_tds))
	}
	for i, td := range category_tds {
		table := nextDescendantWithClass(td, "table", "")
		tbody := nextDescendantWithClass(table, "tbody", "")
		trs := childrenWithClass(tbody, "tr", "")
		if len(trs) != 2 {
			log.Fatal("length of trs expected 2 but have ", len(trs))
		}
		title, err := innerTextMarkdown(
			nextDescendantWithClass(trs[0], "td", "category_name"))
		if err != nil {
			log.Fatal("failed to parse category name", err)
		}
		comments := innerText(
			nextDescendantWithClass(trs[1], "td", "category_comments"))
		if title == "" {
			// Sometimes the category comment appears before the category name.
			title, err = innerTextMarkdown(
				nextDescendantWithClass(trs[1], "td", "category_name"))
			if err != nil {
				log.Fatal("failed to parse category name (2x)", err)
			}
			comments = innerText(
				nextDescendantWithClass(trs[0], "td", "category_comments"))
		}

		if title == "" {
			log.Fatal("FOUND EMPTY CATEGORY TITLE")
		}

		board.Columns[i].Title = schema.CategoryName(title)
		board.Columns[i].Comments = comments
	}

	row_count := 5
	for range row_count { // for each row of the board.
		category_tr = nextSiblingWithClass(category_tr, "tr", "")
		clue_tds := childrenWithClass(category_tr, "td", "clue")
		if len(clue_tds) != 6 {
			log.Fatal("expected 6 clue entries, found", len(clue_tds))
		}

		for i, clue_td := range clue_tds { // for each category column.
			column := board.Columns[i]
			nextChallenge, err := ParseChallenge(clue_td, column.Title)
			if err != nil {
				log.Fatal("failed to parse clue entry\n", err)
			}
			column.Challenges = append(column.Challenges, nextChallenge)
		}
	}
	return board
}
