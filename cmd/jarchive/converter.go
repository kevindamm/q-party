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
// github:kevindamm/q-party/cmd/jarchive/converter.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

// The de-normed representation as found in some datasets, e.g. on Kaggle.
type JArchiveChallenge struct {
	Category string            `json:"category"`
	AirDate  `json:"air_date"` // YYYY-MM-DD

	Value    DollarValue  `json:"value"` // '$' (\d+)
	Question string       `json:"question"`
	Answer   string       `json:"answer"` // excluding "what is..." preface
	Round    EpisodeRound `json:"round"`
}

func ConvertGamesInDir(dir_path string) <-chan *JArchiveEpisode {
	channel := make(chan *JArchiveEpisode)

	go func() {
		// for file, _, _ in os.WalkDir(path)
		//   convert and write to channel
		channel <- ParseEpisode("...") // parse(filepath)
		// ...
		close(channel)
	}()

	return channel
}

func ConvertAllSeasons(all_seasons []JArchiveSeason, convert_path string, out_path string) error {
	for _, season := range all_seasons {
		fmt.Println(season.Season)

		season_path := path.Join(convert_path, season.Season)
		err := os.MkdirAll(season_path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create output path for season '%s' episodes: %s", season_path, err)
		}

		for jgame := range ConvertGamesInDir(season_path) {
			filename := fmt.Sprintf("%d.json", jgame.ShowNumber)
			filepath := path.Join(season_path, filename)
			outfile, err := os.Create(filepath)
			if err != nil {
				log.Fatalf("failed to create file '%s': %s", filepath, err)
			}
			jgame_json, err := json.MarshalIndent(jgame, "", "  ")
			if err != nil {
				log.Fatalf("failed to write JSON for episode %s/%s: %s",
					season.Season, filename, err)
			}
			outfile.Write(jgame_json)
		}
	}

	return nil
}

func ParseEpisode(file_path string) *JArchiveEpisode {
	episode := new(JArchiveEpisode)

	return episode
}
