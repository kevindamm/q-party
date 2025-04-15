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
// github:kevindamm/q-party/cmd/fetch/jarchive/fetch.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/kevindamm/q-party/schema"
)

type Fetcher interface {
	FetchIndex() JarchiveIndex
	FetchSeason(schema.SeasonSlug) <-chan JarchiveSeason
	FetchEpisode(schema.MatchNumber) <-chan JarchiveEpisode

	Errors() <-chan error
	Close()
}

func NewFetcher(pause time.Duration) Fetcher {
	fetcher := new(fetcher)
	fetcher.urls = make(chan fetch_task)
	fetcher.ticker = time.NewTicker(pause)
	fetcher.errors = make(chan error)

	return fetcher
}

type fetcher struct {
	urls   chan fetch_task
	ticker *time.Ticker
	errors chan error
}

type fetch_task struct {
	url      string
	response chan []byte
}

func (fetcher *fetcher) FetchIndex() JarchiveIndex {
	index := new(index)
	//url := IndexURL()
	//response, err := http.Get(url)
	//if err != nil {
	//	fetcher.errors <- err
	//	return nil
	//}

	// TODO
	return index
}

func (fetcher *fetcher) FetchSeason(season_slug schema.SeasonSlug) <-chan JarchiveSeason {
	url := SeasonURL(season_slug)
	response, err := http.Get(url)
	if err != nil {
		fetcher.errors <- err
		return nil
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fetcher.errors <- err
		return nil
	}

	filepath := SeasonIndexHtmlPath(season_slug)
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		fetcher.errors <- err
		return nil
	}

	season_index, err := ParseSeasonIndexHtml(data)
	if err != nil {
		fetcher.errors <- err
		return nil
	}
	season_chan := make(chan JarchiveSeason)
	go func() { season_chan <- season_index; close(season_chan) }()
	return season_chan
}

func (fetcher *fetcher) FetchEpisode(episode schema.MatchNumber) <-chan JarchiveEpisode {

	return make(chan JarchiveEpisode)
}

func (fetcher *fetcher) Errors() <-chan error {
	return fetcher.errors
}

func (fetcher *fetcher) Close() {
	fetcher.ticker.Stop()
	close(fetcher.errors)
	close(fetcher.urls)
}

func IndexURL() string {
	return "https://j-archive.com/listseasons.php"
}

func SeasonURL(season_slug schema.SeasonSlug) string {
	const SEASON_INDEX_FMT = "https://j-archive.com/showseason.php?season=%s"
	return fmt.Sprintf(SEASON_INDEX_FMT, season_slug)
}

func EpisodeURL(id EpisodeID) string {
	const FULL_EPISODE_FMT = "https://j-archive.com/showgame.php?game_id=%d"
	return fmt.Sprintf(FULL_EPISODE_FMT, id)
}

func SeasonIndexHtmlPath(season_slug schema.SeasonSlug) string {
	return fmt.Sprintf("%s.html", season_slug)
}

func EpisodeHtmlPath(season schema.SeasonSlug, match schema.MatchNumber) string {
	return fmt.Sprintf("%s-%d.html", season, match)
}
