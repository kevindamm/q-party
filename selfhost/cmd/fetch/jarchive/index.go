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
// github:kevindamm/q-party/cmd/fetch/jarchive/index.go

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"slices"
	"sync"

	"github.com/kevindamm/q-party/schema"
	"golang.org/x/sync/errgroup"
)

var CURRENT_VERSION = []int{0, 0, 0}

// Concurrency-safe interface to a collection of seasons, episodes and categories.
type JarchiveIndex interface {
	Version() []int

	AddSeasonMetadata(*schema.SeasonMetadata) error
	AddEpisodeMetadata(*schema.EpisodeMetadata) error
	AddCategoryMetadata(*schema.CategoryMetadata) error
	AddEpisode(*JarchiveEpisode) error

	GetSeasonList() []schema.SeasonSlug
	GetSeasonInfo(schema.SeasonSlug) schema.SeasonMetadata
	GetEpisodeList(schema.SeasonSlug) []schema.MatchNumber
	GetEpisodeInfo(schema.MatchNumber) schema.EpisodeMetadata
	GetCategoryCalendar(schema.CategoryName) []schema.CategoryAired

	ExtendOverwrite(JarchiveIndex) error
	WriteJSONLines(io.Writer) error
}

func NewJarchiveIndex(version []int) JarchiveIndex {
	index := new(index)
	index.SemVer = version
	index.Seasons = make(schema.SeasonIndex)
	index.Episodes = make(schema.EpisodeIndex)
	index.Categories = make(schema.CategoryIndex)
	return index
}

func ParseJSONLines(reader io.Reader) (JarchiveIndex, error) {
	var group errgroup.Group
	index := NewJarchiveIndex(CURRENT_VERSION)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Bytes()
		season := new(schema.SeasonMetadata)

		group.Go(func() error {
			err := json.Unmarshal(line, season)
			if err != nil {
				return err
			}
			return index.AddSeasonMetadata(season)
		})
	}

	if err := group.Wait(); err != nil {
		return index, err
	}
	return index, nil
}

func LoadLocalIndex(datapath string) (JarchiveIndex, error) {
	index_path := path.Join(datapath, "jarchive.jsonl")
	file, err := os.Open(index_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	index, err := ParseJSONLines(file)
	if err != nil {
		return nil, err
	}
	return index, err
}

// Parse the contents of the HTML-formatted byte-slice into a [JarchiveIndex],
// representing metadata about all seasons in the archive.
func ParseIndexHtml(data []byte) (JarchiveIndex, error) {
	// The layout of this index is simple enough that it can be parsed with regex.
	// It could written with x/net/html (see parsing in episode.go) but any change
	// to the source material would need a corresponding update here, and would be
	// and be as much (or more!) work to make the appropriate changes.
	// <xx-meme>
	//   I don't always parse my HTML with regexes,
	//   but when I do, I comment at length about it.
	// </xx-meme>
	reSeasonMetadata := regexp.MustCompile(
		`<tr><td><a href="showseason.php\?season=([^"]+)">(?:<i>)?([^<]+)(?:</i>)?</a></td>(.*)`)
	season_matches := reSeasonMetadata.FindAllSubmatch(data, -1)
	if len(season_matches) == 0 {
		return nil, errors.New("no seasons found in jarchive index HTML")
	}

	index := NewJarchiveIndex(CURRENT_VERSION)
	errs := make([]error, 0)
	reAirDateRange := regexp.MustCompile(`<td(?: class="[^"]+")?>(\d+)-(\d+)-(\d+) to (\d+)-(\d+)-(\d+)</td>`)
	for _, smatch := range season_matches {
		season := new(schema.SeasonMetadata)
		season.Season.Slug = schema.SeasonSlug(string(smatch[1]))
		season.Season.Title = string(smatch[2])

		date_match := reAirDateRange.FindSubmatch(smatch[3])
		if len(date_match) > 0 {
			season.Aired.From.Year = atoi(date_match[1])
			season.Aired.From.Month = atoi(date_match[2])
			season.Aired.From.Day = atoi(date_match[3])
			season.Aired.Until.Year = atoi(date_match[4])
			season.Aired.Until.Month = atoi(date_match[5])
			season.Aired.Until.Day = atoi(date_match[6])
		}
		err := index.AddSeasonMetadata(season)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return index, errors.Join(errs...)
}

func (index *index) WriteJSONLines(writer io.Writer) error {
	// TODO
	return nil
}

// Internal representation of [JarchiveIndex] with safe read and update access.
type index struct {
	SemVer     []int                `json:"version"`
	Seasons    schema.SeasonIndex   `json:"seasons"`
	Episodes   schema.EpisodeIndex  `json:"episodes"`
	Categories schema.CategoryIndex `json:"categories"`

	lock sync.RWMutex `json:"-"`
}

func (index *index) Version() []int {
	return index.SemVer
}

func (index *index) AddSeasonMetadata(season *schema.SeasonMetadata) error {
	index.lock.Lock()
	defer index.lock.Unlock()
	index.Seasons[season.Season.Slug] = season
	// TODO is overwriting an existing key considered an error?

	return nil
}

func (index *index) AddEpisodeMetadata(episode *schema.EpisodeMetadata) error {
	index.lock.Lock()
	defer index.lock.Unlock()
	// TODO
	return nil
}

func (index *index) AddCategoryMetadata(category *schema.CategoryMetadata) error {
	index.lock.Lock()
	defer index.lock.Unlock()
	// TODO
	return nil
}

func (index *index) AddEpisode(episode *JarchiveEpisode) error {
	index.lock.Lock()
	defer index.lock.Unlock()
	// TODO
	return nil
}

func (this *index) ExtendOverwrite(other JarchiveIndex) error {
	if !slices.Equal(this.SemVer[:3], other.Version()[:3]) {
		log.Printf("WARNING incompatbile versions merging JarchiveIndex %v != %v",
			this.SemVer, other.Version())
	}

	other_index, ok := other.(*index)
	if !ok {
		return errors.New("invalid type of other (expected an index)")
	}
	if len(other_index.Seasons) > 0 {
		this.lock.Lock()
		defer this.lock.Unlock()
		for key := range other_index.Seasons {
			// expect there won't be competing partial writes
			this.Seasons[key] = other_index.Seasons[key]
		}
	}
	// TODO inspect other.(Episodes|Categories), nonempty ones are written
	return nil
}

func (index *index) GetSeasonList() []schema.SeasonSlug {
	slugs := make([]schema.SeasonSlug, len(index.Seasons))
	i := 0
	for _, season := range index.Seasons {
		slugs[i] = season.Season.Slug
		i++
	}
	return slugs
}

func (index *index) GetSeasonInfo(key schema.SeasonSlug) schema.SeasonMetadata {
	index.lock.RLock()
	defer index.lock.RUnlock()
	return *index.Seasons[key]
}

func (index *index) GetEpisodeList(schema.SeasonSlug) []schema.MatchNumber {
	// TODO
	return []schema.MatchNumber{}
}

func (index *index) GetEpisodeInfo(key schema.MatchNumber) schema.EpisodeMetadata {
	index.lock.RLock()
	defer index.lock.RUnlock()
	return *index.Episodes[key]
}
func (index *index) GetCategoryCalendar(key schema.CategoryName) []schema.CategoryAired {
	index.lock.RLock()
	defer index.lock.RUnlock()
	return index.Categories[key]
}
