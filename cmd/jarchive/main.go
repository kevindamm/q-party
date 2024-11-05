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
)

func main() {
	data_path := flag.String("data", "./.data",
		"path where converted and created games are written")
	flag.Usage = func() {
		fmt.Printf("%s command episode# [flags]\n", path.Base(os.Args[0]))
		fmt.Println("  where")
		fmt.Println("    command is either 'fetch' or 'convert'")
		fmt.Println("    episode# is the index ID for the episode")
		fmt.Println()
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	switch flag.Arg(0) {

	case "fetch":
		episodes_path := path.Join(*data_path, "episodes")
		jeid := MustParseJEID(flag.Arg(1))
		filepath := path.Join(episodes_path, jeid.HTML())

		err := FetchEpisode(jeid, filepath)
		if err != nil {
			log.Fatal(err)
		}

	case "convert":
		jeid := MustParseJEID(flag.Arg(1))
		html_path := path.Join(*data_path, "episodes", jeid.HTML())

		reader, err := os.Open(html_path)
		if err != nil {
			log.Fatal("could not open episode", html_path,
				"\n", err)
		}
		defer reader.Close()

		json_path := path.Join(*data_path, "episodes", jeid.JSON())
		writer, err := os.Create(json_path)
		if err != nil {
			log.Fatal("could not create json file for episode", json_path,
				"\n", err)
		}
		defer writer.Close()

		err = ConvertEpisode(jeid, reader, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}

	default:
		flag.Usage()
	}
}
