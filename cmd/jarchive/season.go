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
// github:kevindamm/q-party/cmd/jarchive/season.go

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type JArchiveSeason struct {
	Season JSID           `json:"season"`
	Name   string         `json:"name"`
	Aired  AiredDateRange `json:"aired"`
	Count  int            `json:"count"`

	Episodes map[JEID]JArchiveEpisodeMetadata `json:"episodes"`
}

func (season *JArchiveSeason) LoadSeasonMetadata(reader io.Reader) error {
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

	season.Episodes = make(map[JEID]JArchiveEpisodeMetadata)

	matches := reEpisodeLink.FindAllSubmatch(bytes, -1)
	for _, match := range matches {
		episode := JArchiveEpisodeMetadata{}
		episode.JEID = MustParseJEID(string(match[1]))
		taped := reTapedDate.FindSubmatch(match[2])
		if taped != nil {
			episode.Taped = parseTimeYYYYMMDD(taped[1], taped[2], taped[3])
		}
		aired := reAiredDate.FindSubmatch(match[2])
		if aired != nil {
			episode.Aired = parseTimeYYYYMMDD(aired[1], aired[2], aired[3])
		}
		season.Episodes[episode.JEID] = episode
	}

	return nil
}

func (season JArchiveSeason) FetchIndex(filepath string) error {
	url := fmt.Sprintf("https://j-archive.com/showseason.php?season=%s", season.Season)
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

// Unique (sometimes numeric) identifier for seasons in the archive.
type JSID string

// Returns a non-zero value if this season is part of regular play,
// zero otherwise.
func (id JSID) RegularSeason() int {
	number, err := strconv.Atoi(string(id))
	if err != nil {
		return 0
	}
	return number
}

const all_seasons string = `[
{"season": "41", "name": "Season 41", "aired": {"from": "2024/09/09", "until": "2025/07/25"}, "count": 37},
{"season": "cwcpi", "name": "Champions Wildcard play-in games", "aired": {"from": "2024/01/12", "until": "2024/02/02"}, "count": 3},
{"season": "40", "name": "Season 40", "aired": {"from": "2023/09/11", "until": "2024/07/26"}, "count": 230},
{"season": "jm", "name": "Jeopardy! Masters", "aired": {"from": "2023/05/08", "until": "2023/05/24"}, "count": 38},
{"season": "pcj", "name": "Primetime Celebrity Jeopardy!", "aired": {"from": "2022/09/25", "until": "2024/02/02"}, "count": 26},
{"season": "39", "name": "Season 39", "aired": {"from": "2022/09/12", "until": "2023/07/28"}, "count": 230},
{"season": "ncc", "name": "Jeopardy! National College Championship", "aired": {"from": "2022/02/08", "until": "2022/02/22"}, "count": 18},
{"season": "38", "name": "Season 38", "aired": {"from": "2021/09/13", "until": "2022/07/29"}, "count": 230},
{"season": "37", "name": "Season 37", "aired": {"from": "2020/09/14", "until": "2021/08/13"}, "count": 230},
{"season": "goattournament", "name": "Jeopardy!: The Greatest of All Time", "aired": {"from": "2020/01/07", "until": "2020/01/14"}, "count": 8},
{"season": "36", "name": "Season 36", "aired": {"from": "2019/09/09", "until": "2020/06/12"}, "count": 190},
{"season": "35", "name": "Season 35", "aired": {"from": "2018/09/10", "until": "2019/07/26"}, "count": 230},
{"season": "34", "name": "Season 34", "aired": {"from": "2017/09/11", "until": "2018/07/27"}, "count": 230},
{"season": "33", "name": "Season 33", "aired": {"from": "2016/09/12", "until": "2017/07/28"}, "count": 230},
{"season": "32", "name": "Season 32", "aired": {"from": "2015/09/14", "until": "2016/07/29"}, "count": 230},
{"season": "31", "name": "Season 31", "aired": {"from": "2014/09/15", "until": "2015/07/31"}, "count": 230},
{"season": "30", "name": "Season 30", "aired": {"from": "2013/09/16", "until": "2014/08/01"}, "count": 230},
{"season": "29", "name": "Season 29", "aired": {"from": "2012/09/17", "until": "2013/08/02"}, "count": 230},
{"season": "28", "name": "Season 28", "aired": {"from": "2011/09/19", "until": "2012/08/03"}, "count": 230},
{"season": "27", "name": "Season 27", "aired": {"from": "2010/09/13", "until": "2011/07/29"}, "count": 230},
{"season": "26", "name": "Season 26", "aired": {"from": "2009/09/14", "until": "2010/07/30"}, "count": 230},
{"season": "25", "name": "Season 25", "aired": {"from": "2008/09/08", "until": "2009/07/24"}, "count": 230},
{"season": "24", "name": "Season 24", "aired": {"from": "2007/09/10", "until": "2008/07/25"}, "count": 230},
{"season": "23", "name": "Season 23", "aired": {"from": "2006/09/11", "until": "2007/07/27"}, "count": 230},
{"season": "22", "name": "Season 22", "aired": {"from": "2005/09/12", "until": "2006/07/28"}, "count": 230},
{"season": "21", "name": "Season 21", "aired": {"from": "2004/09/06", "until": "2005/07/22"}, "count": 230},
{"season": "20", "name": "Season 20", "aired": {"from": "2003/09/08", "until": "2004/07/23"}, "count": 230},
{"season": "19", "name": "Season 19", "aired": {"from": "2002/09/02", "until": "2003/07/18"}, "count": 230},
{"season": "18", "name": "Season 18", "aired": {"from": "2001/09/03", "until": "2002/07/19"}, "count": 229},
{"season": "17", "name": "Season 17", "aired": {"from": "2000/09/04", "until": "2001/07/20"}, "count": 230},
{"season": "16", "name": "Season 16", "aired": {"from": "1999/09/06", "until": "2000/07/21"}, "count": 230},
{"season": "15", "name": "Season 15", "aired": {"from": "1998/09/07", "until": "1999/07/23"}, "count": 230},
{"season": "bbab", "name": "Battle of the Bay Area Brains", "aired": {"from": "1998/05/03", "until": "1998/05/03"}, "count": 1},
{"season": "14", "name": "Season 14", "aired": {"from": "1997/09/01", "until": "1998/07/17"}, "count": 230},
{"season": "13", "name": "Season 13", "aired": {"from": "1996/09/02", "until": "1997/07/18"}, "count": 218},
{"season": "12", "name": "Season 12", "aired": {"from": "1995/09/04", "until": "1996/07/19"}, "count": 229},
{"season": "11", "name": "Season 11", "aired": {"from": "1994/09/05", "until": "1995/07/21"}, "count": 214},
{"season": "10", "name": "Season 10", "aired": {"from": "1993/09/06", "until": "1994/07/22"}, "count": 230},
{"season": "9", "name": "Season 9", "aired": {"from": "1992/09/07", "until": "1993/07/23"}, "count": 229},
{"season": "8", "name": "Season 8", "aired": {"from": "1991/09/02", "until": "1992/07/17"}, "count": 194},
{"season": "7", "name": "Season 7", "aired": {"from": "1990/09/03", "until": "1991/07/19"}, "count": 212},
{"season": "superjeopardy", "name": "Super Jeopardy!", "aired": {"from": "1990/06/16", "until": "1990/09/08"}, "count": 13},
{"season": "6", "name": "Season 6", "aired": {"from": "1989/09/04", "until": "1990/07/20"}, "count": 202},
{"season": "5", "name": "Season 5", "aired": {"from": "1988/09/05", "until": "1989/07/21"}, "count": 215},
{"season": "4", "name": "Season 4", "aired": {"from": "1987/09/07", "until": "1988/07/22"}, "count": 221},
{"season": "3", "name": "Season 3", "aired": {"from": "1986/09/08", "until": "1987/07/24"}, "count": 214},
{"season": "2", "name": "Season 2", "aired": {"from": "1985/09/09", "until": "1986/06/06"}, "count": 179},
{"season": "1", "name": "Season 1", "aired": {"from": "1984/09/10", "until": "1985/06/07"}, "count": 164}]`

//{"season": "trebekpilots", "name": "Trebek pilots", "aired": {"from": "1983/09/18", "until": "1984/01/09"}, "count": 2}]`
