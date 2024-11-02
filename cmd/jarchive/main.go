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
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

func main() {
	out_path := flag.String("out", "./.data",
		"path where converted and created games are written")
	flag.Usage = func() {
		fmt.Printf("%s command episode# [flags]\n", os.Args[0])
		fmt.Println("  where")
		fmt.Println("    command is either 'fetch' or 'convert'")
		fmt.Println("    episode# is the index ID for the episode")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
	}

	switch os.Args[1] {
	case "fetch":
		episode_id := os.Args[2]
		episodes_path := path.Join(*out_path, "episodes")
		filename := fmt.Sprintf("%s.html", episode_id)
		err := FetchEpisode(episode_id, path.Join(episodes_path, filename))
		if err != nil {
			log.Fatal(err)
		}

	case "convert":
		episode_id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("expected integer for episode id (got '%s')", os.Args[2])
		}
		ep_path := path.Join(*out_path, "episodes")
		filename := fmt.Sprintf("%d.html", episode_id)
		reader, err := os.Open(path.Join(ep_path, filename))
		if err != nil {
			log.Fatalf("could not open '%s'\n%s", ep_path, err)
		}
		defer reader.Close()

		// filename = filename[:len(filename)-4] + "json"
		// filepath := path.Join(*out_path, "episodes", filename)
		// writer, err := os.Create(filepath)
		// if err != nil {
		// 	log.Fatalf("could not create json file for episode %s\n%s", filename, err)
		// }
		// defer writer.Close()

		err = ConvertEpisode(filename, reader, os.Stdout)
		if err != nil {
			log.Fatalf("could not convert episode %s\n%s", filename, err)
		}
	}
}
