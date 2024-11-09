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
	"os"
	"regexp"
	"strconv"

	"github.com/kevindamm/q-party/json"
	"golang.org/x/net/html"
)

type JArchiveEpisode struct {
	JEID `json:"-"`
	json.EpisodeMetadata

	Contestants [3]JArchiveContestant `json:"contestants"`
	Comments    string                `json:"comments,omitempty"`
	Media       []json.Media          `json:"media,omitempty"`

	Single     [6]CategoryChallenges `json:"single,omitempty"`
	Double     [6]CategoryChallenges `json:"double,omitempty"`
	Final      *JArchiveChallenge    `json:"final"`
	TieBreaker *JArchiveChallenge    `json:"tiebreaker,omitempty"`
}

type JEID uint

func (id JEID) String() string {
	return fmt.Sprintf("%d", uint(id))
}

func (id JEID) HTML() string {
	return fmt.Sprintf("%d.html", id)
}

func (id JEID) URL() string {
	return fmt.Sprintf("https://j-archive.com/showgame.php?game_id=%d", id)
}

func MustParseJEID(numeric string) JEID {
	id, err := strconv.Atoi(numeric)
	if err != nil {
		log.Fatalf("failed to parse JEID from string '%s'\n%s", numeric, err)
	}
	return JEID(id)
}

func LoadEpisode(html_path string, metadata JArchiveEpisodeMetadata) (*json.Episode, error) {
	reader, err := os.Open(html_path)
	if err != nil {
		return nil, err
	}
	jaepisode := ParseEpisode(metadata.JEID, reader)
	episode := new(json.Episode)
	episode.ShowNumber = json.ShowNumber(jaepisode.ShowNumber)
	episode.ShowTitle = jaepisode.ShowTitle
	// TODO more properties?

	return episode, nil
}

func parseContent(content *html.Node, episode *JArchiveEpisode) {
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
			episode.Single = parseBoard(child)
		case "double_jeopardy_round":
			episode.Double = parseBoard(child)
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
		episode.ShowNumber = json.ShowNumber(number)
		episode.ShowTitle = match[3]
	} else {
		log.Fatal("title does not match expected format", text)
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

	episode.Final = new(JArchiveChallenge)
	parseFinalChallenge(rounds[0], episode.Final)
	if len(rounds) == 2 {
		episode.TieBreaker = new(JArchiveChallenge)
		parseTiebreakerChallenge(rounds[1], episode.TieBreaker)
	}
}

type JArchiveEpisodeMetadata struct {
	JEID       `json:"-"`
	ShowNumber uint          `json:"show_number,omitempty"`
	Aired      json.ShowDate `json:"aired,omitempty"`
	Taped      json.ShowDate `json:"taped,omitempty"`
}
