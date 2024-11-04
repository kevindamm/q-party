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
	"time"

	"golang.org/x/net/html"
)

type JArchiveEpisode struct {
	JEID
	ShowTitle  string  `json:"show_title"`
	ShowNumber int     `json:"show_number" cue:">0"`
	Aired      AirDate `json:"aired"`
	Comments   string  `json:"comments"`

	Contestants [3]JArchiveContestant `json:"contestants"`

	Single     JArchiveBoard           `json:"single"`
	Double     JArchiveBoard           `json:"double"`
	Final      JArchiveFinalChallenge  `json:"final"`
	TieBreaker *JArchiveFinalChallenge `json:"tiebreaker,omitempty"`
}

func (episode JArchiveEpisode) Filename() string {
	filename := fmt.Sprintf("%d.json", episode.ShowNumber)
	if episode.ShowNumber < 1000 {
		filename = fmt.Sprintf("%03d.json", episode.ShowNumber)
	}
	return filename
}

func episode_url(episode_id int) string {
	return fmt.Sprintf("https://j-archive.com/showgame.php?game_id=%d", episode_id)
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
			episode.parseTitle(child)
		case "game_comments":
			episode.parseComments(child)
		case "contestants":
			episode.parseContestants(child)
		case "jeopardy_round":
			episode.parseBoard(child, ROUND_SINGLE_JEOPARDY)
		case "double_jeopardy_round":
			episode.parseBoard(child, ROUND_DOUBLE_JEOPARDY)
		case "final_jeopardy_round":
			episode.parseFinalRound(child)
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
	episode.Comments = innerText(game_comments)
}

func (episode *JArchiveEpisode) parseBoard(div *html.Node, round EpisodeRound) {
	if round == ROUND_SINGLE_JEOPARDY {
		episode.Single.parseBoard(div)
	} else {
		episode.Double.parseBoard(div)
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
		episode.TieBreaker = new(JArchiveFinalChallenge)
		episode.TieBreaker.parseChallenge(rounds[1])
		episode.TieBreaker.Round = ROUND_TIE_BREAKER
	}
}

type JArchiveEpisodeMetadata struct {
	JEID  `json:"-"`
	Taped AirDate `json:"taped"`
	Aired AirDate `json:"aired"`
}

var TimeUnknown = time.Unix(0, 0)

// Unique numeric identifier for episodes in the archive.
// May be different than the sequential show number used in display.
type JEID int

func (id JEID) String() string {
	return fmt.Sprintf("%d", int(id))
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
