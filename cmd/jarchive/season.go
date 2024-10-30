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
)

type JArchiveSeason struct {
	Season string         `json:"season"`
	Name   string         `json:"name"`
	Aired  AiredDateRange `json:"aired"`
	Count  int            `json:"count"`
}

func (season JArchiveSeason) LocalEpisodes() {
	// ...
}

func (season JArchiveSeason) FetchIndex(filename string) error {
	url := fmt.Sprintf("https://j-archive.com/showseason.php?season=%s", season.Season)
	log.Print("Fetching ", url, "  -> ", filename)

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	os.WriteFile(filename, data, 0644)

	return nil
}
