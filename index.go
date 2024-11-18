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
// github:kevindamm/q-party/index.go

package qparty

import "encoding/json"

type JArchiveIndex struct {
	Version    []uint                      `json:"version,omitempty"`
	Seasons    map[SeasonID]SeasonMetadata `json:"seasons"`
	Categories map[string]CategoryMetadata `json:"categories"`
	Episodes   map[EpisodeID]EpisodeStats  `json:"episodes"`
}

func LoadJArchiveIndex(jarchive_json []byte) (*JArchiveIndex, error) {
	index := new(JArchiveIndex)
	err := json.Unmarshal(jarchive_json, index)
	if err != nil {
		return nil, err
	}
	// TODO create additional (in-memory) indexes for easier retrieval
	// TODO also cache recently loaded episodes?  that can be done separately in main, though
	return index, nil
}
