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
// github:kevindamm/q-party/selfhost/cmd/fetch/jarchive/episodes.go

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"path"
	"regexp"
	"strconv"

	"github.com/kevindamm/q-party/schema"
	"github.com/kevindamm/q-party/selfhost/cmd/fetch"
	"golang.org/x/net/html"
)

type EpisodeID int

// All details of the episode, including correct answers & the contestants' bios.
// Properties are set via Fetchable interface: ParseHTML, LoadJSON.
type JarchiveEpisode interface {
	fetch.Fetchable
	Metadata() *schema.MatchMetadata
}

// Constructor for an episode when only its (jarchive) EpisodeID is known.
func NewEpisodeStub(id EpisodeID) JarchiveEpisode {
	if id == 0 {
		log.Fatal("zero value invalid for episode ID (equiv to UNKNOWN episode)")
	}
	return &episode{
		EpisodeID: id,
	}
}

func UnknownEpisode() JarchiveEpisode {
	return &episode{EpisodeID: 0}
}

type episode struct {
	schema.MatchMetadata `json:",inline"`

	EpisodeID EpisodeID         `json:"episode_id,omitempty"`
	Comments  string            `json:"comments,omitempty"`
	Media     []schema.MediaRef `json:"media,omitempty"`

	// Due to absence of archival evidence,
	// not every episode has both single & double rounds.
	Single     *JarchiveBoard          `json:"single,omitempty"`
	Double     *JarchiveBoard          `json:"double,omitempty"`
	Final      *JarchiveFinalChallenge `json:"final,omitempty"`
	TieBreaker *JarchiveTiebreaker     `json:"tiebreaker,omitempty"`
}

func (episode *episode) String() string {
	return fmt.Sprintf("Match %d (episode #%d)",
		episode.MatchNumber,
		episode.EpisodeID)
}

func (episode *episode) Metadata() *schema.MatchMetadata {
	return &episode.MatchMetadata
}

func (episode *episode) URL() string {
	const FULL_EPISODE_FMT = "https://j-archive.com/showgame.php?game_id=%d"
	return fmt.Sprintf(FULL_EPISODE_FMT, episode.EpisodeID)
}

func (episode *episode) FilepathHTML() string {
	return path.Join("jarchive", "episode",
		fmt.Sprintf("%d.html", episode.EpisodeID))
}

func (episode *episode) FilepathJSON() string {
	return path.Join("json", "episode",
		fmt.Sprintf("%d.html", episode.MatchNumber))
}

func (episode *episode) ParseHTML(html_bytes []byte) error {
	doc, err := html.Parse(bytes.NewReader(html_bytes))
	if err != nil {
		log.Fatal("error parsing HTML\n", err)
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
			parseContent(child, episode)
			break
		}
		child = child.NextSibling
	}

	return nil
}

func (episode *episode) WriteJSON(output io.WriteCloser) error {
	defer output.Close()
	// TODO json.Unmarshal
	return nil
}

func (episode *episode) LoadJSON(input io.ReadCloser) error {
	defer input.Close()
	// TODO json.Marshal
	return nil
}

func MustParseEpisodeID(numeric string) EpisodeID {
	id, err := strconv.Atoi(numeric)
	if err != nil {
		log.Fatalf("failed to parse JEID from string '%s'\n%s", numeric, err)
	}
	return EpisodeID(id)
}

func parseContent(content *html.Node, episode *episode) {
	child := content.FirstChild
	for child != nil {
		id := divWithID(child)
		if id == "" {
			child = child.NextSibling
			continue
		}

		switch id {
		case "game_title":
			nextChild := child.FirstChild
			for nextChild.Type != html.ElementNode {
				nextChild = child.NextSibling
			}
			if nextChild.Data != "h1" {
				log.Print("odd, div#game_title does not have an H1 as first sub-element")
				break
			}

			text, media := innerTextMarkdown(nextChild)
			episode.MatchID.MatchNumber = schema.MatchNumber(parseShowNumber(text))
			// derived from content, not <head>...</head>
			episode.ShowTitle = text
			if media != nil {
				episode.Media = media
			}

		case "game_comments":
			text, media := innerTextMarkdown(child)
			episode.Comments = text
			if len(media) > 0 {
				episode.Media = media
			}

		case "jeopardy_round":
			episode.Single = ParseBoard(child)

		case "double_jeopardy_round":
			episode.Double = ParseBoard(child)

		case "final_jeopardy_round":
			episode.Final, episode.TieBreaker = parseFinalRound(child)

		default: // pass on unrecognized class names.
		}

		child = child.NextSibling
	}
}

var reShowNumberMatcher = regexp.MustCompile(`#(\d+)`)

func parseShowNumber(full_title string) uint {
	showNumMatch := reShowNumberMatcher.FindAllStringSubmatch(full_title, 2)
	if showNumMatch == nil {
		log.Fatal("title does not match expected format", full_title)
	}
	if len(showNumMatch) > 1 {
		log.Fatal("more than one pattern match for #\\d+ in title")
	}
	// By regex we know this to be a positive integer.
	number, _ := strconv.Atoi(showNumMatch[0][1])
	if number < 1 {
		log.Fatal("show number should be a positive number")
	}
	return uint(number)
}

func parseFinalRound(div *html.Node) (*JarchiveFinalChallenge, *JarchiveTiebreaker) {
	// On a rare occasion there is also a tiebreaker question,
	// with two instead of one <div class="final_round">
	rounds := childrenWithClass(div, "table", "final_round")
	if len(rounds) == 0 {
		panic("did not find any final_round in this episode")
	}

	final, err := ParseFinalChallenge(rounds[0])
	if err != nil {
		log.Println("ERROR\n", err)
		return nil, nil
	}
	var tiebreak *JarchiveTiebreaker
	if len(rounds) == 2 {
		tiebreak, err = ParseTiebreakerChallenge(rounds[1])
		if err != nil {
			log.Println("ERROR\n", err)
			return final, nil
		}
	}
	return final, tiebreak
}
