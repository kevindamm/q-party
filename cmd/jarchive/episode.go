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
	"strings"

	"golang.org/x/net/html"
)

type JArchiveEpisode struct {
	ShowTitle  string  `json:"show_title"`
	ShowNumber int     `json:"show_number" cue:">0"`
	Aired      AirDate `json:"aired"`
	Comments   string  `json:"comments"`

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

func ParseEpisode(ep_id string, html_reader io.Reader) *JArchiveEpisode {
	episode := new(JArchiveEpisode)
	doc, err := html.Parse(html_reader)
	if err != nil {
		log.Fatalf("error parsing HTML of %s\n\n%s", ep_id, err)
	}

	child := doc.FirstChild
	for child != nil {
		if child.Type == html.DocumentNode ||
			(child.Type == html.ElementNode && child.Data == "html") ||
			(child.Type == html.ElementNode && child.Data == "body") {
			child = child.FirstChild
			continue
		}
		if divWithID(child) == "content" {
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
	// Expect the only child element to be a text node.
	if game_comments.FirstChild.Type != html.TextNode {
		return
	}
	episode.Comments = game_comments.FirstChild.Data
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

// Returns a list of the direct children elements that have the indicated class.
// If elType is not the empty string, only returns elements of that type.  Only
// elements are returned; text nodes and other non-element children are ignored.
func childrenWithClass(node *html.Node, elType string, elClass string) []*html.Node {
	matching_children := make([]*html.Node, 0)
	child := node.FirstChild

	for child != nil {
		if child.Type == html.ElementNode &&
			(elType == "" || child.Data == elType) {
			for _, attr := range child.Attr {
				if attr.Key == "class" {
					// Handles the situation where an element has multiple classes.
					for _, aclass := range strings.Split(attr.Val, " ") {
						if aclass == elClass {
							matching_children = append(matching_children, child)
							break
						}
					}
					break
				}
			}
		}

		child = child.NextSibling
	}
	return matching_children
}

// Searches recursively through descendents (DFS) looking for the next element
// with the indicated type and class.  If class=="" then any (or no) class will
// satisfy the search.  It returns the first matching subelement, depth first.
// If there is no matching element (and class) then nil is returned instead.
func nextDescendantWithClass(node *html.Node, elType string, elClass string) *html.Node {
	var found *html.Node = nil
	var recursiveFind func(*html.Node, string, string)
	recursiveFind = func(next *html.Node, elType string, elClass string) {
		child := next.FirstChild
		for child != nil {
			if child.Type == html.ElementNode &&
				(elType == "" || child.Data == elType) {
				if elClass == "" {
					found = child
					return
				}
				for _, attr := range child.Attr {
					if attr.Key == "class" {
						// Handles the situation where an element has multiple classes.
						for _, aclass := range strings.Split(attr.Val, " ") {
							if aclass == elClass {
								found = child
								return
							}
						}
						break
					}
				}
			}
			if child.FirstChild != nil {
				recursiveFind(child, elType, elClass)
				if found != nil {
					return
				}
			}
			child = child.NextSibling
		}
	}
	recursiveFind(node, elType, elClass)
	return found
}

// Returns the ID of the node if it is a <div> element,
// otherwise returns the empty string.
func divWithID(node *html.Node) string {
	if node.Type != html.ElementNode || node.Data != "div" {
		// Not a <div>.
		return ""
	}
	for _, attr := range node.Attr {
		if attr.Key == "id" {
			return attr.Val
		}
	}

	// Is a <div> but has no ID.
	return ""
}

func innerText(node *html.Node) string {
	text := make([]string, 0)
	fmt.Printf("%v\n", node)

	var recursiveFind func(*html.Node)
	recursiveFind = func(node *html.Node) {
		if node == nil {
			return
		}
		fmt.Println("node data " + node.Data)

		child := node.FirstChild
		for child != nil {
			if child.Type == html.TextNode {
				text = append(text, child.Data)
			}
			if child.FirstChild != nil {
				recursiveFind(child)
			}
			child = child.NextSibling
		}
	}
	recursiveFind(node)

	return strings.ReplaceAll(strings.Join(text, " "), "  ", " ")
}
