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
// github:kevindamm/q-party/cmd/jarchive/episode.go

package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"golang.org/x/net/html"
)

type JArchiveEpisode struct {
	JEID
	ShowTitle  string   `json:"show_title"`
	ShowNumber int      `json:"show_number" cue:">0"`
	Aired      ShowDate `json:"aired,omitempty"`
	Taped      ShowDate `json:"taped,omitempty"`
	Comments   string   `json:"comments,omitempty"`
	Media      []Media  `json:"media,omitempty"`

	Contestants [3]JArchiveContestant `json:"contestants"`

	Single     *JArchiveBoard         `json:"single,omitempty"`
	Double     *JArchiveBoard         `json:"double,omitempty"`
	Final      JArchiveFinalChallenge `json:"final"`
	TieBreaker *JArchiveTiebreaker    `json:"tiebreaker,omitempty"`
}

func NewEpisode(metadata JArchiveEpisodeMetadata) *JArchiveEpisode {
	episode := new(JArchiveEpisode)
	episode.JEID = metadata.JEID
	if metadata.Aired != unknown_airing {
		episode.Aired = metadata.Aired
	}
	if metadata.Taped != unknown_taping {
		episode.Taped = metadata.Taped
	}
	return episode
}

func (episode *JArchiveEpisode) parseContent(content *html.Node) {
	child := content.FirstChild
	for child != nil {
		id := divWithID(child)
		if id == "" {
			child = child.NextSibling
			continue
		}

		switch id {
		case "game_title":
			episode.parseTitle(child) // derived from content, not <head>
		case "game_comments":
			episode.parseComments(child) // minutiae about the episode
		case "contestants":
			episode.parseContestants(child) // name & bio
		case "jeopardy_round":
			episode.Single = NewBoard(episode.Aired, ROUND_SINGLE_JEOPARDY)
			episode.Single.parseBoard(child)
		case "double_jeopardy_round":
			episode.Double = NewBoard(episode.Aired, ROUND_DOUBLE_JEOPARDY)
			episode.Double.parseBoard(child)
		case "final_jeopardy_round":
			episode.parseFinalRound(child) // may include tie-breaker
		}

		child = child.NextSibling
	}

}

var reTitleMatcher = regexp.MustCompile(`.*([Ss]how|pilot|game) #(\d+),? - (.*)`)

func (episode *JArchiveEpisode) parseTitle(game_title *html.Node) {
	// Expect first child to be an H1 tag
	child := game_title.FirstChild
	if child.Type != html.ElementNode || child.Data != "h1" || child.FirstChild.Type != html.TextNode {
		return
	}

	text := child.FirstChild.Data
	match := reTitleMatcher.FindStringSubmatch(text)
	if match != nil {
		// Pattern matching determines match[2] will always be numeric.
		number, _ := strconv.Atoi(match[2])
		episode.ShowNumber = number
		episode.ShowTitle = match[3] // TODO the AirDate can usually be determined from this.
	}
}

func (episode *JArchiveEpisode) parseComments(game_comments *html.Node) {
	text, media := parseIntoMarkdown(game_comments)
	episode.Comments = text
	if len(media) > 0 {
		episode.Media = media
	}
}

func (episode *JArchiveEpisode) parseFinalRound(div *html.Node) {
	// On a rare occasion there is also a tiebreaker question,
	// with two instead of one <div class="final_round">
	rounds := childrenWithClass(div, "table", "final_round")
	if len(rounds) == 0 {
		panic("did not find any final_round in this episode")
	}

	episode.Final.parseChallenge(rounds[0])
	if len(rounds) == 2 {
		episode.TieBreaker = new(JArchiveTiebreaker)
		episode.TieBreaker.parseChallenge(rounds[1])
		episode.TieBreaker.ShowDate = episode.Aired
	}
}

type JArchiveEpisodeMetadata struct {
	JEID  `json:"-"`
	Taped ShowDate `json:"taped,omitempty"`
	Aired ShowDate `json:"aired,omitempty"`
}

// Unique numeric identifier for episodes in the archive.
// May be different than the sequential show number used in display.
type JEID int

func (id JEID) String() string {
	return fmt.Sprintf("%d", int(id))
}

func (id JEID) HTML() string {
	return fmt.Sprintf("%d.html", id)
}

func (id JEID) JSON() string {
	return fmt.Sprintf("%d.json", id)
}

func (id JEID) URL() string {
	return fmt.Sprintf("https://j-archive.com/showgame.php?game_id=%d", id)
}

// Parses the numeric value from a string.
// Fatal error if the value cannot be converted into a number.
func MustParseJEID(numeric string) JEID {
	id, err := strconv.Atoi(numeric)
	if err != nil {
		log.Fatalf("failed to parse JEID from string '%s'\n%s", numeric, err)
	}
	return JEID(id)
}
