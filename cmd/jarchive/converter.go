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

func ConvertGamesInDir(dir_path string) <-chan *JArchiveEpisode {
	file, err := os.Open(dir_path)
	if err != nil {
		log.Fatal(err)
	}

	channel := make(chan *JArchiveEpisode)
	go func() {
		defer file.Close()
		defer close(channel)

		names, err := file.Readdirnames(0)
		if err != nil {
			log.Printf("error reading episodes directory: %s\n", err)
			return
		}

		for _, name := range names {
			//reader, err := os.Open(name)
			if err != nil {
				log.Print(err)
				continue
			}
			//channel <- ParseEpisode(name, reader)
			fmt.Println(name)
		}
	}()

	return channel
}

func ConvertAllEpisodes(convert_path string, out_path string) {
	err := os.MkdirAll(out_path, 0755)
	if err != nil {
		log.Fatalf("failed to create directory for converted episodes %s\n%s", out_path, err)
	}

	for jgame := range ConvertGamesInDir(convert_path) {
		filename := fmt.Sprintf("%d.json", jgame.ShowNumber)
		filepath := path.Join(out_path, filename)
		outfile, err := os.Create(filepath)
		if err != nil {
			log.Fatalf("failed to create file '%s': %s", filepath, err)
		}

		jgame_json, err := json.MarshalIndent(jgame, "", "  ")
		if err != nil {
			log.Fatalf("failed to encode episode %d", jgame.ShowNumber)
		}
		outfile.Write(jgame_json)
	}
}
