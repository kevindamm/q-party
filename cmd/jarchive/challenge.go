package main

import (
	"errors"
	"fmt"
	"regexp"
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

	Prompt  string  `json:"prompt"`
	Media   []Media `json:"media,omitempty"`
	Correct string  `json:"correct"` // excluding "what is..." preface
}

type JArchiveFinalChallenge struct {
	JArchiveChallenge
}

type JArchiveTiebreaker struct {
	JArchiveChallenge
}

type Media struct {
	MediaType `json:"type,omitempty"`
	MediaURL  string `json:"url"`
}

// This enumeration over available media types is modeled after its equivalent
// MIME type such as image/jpeg, image/png, audio/mpeg, etc.  The default (its
// zero value) is an empty string which implicitly represents text/plain, UTF-8.
type MediaType string

const (
	MediaTextUTF8 MediaType = "" // default is text/plain;charset=UTF-8
	MediaImageJPG MediaType = "image/jpeg"
	MediaAudioMP3 MediaType = "audio/mpeg"
	MediaVideoMP4 MediaType = "video/mp4"
)

// Assumes that a file extension is present.
var reMediaPath = regexp.MustCompile(`^https?://.*\.com/media/([^.]+)(\.[a-zA-Z0-9]+)`)
var unknown_media = Media{"", ""}

func MakeMedia(href string) Media {
	match := reMediaPath.FindStringSubmatch(href)
	if match == nil {
		return unknown_media
	}
	filename := fmt.Sprintf("/media/%s%s", match[1], match[2])
	mimetype := inferMediaType(match[2])

	return Media{
		MediaType: mimetype,
		MediaURL:  filename}
}

func inferMediaType(ext string) MediaType {
	switch ext {
	case ".jpg", ".jpeg":
		return MediaImageJPG
	case ".mp3":
		return MediaAudioMP3
	case ".mp4":
		return MediaVideoMP4
	default:
		panic("unrecognized media type for " + ext)
	}
}

func NewChallenge() *JArchiveChallenge {
	challenge := new(JArchiveChallenge)
	challenge.Media = make([]Media, 0)
	return challenge
}

func (challenge JArchiveChallenge) IsEmpty() bool {
	if challenge.Value != 0 ||
		challenge.Prompt != "" ||
		challenge.Correct != "" {
		return false
	}
	return true
}

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

	clue_td := nextDescendantWithClass(table, "td", "clue_text")
	text, media := parseIntoMarkdown(clue_td)
	challenge.Prompt = text
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

func (tiebreaker *JArchiveTiebreaker) parseChallenge(div *html.Node) error {
	table := nextDescendantWithClass(div, "table", "")
	tiebreaker.Round = ROUND_TIE_BREAKER
	tiebreaker.Category = JArchiveCategory(innerText(
		nextDescendantWithClass(table, "td", "category_name")))
	tiebreaker.Commentary = innerText(
		nextDescendantWithClass(table, "td", "category_comments"))

	clue := nextDescendantWithClass(div, "td", "clue")
	clue_td := nextDescendantWithClass(clue, "td", "clue_text")
	tiebreaker.Prompt = innerText(clue_td)
	clue_td = nextSiblingWithClass(clue_td, "td", "clue_text")
	if clue_td == nil {
		return errors.New("could not find tiebreaker challenge response")
	}
	tiebreaker.Correct = innerText(
		nextDescendantWithClass(clue_td, "em", "correct_response"))
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
