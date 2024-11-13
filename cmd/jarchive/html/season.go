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
// github:kevindamm/q-party/cmd/jarchive/html/season.go

package html

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"time"

	qparty "github.com/kevindamm/q-party"
)

// Loads the season index and the metadata of each season in the index.
// Parameter [data_path] is the location of the seasons.json for metadata.
// TODO replace with instead calling to [LoadSeasonIndex()] and [LoadEpisode()].
func MustLoadAllSeasons(data_path string) *qparty.SeasonIndex {
	season_index := qparty.LoadSeasonsJSON(path.Join(data_path, "seasons.json"))

	seasons_dir := path.Join(data_path, "seasons")
	for season_id, season := range season_index.Seasons {
		json_path := path.Join(seasons_dir, fmt.Sprintf("%s.json", season_id))
		if _, err := os.Stat(json_path); os.IsNotExist(err) {
			log.Fatal("metadata file does not exist", json_path)
		}
		reader, err := os.Open(json_path)
		if err != nil {
			log.Fatal("failed to open file path", json_path, "\n", err)
		}
		err = ParseSeasonMetadata(reader, &season)
		if err != nil {
			log.Fatal("failed to parse season", season_id, "index file")
		}
	}

	return season_index
}

func ParseSeasonMetadata(reader io.Reader, season *qparty.SeasonMetadata) error {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	reEpisodeLink := regexp.MustCompile(`"showgame\.php\?game_id=(\d+)"(.*)\n`)
	reTapedDate := regexp.MustCompile(`[tT]aped\s+(\d{4})-(\d{2})-(\d{2})`)
	reAiredDate := regexp.MustCompile(`[aA]ired.*(\d{4})-(\d{2})-(\d{2})`)
	reSeasonName := regexp.MustCompile(`<h2 class="season">(.*)</h2>`)

	match := reSeasonName.FindSubmatch(bytes)
	if len(match) > 0 {
		season.Name = string(match[1])
	}

	matches := reEpisodeLink.FindAllSubmatch(bytes, -1)
	for _, match := range matches {
		episode := EpisodeMetadata{}
		// episode.JEID = MustParseJEID(string(match[1]))
		taped := reTapedDate.FindSubmatch(match[2])
		if taped != nil {
			episode.Taped = parseTimeYYYYMMDD(taped[1], taped[2], taped[3])
		}
		aired := reAiredDate.FindSubmatch(match[2])
		if aired != nil {
			episode.Aired = parseTimeYYYYMMDD(aired[1], aired[2], aired[3])
		}
		season.Count += 1
	}
	return nil
}

func FetchSeasonIndexHTML(jsid qparty.SeasonID, filepath string) error {
	url := SeasonURL(jsid)
	log.Print("Fetching ", url, "  -> ", filepath)

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	return nil
}

func SeasonURL(id qparty.SeasonID) string {
	return fmt.Sprintf("https://j-archive.com/showseason.php?season=%s", id)
}
