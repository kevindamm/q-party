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
	"io"
	"log"
	"regexp"
	"strconv"

	"golang.org/x/net/html"
)

type JArchiveEpisode struct {
	ShowTitle  string  `json:"show_title"`
	ShowNumber int     `json:"show_number" cue:">0"`
	Aired      AirDate `json:"aired"`
	Comments   string  `json:"comments"`

	Single     JArchiveBoard          `json:"single"`
	Double     JArchiveBoard          `json:"double"`
	Final      JArchiveFinalChallenge `json:"final"`
	TieBreaker *JArchiveChallenge     `json:"tiebreaker,omitempty"`
}

func (episode JArchiveEpisode) Filename() string {
	filename := fmt.Sprintf("%d.json", episode.ShowNumber)
	if episode.ShowNumber < 1000 {
		filename = fmt.Sprintf("%03d.json", episode.ShowNumber)
	}
	return filename
}

func episode_url(episode_id int) string {
	return fmt.Sprintf("https://j-archive.com/showgame.php?game_id=%s", episode_id)
}

func ParseEpisode(ep_id string, html_reader io.Reader) *JArchiveEpisode {
	episode := new(JArchiveEpisode)
	doc, err := html.Parse(html_reader)
	if err != nil {
		log.Fatalf("error parsing HTML of %s\n\n%s", ep_id, err)
	}

	child := doc.FirstChild
	for child != nil {
		if isDivWithID(child, "content") {
			episode.parseContent(child)
			break
		}
		child = child.NextSibling
	}

	return episode
}

func (episode *JArchiveEpisode) parseContent(content *html.Node) {
	child := content.FirstChild
	for child != nil {
		if isDivWithID(child, "game_title") {
			episode.parseTitle(child)
		}
		if isDivWithID(child, "game_comments") {
			episode.parseComments(child)
		}
		if isDivWithID(child, "contestants") {
			episode.parseContestants(child)
		}
		if isDivWithID(child, "jeopardy_round") {
			episode.parseBoard(child, ROUND_SINGLE_JEOPARDY)
		}
		if isDivWithID(child, "double_jeopardy_round") {
			episode.parseBoard(child, ROUND_DOUBLE_JEOPARDY)
		}
		if isDivWithID(child, "final_jeopardy_round") {
			episode.parseFinalChallenge(child)
			// TODO on a rare occasion there is also a tiebreaker question
		}

		child = child.NextSibling
	}

}

func isDivWithID(node *html.Node, id string) bool {
	if node.Type != html.ElementNode || node.Data != "div" {
		return false
	}
	for _, attr := range node.Attr {
		if attr.Key == "id" && attr.Val == id {
			return true
		}
	}
	return false
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
	// Expect the only child element to be a text node.
	if game_comments.FirstChild.Type != html.TextNode {
		return
	}
	episode.Comments = game_comments.FirstChild.Data
}

func (episode *JArchiveEpisode) parseContestants(game_comments *html.Node) {
	// TODO not necessary but could be nice for tracking a contestant's career
}

func (episode *JArchiveEpisode) parseBoard(board *html.Node, round EpisodeRound) {

}

func (episode *JArchiveEpisode) parseFinalChallenge(div *html.Node) {

}

func (episode *JArchiveEpisode) parseTiebreaker(div *html.Node) {

}
