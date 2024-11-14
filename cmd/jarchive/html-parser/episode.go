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

type EpisodeMetadata struct {
	JEID                   `json:"-"`
	qparty.EpisodeMetadata `json:",inline"`
}

type JArchiveEpisode struct {
	EpisodeMetadata
	Contestants [3]qparty.Contestant `json:"contestants"`
	Comments    string               `json:"comments,omitempty"`
	Media       []qparty.Media       `json:"media,omitempty"`

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

func LoadEpisode(html_path string, metadata EpisodeMetadata) (*qparty.Episode, error) {
	reader, err := os.Open(html_path)
	if err != nil {
		return nil, err
	}
	jaepisode := ParseEpisode(metadata.JEID, reader)
	episode := new(qparty.Episode)
	episode.ShowNumber = qparty.ShowNumber(jaepisode.ShowNumber)
	episode.ShowTitle = jaepisode.ShowTitle
	for i := range 3 {
		episode.ContestantIDs[i] = jaepisode.Contestants[i].ContestantID
	}

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
		episode.ShowNumber = qparty.ShowNumber(number)
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

func FetchEpisode(episode JEID, filepath string) error {
	url := episode.URL()
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

func parseBoard(root *html.Node) [6]CategoryChallenges {
	categories := [6]CategoryChallenges{}
	round_table := nextDescendantWithClass(root, "table", "round")
	category_tr := nextDescendantWithClass(round_table, "tr", "")
	category_tds := childrenWithClass(category_tr, "td", "category")
	if len(category_tds) != 6 {
		log.Fatal("expected 6 category entries, found", len(category_tds))
	}
	for i, td := range category_tds {
		err := parseCategoryHeader(td, &categories[i])
		if err != nil {
			log.Fatal("failed to parse category header (name and comments)")
		}
	}

	for range 5 {
		category_tr = nextSiblingWithClass(category_tr, "tr", "")
		clue_tds := childrenWithClass(category_tr, "td", "clue")
		if len(clue_tds) != 6 {
			log.Fatal("expected 6 clue entries, found", len(clue_tds))
		}
		for i, clue_td := range clue_tds {
			err := parseCategoryChallenge(clue_td, &categories[i])
			if err != nil {
				log.Fatal("failed to parse clue entry", err)
			}
		}
	}
	return categories
}
