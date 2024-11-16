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
// github:kevindamm/q-party/cmd/jarchive/html/episode.go

package html

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	qparty "github.com/kevindamm/q-party"
	"golang.org/x/net/html"
)

func MustParseJEID(numeric string) qparty.EpisodeID {
	id, err := strconv.Atoi(numeric)
	if err != nil {
		log.Fatalf("failed to parse JEID from string '%s'\n%s", numeric, err)
	}
	return qparty.EpisodeID(id)
}

func parseContent(content *html.Node, episode *qparty.FullEpisode) {
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

			text, media := parseIntoMarkdown(nextChild)
			episode.ShowNumber = parseShowNumber(text)
			// derived from content, not <head>...</head>
			episode.ShowTitle = text
			if media != nil {
				episode.Media = media
			}

		case "game_comments":
			text, media := parseIntoMarkdown(child)
			episode.Comments = text
			if len(media) > 0 {
				episode.Media = media
			}

		case "contestants":
			episode.Contestants = parseContestants(child) // name & bio

		case "jeopardy_round":
			episode.Single = parseBoard(child)

		case "double_jeopardy_round":
			episode.Double = parseBoard(child)

		case "final_jeopardy_round":
			episode.Final, episode.TieBreaker = parseFinalRound(child)

		default: // pass on unrecognized class names.
		}

		child = child.NextSibling
	}

}

var reShowNumberMatcher = regexp.MustCompile(`#(\d+)`)

func parseShowNumber(full_title string) qparty.ShowNumber {
	showNumMatch := reShowNumberMatcher.FindAllStringSubmatch(full_title, 2)
	if showNumMatch == nil {
		log.Fatal("title does not match expected format", full_title)
	}
	if len(showNumMatch) > 1 {
		log.Fatal("more than one pattern match for #\\d+ in title")
	}
	number := showNumMatch[0][1]
	return qparty.ShowNumber(number)
}

func parseFinalRound(div *html.Node) (*qparty.FullChallenge, *qparty.FullChallenge) {
	// On a rare occasion there is also a tiebreaker question,
	// with two instead of one <div class="final_round">
	rounds := childrenWithClass(div, "table", "final_round")
	if len(rounds) == 0 {
		panic("did not find any final_round in this episode")
	}

	final, err := parseFinalChallenge(rounds[0])
	if err != nil {
		log.Println("ERROR\n", err)
		return nil, nil
	}
	var tiebreak *qparty.FullChallenge
	if len(rounds) == 2 {
		tiebreak, err = parseTiebreakerChallenge(rounds[1])
		if err != nil {
			log.Println("ERROR\n", err)
			return final, nil
		}
	}
	return final, tiebreak
}

func FetchEpisode(episode qparty.EpisodeID, filepath string) error {
	url := EpisodeURL(episode)
	log.Print("Fetching ", url, "  -> ", filepath)

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	return nil
}

func EpisodeURL(id qparty.EpisodeID) string {
	return fmt.Sprintf("https://j-archive.com/showgame.php?game_id=%d", id)
}
