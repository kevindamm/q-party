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
// github:kevindamm/q-party/selfhost/cmd/fetch/jarchive/challenges.go

package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kevindamm/q-party/schema"
	"golang.org/x/net/html"
)

// Full representation of the challenge, including the correct response.
type JarchiveChallenge struct {
	schema.Challenge
	Responses []string
	Correct   []string
	Stumped   int // 0..3 how many incorrect responses were given (not counting silence)
}

// Full representation, as with the above, as well as each player's wager.
type JarchiveFinal struct {
	schema.Challenge
	Wagers    []schema.PlayerWager
	Responses []string
	Correct   []string
	Stumped   int // 0..|players|
}

// Sentinel value for board entries that are missing/blank.
func UnknownChallenge() JarchiveChallenge {
	return JarchiveChallenge{
		schema.UnknownChallenge(),
		[]string{},
		[]string{},
		0}
}

// There may not be a final jeopardy (or it may not have been entered yet).
func UnknownFinal() JarchiveFinal {
	return JarchiveFinal{
		schema.UnknownChallenge(),
		[]schema.PlayerWager{},
		[]string{},
		[]string{},
		0}
}

// Assumes that a file extension is present.
var reMediaPath = regexp.MustCompile(`^https?://.*\.com/media/([^.]+)(\.[a-zA-Z0-9]+)`)
var reCluePath = regexp.MustCompile(`^suggestcorrection\.php\?clue_id=(\d+)`)

func MakeMediaClue(href string) schema.MediaRef {
	match := reMediaPath.FindStringSubmatch(href)
	if match == nil {
		return schema.MediaRef{}
	}
	filename := fmt.Sprintf("/media/%s%s", match[1], match[2])
	var mimetype schema.MimeType
	switch match[2] {
	case ".jpg", ".jpeg":
		mimetype = schema.MediaImageJPG
	case ".png":
		mimetype = schema.MediaImagePNG
	case ".svg":
		mimetype = schema.MediaImageSVG
	case ".mp3":
		mimetype = schema.MediaAudioMP3
	case ".mp4":
		mimetype = schema.MediaVideoMP4
	case ".mov":
		mimetype = schema.MediaVideoMOV
	default:
		panic("unrecognized media type for " + match[2])
	}

	return schema.MediaRef{
		MimeType: mimetype,
		MediaURL: filename}
}

func NewChallenge(category string) *JarchiveChallenge {
	challenge := new(JarchiveChallenge)
	challenge.Category = category
	challenge.Media = make([]schema.MediaRef, 0)
	return challenge
}

func parseChallenge(div *html.Node, challenge *JarchiveChallenge) error {
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

	td_order_number := nextDescendantWithClass(table, "td", "clue_order_number")
	edit_link := nextDescendantWithClass(td_order_number, "a", "")
	if edit_link != nil {
		for _, attr := range edit_link.Attr {
			if attr.Key == "href" {
				match := reCluePath.FindStringSubmatch(attr.Val)
				if match != nil {
					// We know from the regex that this is an integer.
					clue_id, _ := strconv.Atoi(match[1])
					challenge.ChallengeID = schema.ChallengeID(clue_id)
				}
			}
		}
	}

	clue_td := nextDescendantWithClass(table, "td", "clue_text")
	text, media := parseIntoMarkdown(clue_td)
	challenge.Clue = text
	if len(media) > 0 {
		challenge.Media = media
	}

	if err != nil {
		return fmt.Errorf("failed to parse challenge prompt %s...\n%s",
			text[:18], err.Error())
	}
	clue_td = nextSiblingWithClass(clue_td, "td", "clue_text")
	if clue_td == nil {
		return errors.New("could not find challenge response")
	}
	correct := nextDescendantWithClass(clue_td, "em", "correct_response")
	if correct == nil {
		return errors.New("could not find correct response")
	}
	challenge.Correct = []string{innerText(correct)}

	judgement := nextDescendantWithClass(clue_td, "table", "")
	if innerText(judgement) == "Triple Stumper" {
		challenge.Stumped = 3
	}

	return nil
}

func parseTiebreakerChallenge(div *html.Node) (*JarchiveChallenge, error) {
	table := nextDescendantWithClass(div, "table", "")
	tiebreaker := new(JarchiveChallenge)
	tiebreaker.Category = innerText(
		nextDescendantWithClass(table, "td", "category_name"))
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

func parseFinalChallenge(div *html.Node) (*JarchiveChallenge, error) {
	table := nextDescendantWithClass(div, "table", "")

	final := new(JarchiveChallenge)
	final.Category = innerText(
		nextDescendantWithClass(table, "td", "category_name"))
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
