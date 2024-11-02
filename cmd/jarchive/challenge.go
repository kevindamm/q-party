package main

import (
	"fmt"

	"golang.org/x/net/html"
)

// The de-normed representation as found in some datasets, e.g. on Kaggle.
type JArchiveChallenge struct {
	Round      EpisodeRound `json:"round"`
	Category   string       `json:"category"`           // ALL \"U\"PPERCASE
	AirDate    *AirDate     `json:"air_date,omitempty"` // YYYY-MM-DD
	Commentary string       `json:"comment,omitempty"`

	Value   DollarValue `json:"value"` // /$(\d+)/ use negative value for wagers
	Prompt  string      `json:"prompt"`
	Correct string      `json:"correct"` // excluding "what is..." preface
	Accept  []string    `json:"accept,omitempty"`
}

type JArchiveFinalChallenge JArchiveChallenge

func (challenge *JArchiveChallenge) parseChallenge(div *html.Node) {
	// TODO
	fmt.Println("TODO parse challenge")
}

func (final *JArchiveFinalChallenge) parseChallenge(div *html.Node) {
	table := nextDescendantWithClass(div, "table", "")

	final.Round = ROUND_FINAL_JEOPARDY
	final.Category = innerText(
		nextDescendantWithClass(table, "td", "category_name"))
	final.Commentary = innerText(
		nextDescendantWithClass(table, "td", "category_comments"))

	clue := nextDescendantWithClass(div, "td", "clue")
	clue_td := nextDescendantWithClass(clue, "td", "clue_text")
	final.Prompt = innerText(clue_td)
	clue_td = clue_td.NextSibling
	for clue_td != nil && clue_td.Type != html.ElementNode && clue_td.Data != "td" {
		clue_td = clue_td.NextSibling
	}
	if clue_td == nil {
		panic("could not find response")
	}
	final.Correct = innerText(nextDescendantWithClass(clue_td, "em", "correct_response"))
}
