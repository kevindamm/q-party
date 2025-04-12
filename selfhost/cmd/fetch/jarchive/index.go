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
	"os"
	"path"
	"slices"
	"sync"

	"github.com/kevindamm/q-party/schema"
	"golang.org/x/sync/errgroup"
)

// Concurrency-safe interface to a collection of seasons, episodes and categories.
type JarchiveIndex interface {
	Version() []int

	AddSeasonMetadata(*schema.SeasonMetadata) error
	AddEpisodeMetadata(*schema.EpisodeMetadata) error
	AddCategoryMetadata(*schema.CategoryMetadata) error
	AddEpisode(*JarchiveEpisode) error

	Extend(JarchiveIndex) error

	GetSeasonList() []schema.SeasonSlug
	GetSeasonInfo(schema.SeasonSlug) schema.SeasonMetadata
	GetEpisodeList(schema.SeasonSlug) []schema.MatchNumber
	GetEpisodeInfo(schema.MatchNumber) schema.EpisodeMetadata
	GetCategoryCalendar(schema.CategoryName) []schema.CategoryAired

	Fetch() error
	ParseHtml([]byte) error
	ParseJSONLines(io.Reader) error
	WriteJSONLines(io.Writer) error
}

func LoadLocalIndex(datapath string) (JarchiveIndex, error) {
	index_path := path.Join(datapath, "jarchive.jsonl")
	file, err := os.Open(index_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	index := new(index)
	err = index.ParseJSONLines(file)
	if err != nil {
		return nil, err
	}
	return index, err
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
	// TODO
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

func (this *index) Extend(other JarchiveIndex) error {
	if !slices.Equal(this.SemVer[:3], other.Version()[:3]) {
		return fmt.Errorf("incompatbile versions merging JarchiveIndex %v != %v",
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
			// TODO merge season info piecewise? but expect there won't be competing partial writes.
			this.Seasons[key] = other_index.Seasons[key]
		}
	}
	// TODO inspect other.(Episodes|Categories), nonempty ones are written
	return nil
}

func (index *index) ParseJSONLines(reader io.Reader) error {
	var group errgroup.Group
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		season := new(schema.SeasonMetadata)
		line := scanner.Bytes()
		group.Go(func() error {
			err := json.Unmarshal(line, season)
			if err != nil {
				return err
			}
			return index.AddSeasonMetadata(season)
		})
	}

	if err := group.Wait(); err != nil {
		return err
	}
	return nil
}

func (index *index) WriteJSONLines(writer io.Writer) error {
	// TODO
	return nil
}

func (index *index) GetSeasonList() []schema.SeasonSlug {
	// TODO
	return []schema.SeasonSlug{}
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

const LIST_SEASONS_URL = "https://j-archive.com/listseasons.php"

func (index *index) Fetch() error {

	// TODO
	return nil
}

func (index *index) ParseHtml(data []byte) error {

	// TODO
	return nil
}
