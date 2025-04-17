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
// github:kevindamm/q-party/cmd/fetch/jarchive/fetcher.go

package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var CURRENT_VERSION = []int{0, 1, 0}

// Abstraction over time-delayed fetching of URLs.
type Fetcher interface {
	// Fetches the resource and sends the populated object on the returned channel.
	Fetch(Fetchable) <-chan Fetchable

	// Set a common channel that all fetched resources will be sent to.
	// By default, the fetcher creates a new channel for each Fetch,
	// this allows callers to build a fan-in channel (and Fetch will still
	// return a channel but it will be the same one as provided here).
	// Fetcher takes ownership of closing the channel, in [Close].
	FanIn(*chan Fetchable)

	// Get the channel where errors are sent.
	Errors() <-chan error
	// Close all dependent resources (channels) this Fetcher owns.
	Close()
}

// Creates a Fetcher instance and initiates its goroutine for delayed fetching.
func NewFetcher(pause time.Duration, useragent string) Fetcher {
	useragent = strings.ReplaceAll(useragent, "'\"", "")
	useragent = strings.Trim(useragent, " \t\n")
	if len(useragent) == 0 {
		useragent = DefaultUserAgent()
	}

	fetcher := fetcher{
		useragent: useragent,
		tasks:     make(chan fetch_task),
		errors:    make(chan error),
	}
	go fetcher.start_titration(pause)

	return &fetcher
}

// Internal representation of the state, metadata and channels of the Fetcher.
type fetcher struct {
	ticker    *time.Ticker
	useragent string

	output *chan Fetchable
	tasks  chan fetch_task
	errors chan error
}

// A goroutine for periodically waiting between http.Get() requests.
func (fetcher *fetcher) start_titration(pause time.Duration) {
	// pause will be at least 2.5 seconds
	min_pause := 2500 * time.Millisecond
	fetcher.ticker = time.NewTicker(max(min_pause, pause))

	for range fetcher.ticker.C {
		select {
		case task, _ := <-fetcher.tasks:
			request, err := http.NewRequest("GET", task.url, nil)
			if err != nil {
				fetcher.errors <- err
				continue
			}
			request.Header.Set("User-Agent", fetcher.useragent)
			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				fetcher.errors <- err
				continue
			}
			defer response.Body.Close()

			bytes, err := io.ReadAll(response.Body)
			if err != nil {
				fetcher.errors <- err
			} else {
				task.output <- bytes
			}
			close(task.output)

		default: // channel is empty, keep waiting
		}
	}
}

// Represents an awaiting task being communicated on the fetcher's internal
// channel.  Each request is assumed to be a GET to the indicated URL.
type fetch_task struct {
	url    string
	output chan<- []byte
}

// First attempts to load a locally-stored representation,
// fetches a remote copy if one is not found (overridable)
// and stores both the HTML (before parsing) and JSON file
// representations relative to the Fetcher's data directory.
func (fetcher *fetcher) Fetch(resource Fetchable) <-chan Fetchable {
	bytes_chan := make(chan []byte)
	task := fetch_task{
		url:    resource.URL(),
		output: bytes_chan,
	}
	fetcher.tasks <- task

	// async write-and-parse when fetcher gets to this task
	var obj_chan chan Fetchable
	if fetcher.output != nil {
		obj_chan = *fetcher.output
	} else {
		obj_chan = make(chan Fetchable)
	}

	go func() {
		defer close(obj_chan)
		html_bytes := <-bytes_chan

		filepath := resource.FilepathHTML()
		err := os.WriteFile(filepath, html_bytes, 0644)
		if err != nil {
			fetcher.errors <- err
			return
		}

		err = resource.ParseHTML(html_bytes)
		if err != nil {
			fetcher.errors <- err
			return
		}
		obj_chan <- resource
	}()

	// caller awaits on the JarchiveIndex that is produced
	return obj_chan
}

func (fetcher *fetcher) FanIn(common_output *chan Fetchable) {
	if fetcher.output != nil {
		close(*fetcher.output)
	}
	if common_output == nil {
		return
	}
	fetcher.output = common_output
}

func (fetcher *fetcher) Errors() <-chan error {
	return fetcher.errors
}

func (fetcher *fetcher) Close() {
	close(fetcher.errors)
	close(fetcher.tasks)
	fetcher.ticker.Stop()
}

// // The UserAgent we identify ourselves as.  Follows sec 3.7 of RFC 1945.
func DefaultUserAgent() string {
	var major, minor int
	if len(CURRENT_VERSION) > 1 {
		minor = CURRENT_VERSION[1]
	}
	if len(CURRENT_VERSION) > 0 {
		major = CURRENT_VERSION[0]
	}

	// I append a special signature at the end of legitimate requests from the original author.  This signature is not published on github.
	return fmt.Sprintf("q-party/golang/%d.%d +https://github.com/kevindamm/q-party (FORKED)", major, minor)
}
