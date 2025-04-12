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
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kevindamm/q-party/schema"
)

type JarchiveSeason struct {
	Metadata schema.SeasonMetadata `json:",inline"`
	Episodes schema.EpisodeIndex   `json:"episodes"`
}

func FetchSeason(season_id schema.SeasonID, outpath string) (*JarchiveSeason, error) {
	url := SeasonURL(season_id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	filepath := HtmlFilePath(season_id)
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return nil, err
	}

	season_index := new(JarchiveSeason)
	err = ParseSeasonIndexHtml(data, season_index)
	return season_index, err
}

const season_index_url_prefix = "https://j-archive.com/showseason.php?season="

func SeasonURL(season_id schema.SeasonID) string {
	return fmt.Sprint(season_index_url_prefix, season_id.Slug)
}

func HtmlFilePath(season_id schema.SeasonID) string {
	return fmt.Sprintf("%s.html", season_id.Slug)
}

func LoadSeasonIndex(filepath string) (*JarchiveSeason, error) {
	// TODO
	return nil, nil
}

func ParseSeasonIndexHtml(data []byte, season *JarchiveSeason) error {
	errs := make([]string, 0)

	// Since the layout of this file is very simple, we can just use regexes here.
	reSeasonName := regexp.MustCompile(`<h2 class="season">(.*)</h2>`)
	match := reSeasonName.FindSubmatch(data)
	if len(match) > 0 {
		season.Metadata.Season.Title = string(match[1])
	}

	reEpisodeLink := regexp.MustCompile(`"showgame\.php\?game_id=(\d+)"(.*)\n`)
	matches := reEpisodeLink.FindAllSubmatch(data, -1)
	for range matches {
		game_id, err := strconv.Atoi(string(match[1]))
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		episode := schema.EpisodeMetadata{
			MatchID: schema.NewMatchID(game_id),
		}
		episode.MatchID.ShowTitle = string(match[2])
		episode.MatchID.SeasonSlug = season.Metadata.Season.Slug
		season.Episodes.Update(episode)

		season.Metadata.EpisodeCount += 1
	}

	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
