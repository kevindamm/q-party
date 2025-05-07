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
// github:kevindamm/q-party/selfhost/cmd/fetch/jarchive/seasons_test.go

package main_test

import (
	_ "embed"
	"testing"

	"github.com/kevindamm/q-party/schema"
	jarchive "github.com/kevindamm/q-party/selfhost/cmd/fetch/jarchive"
)

//go:embed testdata/pcj-season.html
var CELEBJ_SEASON_HTML []byte

func TestParseSeasonPCJ(t *testing.T) {
	season_index := jarchive.NewJarchiveSeason("pcj")
	err := season_index.ParseHTML(CELEBJ_SEASON_HTML)
	if err != nil {
		t.Error(err)
	}

	expected_episodes := map[jarchive.EpisodeID]schema.MatchMetadata{
		9165: {
			TapedDate: &schema.ShowDate{Year: 2024, Month: 10, Day: 21},
			AiredDate: &schema.ShowDate{Year: 2025, Month: 2, Day: 5},
			Comments:  "2025 Primetime *Celebrity Jeopardy!* quarterfinal game 5.",
		},
	}

	for episode_id, expected_metadata := range expected_episodes {
		metadata := season_index.GetJarchiveEpisode(episode_id).Metadata()
		if metadata.TapedDate.Compare(expected_metadata.TapedDate) != 0 {
			t.Errorf("taped date for %d not expected (%s != %s)",
				episode_id, metadata.TapedDate, expected_metadata.TapedDate)
		}
		if metadata.AiredDate.Compare(expected_metadata.AiredDate) != 0 {
			t.Errorf("aired date for %d not expected (%s != %s)",
				episode_id, metadata.AiredDate, expected_metadata.AiredDate)
		}
		if metadata.Comments != expected_metadata.Comments {
			t.Errorf("comments for %d not expected (%s != %s)",
				episode_id, metadata.Comments, expected_metadata.Comments)
		}
	}
}

//go:embed testdata/goat-season.html
var GOAT_SEASON_HTML []byte

func TestParseSeasonGOAT(t *testing.T) {

}

//go:embed testdata/s25.html
var SEASON_25_HTML []byte

func TestParseSeason25(t *testing.T) {

}
