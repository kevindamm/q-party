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
// github:kevindamm/q-party/cmd/jarchive/html/challenge.go

package html

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	qparty "github.com/kevindamm/q-party"
	"golang.org/x/net/html"
)

// The de-normed representation as found in some datasets, e.g. on Kaggle.
type JArchiveChallenge struct {
	qparty.Challenge
	Correct string `json:"correct"`
}

// Assumes that a file extension is present.
var reMediaPath = regexp.MustCompile(`^https?://.*\.com/media/([^.]+)(\.[a-zA-Z0-9]+)`)

func MakeMedia(href string) qparty.Media {
	match := reMediaPath.FindStringSubmatch(href)
	if match == nil {
		return qparty.Media{}
	}
	filename := fmt.Sprintf("/media/%s%s", match[1], match[2])
	mimetype := inferMediaType(match[2])

	return qparty.Media{
		MimeType: mimetype,
		MediaURL: filename}
}

func inferMediaType(ext string) qparty.MimeType {
	switch ext {
	case ".jpg", ".jpeg":
		return qparty.MediaImageJPG
	case ".mp3":
		return qparty.MediaAudioMP3
	case ".mp4":
		return qparty.MediaVideoMP4
	default:
		panic("unrecognized media type for " + ext)
	}
}

func NewChallenge(category string) *JArchiveChallenge {
	challenge := new(JArchiveChallenge)
	challenge.Category = category
	challenge.Media = make([]qparty.Media, 0)
	return challenge
}

func parseChallenge(div *html.Node, challenge *JArchiveChallenge) error {
	if strings.Trim(innerText(div), " ") == "" {
		return nil
	}
	table := nextDescendantWithClass(div, "table", "")

	var err error
	value_td := nextDescendantWithClass(table, "td", "clue_value")
	if value_td != nil {
		challenge.Value, err = qparty.ParseDollarValue(innerText(value_td))
		if err != nil {
			return errors.New("failed to parse challenge value " + err.Error())
		}
	} else {
		dd_value_td := nextDescendantWithClass(table, "td", "clue_value_daily_double")
		if dd_value_td != nil {
			text := strings.ReplaceAll(innerText(dd_value_td), ",", "")
			challenge.Value, err = qparty.ParseDollarValue(text[4:])
			if err != nil {
				return fmt.Errorf("failed to parse daily double value %s\n%s", text, err.Error())
			}
			challenge.Value = -challenge.Value
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
	challenge.Correct = innerText(correct)
	return nil
}

func parseTiebreakerChallenge(div *html.Node, tiebreaker *JArchiveChallenge) error {
	table := nextDescendantWithClass(div, "table", "")
	tiebreaker.Category = innerText(
		nextDescendantWithClass(table, "td", "category_name"))
	tiebreaker.Comments = innerText(
		nextDescendantWithClass(table, "td", "category_comments"))

	clue := nextDescendantWithClass(div, "td", "clue")
	clue_td := nextDescendantWithClass(clue, "td", "clue_text")
	tiebreaker.Clue = innerText(clue_td)
	clue_td = nextSiblingWithClass(clue_td, "td", "clue_text")
	if clue_td == nil {
		return errors.New("could not find tiebreaker challenge response")
	}
	tiebreaker.Correct = innerText(
		nextDescendantWithClass(clue_td, "em", "correct_response"))
	return nil
}

func parseFinalChallenge(div *html.Node, final *JArchiveChallenge) error {
	table := nextDescendantWithClass(div, "table", "")

	final.Category = innerText(
		nextDescendantWithClass(table, "td", "category_name"))
	final.Comments = innerText(
		nextDescendantWithClass(table, "td", "category_comments"))

	clue := nextDescendantWithClass(div, "td", "clue")
	clue_td := nextDescendantWithClass(clue, "td", "clue_text")
	final.Clue = innerText(clue_td)
	clue_td = nextSiblingWithClass(clue_td, "td", "clue_text")
	if clue_td == nil {
		return errors.New("could not find final challenge response")
	}
	final.Correct = innerText(nextDescendantWithClass(clue_td, "em", "correct_response"))
	return nil
}
