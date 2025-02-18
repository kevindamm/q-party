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
	"regexp"
	"time"

	qparty "github.com/kevindamm/q-party"
)

func ParseSeasonMetadata(reader io.Reader, season *qparty.SeasonMetadata) error {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	reSeasonName := regexp.MustCompile(`<h2 class="season">(.*)</h2>`)
	match := reSeasonName.FindSubmatch(bytes)
	if len(match) > 0 {
		season.Name = string(match[1])
	}

	reEpisodeLink := regexp.MustCompile(`"showgame\.php\?game_id=(\d+)"(.*)\n`)
	matches := reEpisodeLink.FindAllSubmatch(bytes, -1)
	for range matches {
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
