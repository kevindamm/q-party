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
// github:kevindamm/q-party/cmd/jarchive/main.go

package main

import (
	"encoding/json"
	"flag"
	"log"
)

func main() {
	convert_path := flag.String("jgames_path", "",
		"path to collection of j-archive.com games")
	output_path := flag.String("out_path", "./.data",
		"path where converted and created games are written")

	flag.Parse()

	all_seasons := make([]JArchiveSeason, 0, 50)
	err := json.Unmarshal([]byte(seasons), &all_seasons)
	if err != nil {
		log.Fatalf("failed to load season metadata: %s", err)
	}

	if len(*convert_path) > 0 {
		err := ConvertAllSeasons(all_seasons, *convert_path, *output_path)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = FetchAllSeasons(all_seasons, *output_path)
	if err != nil {
		log.Fatal(err)
	}
}
