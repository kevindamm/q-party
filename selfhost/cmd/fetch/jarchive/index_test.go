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
// github:kevindamm/q-party/selfhost/cmd/fetch/jarchive/index_test.go

package main_test

import (
	_ "embed"
	"slices"
	"testing"

	"github.com/kevindamm/q-party/schema"
	jarchive "github.com/kevindamm/q-party/selfhost/cmd/fetch/jarchive"
)

//go:embed testdata/jarchive_index.html
var JARCHIVE_INDEX_HTML []byte

func TestParseIndexHTML(t *testing.T) {
	jarchive_index := jarchive.NewJarchiveIndex()
	err := jarchive_index.ParseHTML(JARCHIVE_INDEX_HTML)
	if err != nil {
		t.Error("Failed to parse jarchive index:", err)
		return
	}

	expected_seasons := []schema.SeasonSlug{
		schema.SeasonSlug("26"),
		schema.SeasonSlug("28"),
		schema.SeasonSlug("23"),
		schema.SeasonSlug("29"),
		schema.SeasonSlug("27"),
		schema.SeasonSlug("bbab"),
		schema.SeasonSlug("superjeopardy"),
	}
	season_list := jarchive_index.GetSeasonList()
	for _, season_slug := range season_list {
		if !slices.Contains(expected_seasons, season_slug) {
			t.Error("Found unexpected season", season_slug)
		}
	}
	for _, season_slug := range expected_seasons {
		if !slices.Contains(season_list, season_slug) {
			t.Error("Expected season not found", season_slug)
		}
	}

	for _, slug := range jarchive_index.GetSeasonList() {
		season := jarchive_index.GetSeasonMetadata(slug)
		if season.Aired.From.Year == 0 ||
			season.Aired.From.Month == 0 ||
			season.Aired.From.Day == 0 {
			t.Errorf("from date (%v) appears invalid", season.Aired.From)
		}
		if season.Aired.Until.Year == 0 ||
			season.Aired.Until.Month == 0 ||
			season.Aired.Until.Day == 0 {
			t.Errorf("until date (%v) appears invalid", season.Aired.Until)
		}
	}
}
