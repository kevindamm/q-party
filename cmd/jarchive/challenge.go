package main

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// The de-normed representation as found in some datasets, e.g. on Kaggle.
type JArchiveChallenge struct {
	ShowDate   `json:"air_date,omitempty"`
	Round      EpisodeRound     `json:"round"`
	Category   JArchiveCategory `json:"category"`
	Commentary string           `json:"comment,omitempty"`

	// String representation has a dollar sign, negated values are wagers.
	Value DollarValue `json:"value,omitempty"`

	Prompt  string   `json:"prompt"`
	Correct string   `json:"correct"` // excluding "what is..." preface
	Accept  []string `json:"accept,omitempty"`
}

var unknown_challenge = JArchiveChallenge{}

func (challenge JArchiveChallenge) IsEmpty() bool {
	if challenge.Value != 0 ||
		challenge.Prompt != "" ||
		challenge.Correct != "" ||
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
	value_td := nextDescendantWithClass(table, "td", "clue_value")
	if value_td != nil {
		challenge.Value, err = ParseDollarValue(innerText(value_td))
		if err != nil {
			return errors.New("failed to parse challenge value " + err.Error())
		}
	} else {
		dd_value_td := nextDescendantWithClass(table, "td", "clue_value_daily_double")
		if dd_value_td != nil {
			text := strings.ReplaceAll(innerText(dd_value_td), ",", "")
			challenge.Value, err = ParseDollarValue(text[4:])
			if err != nil {
				return fmt.Errorf("failed to parse daily double value %s\n%s", text, err.Error())
			}
			challenge.Value = -challenge.Value
		}
	}

	clue := nextDescendantWithClass(table, "td", "clue_text")
	challenge.Prompt = innerText(clue)
	clue = nextSiblingWithClass(clue, "td", "clue_text")
	if clue == nil {
		return errors.New("could not find challenge response")
	}
	correct := nextDescendantWithClass(clue, "em", "correct_response")
	if correct == nil {
		return errors.New("could not find correct response")
	}
	challenge.Correct = innerText(correct)
	return nil
}

func (final *JArchiveFinalChallenge) parseChallenge(div *html.Node) error {
	table := nextDescendantWithClass(div, "table", "")

	final.Round = ROUND_FINAL_JEOPARDY
	final.Category = JArchiveCategory(innerText(
		nextDescendantWithClass(table, "td", "category_name")))
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
