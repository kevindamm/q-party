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
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/kevindamm/q-party/schema"
	"github.com/kevindamm/q-party/selfhost/cmd/fetch"
	"golang.org/x/sync/errgroup"
)

// Concurrency-safe interface to a collection of seasons, episodes and categories.
type JarchiveIndex interface {
	fetch.Fetchable // fetches the top-level index (season slugs and metadata)

	AddSeasonMetadata(*schema.SeasonMetadata) error
	AddSeason(JarchiveSeason) error
	AddEpisode(JarchiveEpisode) error
	ExtendOverwrite(JarchiveIndex) error

	GetSeasonList() []schema.SeasonSlug
	GetSeasonMetadata(schema.SeasonSlug) schema.SeasonMetadata
	GetEpisodeList(schema.SeasonSlug) []schema.MatchNumber
	GetEpisodeInfo(schema.MatchNumber) schema.MatchMetadata

	// Returns the category and its air dates for any episodes currently indexed.
	// Does not fetch or load any files, it only searches the in-memory index.
	GetCategoryCalendar(
		schema.CategoryName,
		*schema.ShowDateRange,
	) []schema.CategoryAired
}

func NewJarchiveIndex() JarchiveIndex {
	jarchive_index := new(jarchive_index)
	jarchive_index.Seasons = make(schema.SeasonIndex)
	jarchive_index.Episodes = make(schema.MatchIndex)
	jarchive_index.Categories = make(schema.CategoryIndex)
	jarchive_index.EpisodeMatch = make(EpisodeMatchNumber)
	jarchive_index.MatchEpisode = make(MatchNumberEpisode)
	return jarchive_index
}

func LoadJarchiveJSONL(filepath string) (JarchiveIndex, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	errs := make([]error, 0)
	jarchive := NewJarchiveIndex()
	for scanner.Scan() {
		var season season_index
		err := json.Unmarshal(scanner.Bytes(), &season)
		if err != nil {
			errs = append(errs, err)
		} else {
			jarchive.AddSeason(&season)
		}
	}
	return jarchive, errors.Join(errs...)
}

type EpisodeMatchNumber map[EpisodeID]schema.MatchNumber
type MatchNumberEpisode map[schema.MatchNumber]EpisodeID

// Internal representation of [JarchiveIndex] with safe read and update access.
type jarchive_index struct {
	fetch.Fetchable

	Seasons    schema.SeasonIndex   `json:"seasons"`
	Episodes   schema.MatchIndex    `json:"episodes"`
	Categories schema.CategoryIndex `json:"categories"`

	EpisodeMatch EpisodeMatchNumber `json:"-"`
	MatchEpisode MatchNumberEpisode `json:"-"`

	lock sync.RWMutex `json:"-"`
}

func (jarchive *jarchive_index) String() string {
	return "_"
}

func (jarchive *jarchive_index) URL() string {
	return "https://j-archive.com/listseasons.php"
}

func (jarchive *jarchive_index) FilepathHtml() string {
	return "index.html"
}

func (jarchive *jarchive_index) FilepathJSON() string {
	return "jarchive.jsonl"
}

func (jarchive *jarchive_index) ParseHTML(data []byte) error {
	// The layout of this index is simple enough that it can be parsed with regex.
	// It could written with x/net/html (see parsing in episode.go) but any change
	// to the source material would need a corresponding update here, and would be
	// as much (or more!) work to make the appropriate changes.
	// <xx-meme>
	//   I don't always parse my HTML with regexes,
	//   but when I do, I comment at length about it.
	// </xx-meme>
	reSeasonMetadata := regexp.MustCompile(
		`<tr><td><a href="showseason.php\?season=([^"]+)">(?:<i>)?([^<]+)(?:</i>)?</a></td>(.*)`)
	season_matches := reSeasonMetadata.FindAllSubmatch(data, -1)
	if len(season_matches) == 0 {
		return errors.New("no seasons found in jarchive index HTML")
	}

	errs := make([]error, 0)
	reAirDateRange := regexp.MustCompile(`<td(?: class="[^"]+")?>(\d+)-(\d+)-(\d+) to (\d+)-(\d+)-(\d+)</td>`)
	for _, smatch := range season_matches {
		var season schema.SeasonMetadata
		season_slug := string(smatch[1])
		season_slug = strings.ReplaceAll(season_slug, "'\"", "")
		season_slug = strings.Trim(season_slug, " \t\n")
		season.Slug = schema.SeasonSlug(season_slug)
		season.Title = string(smatch[2])

		if season_slug == "trebekpilots" { // dubious start and end times for pilots
			season.Aired.From = &schema.ShowDate{Year: 1983, Month: 9, Day: 18}
			season.Aired.Until = &schema.ShowDate{Year: 1984, Month: 1, Day: 1}
		} else {
			date_match := reAirDateRange.FindSubmatch(smatch[3])
			if len(date_match) > 0 {
				season.Aired.From.Year = digits(date_match[1])
				season.Aired.From.Month = digits(date_match[2])
				season.Aired.From.Day = digits(date_match[3])
				season.Aired.Until.Year = digits(date_match[4])
				season.Aired.Until.Month = digits(date_match[5])
				season.Aired.Until.Day = digits(date_match[6])
			}
		}

		err := jarchive.AddSeasonMetadata(&season)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

// Writes JSONL line-separated [schema.SeasonMetadata] JSON for the index.
func (jarchive *jarchive_index) WriteJSON(writer io.WriteCloser) error {
	defer writer.Close()
	for _, season := range jarchive.Seasons {
		json_bytes, err := json.Marshal(season)
		if err != nil {
			return err
		}
		json_bytes = append(json_bytes, byte('\n'))
		_, err = writer.Write(json_bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

// Reads JSONL
func (jarchive *jarchive_index) LoadJSON(reader io.ReadCloser) error {
	defer reader.Close()
	var group errgroup.Group
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Bytes()
		var season schema.SeasonMetadata

		group.Go(func() error {
			err := json.Unmarshal(line, &season)
			if err != nil {
				return err
			}
			return jarchive.AddSeasonMetadata(&season)
		})
	}

	if err := group.Wait(); err != nil {
		return err
	}
	return nil
}

func (jarchive *jarchive_index) AddSeasonMetadata(season *schema.SeasonMetadata) error {
	var err error
	jarchive.lock.Lock()
	defer jarchive.lock.Unlock()

	if _, found := jarchive.Seasons[season.Slug]; found {
		log.Printf("WARNING AddSeasonMetadata called with the same slug '%s' multiple times", season.Slug)
		err = fmt.Errorf("multiple insertion season %s metadata", season.Slug)
	}
	jarchive.Seasons[season.Slug] = season
	return err
}

func (jarchive *jarchive_index) AddSeason(season JarchiveSeason) error {
	var err error
	jarchive.lock.Lock()
	defer jarchive.lock.Unlock()

	if _, found := jarchive.Seasons[season.SeasonSlug()]; found {
		log.Printf("WARNING inserting (more than once) metadata for season '%s'",
			season.SeasonSlug())
		err = fmt.Errorf("multiple insertion season %s data", season.SeasonSlug())
	}
	jarchive.Seasons[season.SeasonSlug()] = season.Metadata()
	return err
}

func (jarchive *jarchive_index) AddEpisode(episode JarchiveEpisode) error {
	var err error
	jarchive.lock.Lock()
	defer jarchive.lock.Unlock()

	if _, found := jarchive.Episodes[episode.Metadata().MatchNumber]; found {
		metadata := episode.Metadata()
		log.Printf("WARNING inserting (more than once) metadata for episode (%s/%d)",
			metadata.SeasonSlug,
			metadata.MatchNumber)
		err = fmt.Errorf("multiple insertion episode (%s) %d",
			metadata.SeasonSlug,
			metadata.MatchNumber)
	}
	jarchive.Episodes[episode.Metadata().MatchNumber] = episode.Metadata()
	return err
}

func (jarchive *jarchive_index) ExtendOverwrite(other JarchiveIndex) error {
	other_index, ok := other.(*jarchive_index)
	if !ok {
		return fmt.Errorf("invalid type of other %T (expected an index)", other)
	}
	if len(other_index.Seasons) > 0 ||
		len(other_index.Episodes) > 0 ||
		len(other_index.Categories) > 0 {
		// We could make this lock more fine-grained but this structure is seldom
		// updated (and mostly during system startup); correctness >> optimization.
		jarchive.lock.Lock()
		defer jarchive.lock.Unlock()
	}

	if len(other_index.Seasons) > 0 {
		for key := range other_index.Seasons {
			if _, found := jarchive.Seasons[key]; found {
				log.Printf("WARNING found multiple updates to season %s (duplicate Slug?)", key)
			}
			jarchive.Seasons[key] = other_index.Seasons[key]
		}
	}

	if len(other_index.Episodes) > 0 {
		for key := range other_index.Episodes {
			if _, found := jarchive.Episodes[key]; found {
				log.Printf("WARNING found multiple updates to episode %d (duplicate MatchNumber?)", key)
			}
			jarchive.Episodes[key] = other_index.Episodes[key]
		}
	}

	if len(other_index.Categories) > 0 {
		for key := range other_index.Categories {
			if _, found := jarchive.Categories[key]; found {
				log.Printf("WARNING found multiple updates to category %s (duplicate CategoryName?)", key)
			}
			jarchive.Categories[key] = other_index.Categories[key]
		}
	}

	return nil
}

func (jarchive *jarchive_index) GetSeasonList() []schema.SeasonSlug {
	slugs := make([]schema.SeasonSlug, len(jarchive.Seasons))
	i := 0
	for slug := range jarchive.Seasons {
		slugs[i] = slug
		i++
	}
	slices.Sort(slugs)
	return slugs
}

func (jarchive *jarchive_index) GetSeasonMetadata(key schema.SeasonSlug) schema.SeasonMetadata {
	jarchive.lock.RLock()
	defer jarchive.lock.RUnlock()
	return *jarchive.Seasons[key]
}

func (jarchive *jarchive_index) GetEpisodeList(season_slug schema.SeasonSlug) []schema.MatchNumber {
	jarchive.lock.RLock()
	defer jarchive.lock.RUnlock()
	matches := make([]schema.MatchNumber, 0)
	for match_num, episode := range jarchive.Episodes {
		if episode.SeasonSlug == season_slug {
			matches = append(matches, match_num)
		}
	}
	return matches
}

func (jarchive *jarchive_index) GetEpisodeInfo(key schema.MatchNumber) schema.MatchMetadata {
	jarchive.lock.RLock()
	defer jarchive.lock.RUnlock()
	return *jarchive.Episodes[key]
}

func (jarchive *jarchive_index) GetCategoryCalendar(key schema.CategoryName, scope *schema.ShowDateRange) []schema.CategoryAired {
	jarchive.lock.RLock()
	defer jarchive.lock.RUnlock()
	if scope == nil {
		// Retrieve from entire known date range.
		return jarchive.Categories[key]
	}

	response := make([]schema.CategoryAired, 0)
	for _, aircat := range jarchive.Categories[key] {
		if scope.Contains(aircat.Aired) {
			response = append(response, aircat)
		}
	}
	return response
}
