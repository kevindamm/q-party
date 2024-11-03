package main

import (
	"errors"
	"strings"

	"golang.org/x/net/html"
)

// The de-normed representation as found in some datasets, e.g. on Kaggle.
type JArchiveChallenge struct {
	Round      EpisodeRound `json:"round"`
	Category   string       `json:"category"`           // ALL \"U\"PPERCASE
	AirDate    *AirDate     `json:"air_date,omitempty"` // YYYY-MM-DD
	Value      DollarValue  `json:"value,omitempty"`    // /$(\d+)/
	Commentary string       `json:"comment,omitempty"`

	Prompt  string   `json:"prompt"`
	Correct string   `json:"correct"` // excluding "what is..." preface
	Accept  []string `json:"accept,omitempty"`
}

func (challenge JArchiveChallenge) IsEmpty() bool {
	if challenge.Category != "" || challenge.Commentary != "" ||
		challenge.Round != ROUND_UNKNOWN || challenge.Value != 0 ||
		challenge.Prompt != "" || challenge.Correct != "" ||
		len(challenge.Accept) != 0 {
		return false
	}
	return true
}

type JArchiveFinalChallenge JArchiveChallenge

func (challenge *JArchiveChallenge) parseChallenge(div *html.Node) error {
	if strings.Trim(innerText(div), " ") == "" {
		return nil
	}

	table := nextDescendantWithClass(div, "table", "")

	var err error
	challenge.Value, err = ParseDollarValue(innerText(
		nextDescendantWithClass(table, "td", "clue_value")))
	if err != nil {
		return errors.New("failed to parse challenge value " + err.Error())
	}

	clue := nextDescendantWithClass(table, "td", "clue_text")
	challenge.Category = innerText(clue)
	clue = nextSiblingWithClass(clue, "td", "clue_text")
	if clue == nil {
		return errors.New("could not find challenge response")
	}
	clue = nextDescendantWithClass(clue, "em", "correct_response")
	if clue == nil {
		return errors.New("could not find correct response")
	}
	challenge.Correct = innerText(clue)
	return nil
}

func (final *JArchiveFinalChallenge) parseChallenge(div *html.Node) error {
	table := nextDescendantWithClass(div, "table", "")

	final.Round = ROUND_FINAL_JEOPARDY
	final.Category = innerText(
		nextDescendantWithClass(table, "td", "category_name"))
	final.Commentary = innerText(
		nextDescendantWithClass(table, "td", "category_comments"))

	clue := nextDescendantWithClass(div, "td", "clue")
	clue_td := nextDescendantWithClass(clue, "td", "clue_text")
	final.Prompt = innerText(clue_td)
	clue_td = nextSiblingWithClass(clue_td, "td", "clue_text")
	if clue_td == nil {
		return errors.New("could not find final challenge response")
	}
	final.Correct = innerText(nextDescendantWithClass(clue_td, "em", "correct_response"))
	return nil
}
