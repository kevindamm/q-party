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
// github:kevindamm/q-party/cmd/jarchive/html/seasons.go

package main

import (
	"fmt"
	"io"
	"log"
	"path"
	"regexp"
	"strconv"

	"github.com/kevindamm/q-party/schema"
	"github.com/kevindamm/q-party/selfhost/cmd/fetch"
)

type JarchiveSeason interface {
	fetch.Fetchable
	SeasonSlug() schema.SeasonSlug
	Metadata() *schema.SeasonMetadata
	GetEpisodeMetadata(schema.MatchNumber) *schema.EpisodeMetadata
	GetJarchiveEpisode(EpisodeID) JarchiveEpisode
}

type JarchiveEpisodeIndex map[EpisodeID]JarchiveEpisode

func NewJarchiveSeason(slug schema.SeasonSlug) JarchiveSeason {
	season_index := new(season_index)
	season_index.Slug = slug
	season_index.Episodes = make(JarchiveEpisodeIndex)
	season_index.Matches = make(schema.EpisodeIndex)
	return season_index
}

type season_index struct {
	schema.SeasonMetadata `json:",inline"`
	Episodes              map[EpisodeID]JarchiveEpisode
	Matches               schema.EpisodeIndex `json:"episodes,omitempty"`
}

func (season *season_index) String() string {
	return fmt.Sprintf("season %s", season.Slug)
}

func (season *season_index) URL() string {
	const SEASON_INDEX_FMT = "https://j-archive.com/showseason.php?season=%s"
	return fmt.Sprintf(SEASON_INDEX_FMT, season.Slug)
}

func (season *season_index) FilepathHTML() string {
	return path.Join("season", fmt.Sprintf("%s.html", season.Slug))
}

func (season *season_index) FilepathJSON() string {
	return path.Join("json", "season", fmt.Sprintf("%s.json", season.Slug))
}

func (season *season_index) LoadJSON(input io.ReadCloser) error {
	defer input.Close()
	// TODO
	return nil
}

func (season *season_index) WriteJSON(output io.WriteCloser) error {
	defer output.Close()
	// TODO
	return nil
}

func (season *season_index) ParseHTML(season_html []byte) error {
	errs := make([]error, 0)

	// Since the layout of this file is very simple, we can just use regexes here.
	reSeasonName := regexp.MustCompile(`<h2 class="season">(.*)</h2>`)
	match := reSeasonName.FindSubmatch(season_html)
	if len(match) > 0 {
		season.Title = string(match[1])
	} else {
		log.Printf("no title found in season index")
	}

	reEpisodeLink := regexp.MustCompile(`"showgame\.php\?game_id=(\d+)" title="([^"]+)">.*(aired&#160;([\d-]+))?.*</a>\n`)
	matches := reEpisodeLink.FindAllSubmatch(season_html, -1)
	for _, match := range matches {
		game_id, err := strconv.Atoi(string(match[1]))
		if err != nil {
			errs = append(errs, err)
			continue
		}
		episode := schema.EpisodeMetadata{
			MatchID: schema.NewMatchID(game_id),
		}

		episode.TapedDate = schema.ParseShowDate(string(match[2]))
		if len(match[3]) > 0 {
			episode.AiredDate = schema.ParseShowDate(string(match[4]))
		}
		episode.MatchID.SeasonSlug = season.Slug

		season.AddEpisode(game_id)
		season.Matches.Update(episode)
	}

	return nil
}

func (season *season_index) AddEpisode(game_id int) {
	ep_id := EpisodeID(game_id)
	if _, exists := season.Episodes[ep_id]; exists {
		log.Fatalf("repeated game_id insertion for season %s, episode %d",
			season.Slug, game_id)
	}
	season.Episodes[ep_id] = NewEpisodeStub(ep_id)
	season.EpisodeCount += 1
}

func (season *season_index) SeasonSlug() schema.SeasonSlug {
	return season.Slug
}

func (season *season_index) Metadata() *schema.SeasonMetadata {
	return &season.SeasonMetadata
}

func (season *season_index) GetJarchiveEpisode(epid EpisodeID) JarchiveEpisode {
	return season.Episodes[epid]
}

func (season *season_index) GetEpisodeMetadata(match schema.MatchNumber) *schema.EpisodeMetadata {
	return season.Matches[match]
}
