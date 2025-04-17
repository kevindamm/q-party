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
	"errors"
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
	GetEpisodeByID(EpisodeID) JarchiveEpisode
	GetEpisodeByNumber(schema.MatchNumber) *schema.EpisodeMetadata
}

func LoadSeasonIndex(filepath string) (JarchiveSeason, error) {
	// TODO
	return nil, nil
}

func ParseSeasonIndexHtml(data []byte) (JarchiveSeason, error) {
	season := new(season_index)
	errs := make([]error, 0)

	// Since the layout of this file is very simple, we can just use regexes here.
	reSeasonName := regexp.MustCompile(`<h2 class="season">(.*)</h2>`)
	match := reSeasonName.FindSubmatch(data)
	if len(match) > 0 {
		season.Season.Title = string(match[1])
	} else {
		log.Printf("no title found in season index")
	}

	reEpisodeLink := regexp.MustCompile(`"showgame\.php\?game_id=(\d+)"(.*)\n`)
	matches := reEpisodeLink.FindAllSubmatch(data, -1)
	for _, match := range matches {
		game_id, err := strconv.Atoi(string(match[1]))
		if err != nil {
			errs = append(errs, err)
			continue
		}
		episode := schema.EpisodeMetadata{
			MatchID: schema.NewMatchID(game_id),
		}
		episode.MatchID.ShowTitle = string(match[2])
		episode.MatchID.SeasonSlug = season.Season.Slug

		ep_id := EpisodeID(game_id)
		season.Episodes[ep_id] = NewEpisode(ep_id)
		season.EpisodeCount += 1
		season.Matches.Update(episode)
	}

	return season, errors.Join(errs...)
}

type season_index struct {
	schema.SeasonMetadata `json:",inline"`
	Episodes              map[EpisodeID]JarchiveEpisode
	Matches               schema.EpisodeIndex `json:"episodes,omitempty"`
}

func (season *season_index) String() string {
	return fmt.Sprintf("season %s", season.Season.Slug)
}

func (season *season_index) URL() string {
	const SEASON_INDEX_FMT = "https://j-archive.com/showseason.php?season=%s"
	return fmt.Sprintf(SEASON_INDEX_FMT, season.Season.Slug)
}

func (season *season_index) FilepathHTML() string {
	return path.Join("season", fmt.Sprintf("%s.html", season.Season.Slug))
}

func (season *season_index) FilepathJSON() string {
	return path.Join("json", "season", fmt.Sprintf("%s.json", season.Season.Slug))
}

func (season *season_index) ParseHTML(html_bytes []byte) error {
	// TODO
	return nil
}

func (season *season_index) WriteJSON(output io.WriteCloser) error {
	defer output.Close()
	// TODO
	return nil
}

func (season *season_index) LoadJSON(input io.ReadCloser) error {
	defer input.Close()
	// TODO
	return nil
}

func (season *season_index) GetEpisodeByID(epid EpisodeID) JarchiveEpisode {
	return season.Episodes[epid]
}

func (season *season_index) GetEpisodeByNumber(match schema.MatchNumber) *schema.EpisodeMetadata {
	return season.Matches[match]
}
