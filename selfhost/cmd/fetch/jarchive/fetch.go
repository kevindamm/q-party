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
	"path"
	"strings"
	"time"

	"github.com/kevindamm/q-party/schema"
)

// Abstraction over time-delayed fetching of URLs.
type Fetcher interface {
	FetchIndex() <-chan JarchiveIndex
	FetchSeason(schema.SeasonSlug) <-chan JarchiveSeason
	FetchEpisode(EpisodeID, schema.MatchNumber) <-chan *JarchiveEpisode

	Errors() <-chan error
	Close()
}

// Creates a Fetcher instance and initiates its goroutine for delayed fetching.
func NewFetcher(pause time.Duration, useragent string) Fetcher {
	useragent = strings.Trim(useragent, " \t\n")
	useragent = strings.ReplaceAll(useragent, "'\"", "")
	if len(useragent) == 0 {
		useragent = DefaultUserAgent()
	}

	fetcher := fetcher{
		useragent: useragent,
		urls:      make(chan fetch_task),
		errors:    make(chan error),
	}
	go fetcher.start_titration(pause)

	return &fetcher
}

// Internal representation of the state, metadata and channels of the Fetcher.
type fetcher struct {
	ticker    *time.Ticker
	useragent string

	urls   chan fetch_task
	errors chan error
}

// A goroutine for periodically waiting between http.Get() requests.
func (fetcher *fetcher) start_titration(pause time.Duration) {
	// pause will be at least 2.5 seconds
	min_pause := 2500 * time.Millisecond
	fetcher.ticker = time.NewTicker(max(min_pause, pause))

	for range fetcher.ticker.C {
		select {
		case task, _ := <-fetcher.urls:
			request, err := http.NewRequest("GET", task.url, nil)
			if err != nil {
				fetcher.errors <- err
				continue
			}
			request.Header.Set("User-Agent", fetcher.useragent)
			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				fetcher.errors <- err
				continue
			}
			defer response.Body.Close()

			bytes, err := io.ReadAll(response.Body)
			if err != nil {
				fetcher.errors <- err
			} else {
				task.output <- bytes
			}
			close(task.output)

		default: // channel is empty, keep waiting
		}
	}
}

// Represents an awaiting task being communicated on the fetcher's internal
// channel.  Each request is assumed to be a GET to the indicated URL.
type fetch_task struct {
	url    string
	output chan<- []byte
}

// Fetch the top-level index of jarchive seasons.
func (fetcher *fetcher) FetchIndex() <-chan JarchiveIndex {
	bytes_chan := make(chan []byte)
	task := fetch_task{
		url:    IndexURL(),
		output: bytes_chan,
	}
	fetcher.urls <- task

	// async write-and-parse when fetcher gets to this task
	index_chan := make(chan JarchiveIndex)
	go func() {
		defer close(index_chan)
		index_html := <-bytes_chan

		filepath := IndexHtmlPath()
		err := os.WriteFile(filepath, index_html, 0644)
		if err != nil {
			fetcher.errors <- err
			return
		}

		index, err := ParseIndexHtml(index_html)
		if err != nil {
			fetcher.errors <- err
			return
		}
		index_chan <- index
	}()

	// caller awaits on the JarchiveIndex that is produced
	return index_chan
}

// Fetch the season metadata and its list of episodes.
func (fetcher *fetcher) FetchSeason(season_slug schema.SeasonSlug) <-chan JarchiveSeason {
	bytes_chan := make(chan []byte)
	task := fetch_task{
		url:    SeasonURL(season_slug),
		output: bytes_chan,
	}
	fetcher.urls <- task

	// async write-and-parse when fetcher gets to this task
	season_chan := make(chan JarchiveSeason)
	go func() {
		defer close(season_chan)
		season_html := <-bytes_chan

		filepath := SeasonIndexHtmlPath(season_slug)
		err := os.WriteFile(filepath, season_html, 0644)
		if err != nil {
			fetcher.errors <- err
			return
		}

		season_index, err := ParseSeasonIndexHtml(season_html)
		if err != nil {
			fetcher.errors <- err
			return
		}
		season_chan <- season_index
	}()

	// caller awaits on the JarchiveSeason that is produced
	return season_chan
}

// Fetch an episode's metadata (categories, challenges & contestants), keyed by the .
func (fetcher *fetcher) FetchEpisode(ja_eid EpisodeID, match schema.MatchNumber) <-chan *JarchiveEpisode {
	bytes_chan := make(chan []byte)
	task := fetch_task{
		url:    EpisodeURL(ja_eid),
		output: bytes_chan,
	}
	fetcher.urls <- task

	// async write-and-parse when fetcher gets to this task
	episode_chan := make(chan *JarchiveEpisode)
	go func() {
		defer close(episode_chan)
		episode_html := <-bytes_chan

		filepath := EpisodeHtmlPath(match)
		err := os.WriteFile(filepath, episode_html, 0644)
		if err != nil {
			fetcher.errors <- err
			return
		}

		episode, err := ParseEpisodeHtml(episode_html)
		if err != nil {
			fetcher.errors <- err
			return
		}
		episode_chan <- episode
	}()

	// caller awaits on the JarchiveSeason that is produced
	return episode_chan
}

func (fetcher *fetcher) Errors() <-chan error {
	return fetcher.errors
}

func (fetcher *fetcher) Close() {
	close(fetcher.errors)
	close(fetcher.urls)
	fetcher.ticker.Stop()
}

// // The UserAgent we identify ourselves as.  Follows sec 3.7 of RFC 1945.
func DefaultUserAgent() string {
	var major, minor int
	if len(CURRENT_VERSION) > 1 {
		minor = CURRENT_VERSION[1]
	}
	if len(CURRENT_VERSION) > 0 {
		major = CURRENT_VERSION[0]
	}

	// I append a special signature at the end of legitimate requests from the original author.  This signature is not published on github.
	return fmt.Sprintf("q-party/golang/%d.%d +https://github.com/kevindamm/q-party (FORKED)", major, minor)
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

func IndexHtmlPath() string {
	return "index.html"
}

func SeasonIndexHtmlPath(season_slug schema.SeasonSlug) string {
	return path.Join("season", fmt.Sprintf("%s.html", season_slug))
}

func EpisodeHtmlPath(match schema.MatchNumber) string {
	return path.Join("episode", fmt.Sprintf("%d.html", match))
}
