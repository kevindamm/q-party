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
// github:kevindamm/q-party/selfhost/cmd/fetch/jarchive/episodes.go

package main

import (
	"fmt"
	"sync"

	"github.com/kevindamm/q-party/schema"
)

type EpisodeID int

func EpisodeURL(id EpisodeID) string {
	return fmt.Sprintf("https://j-archive.com/showgame.php?game_id=%d", id)
}

type EpisodeMatchNumber map[EpisodeID]schema.MatchNumber
type MatchNumberEpisode map[schema.MatchNumber]EpisodeID

// All details of the episode, including correct answers & the contestants' bios.
type JarchiveEpisode struct {
	schema.EpisodeMetadata `json:",inline"`
	Comments               string            `json:"comments,omitempty"`
	Media                  []schema.MediaRef `json:"media,omitempty"`

	// Due to absence of archival evidence, not every episode has both single & double rounds.
	Single     *JarchiveBoard `json:"single,omitempty"`
	Double     *JarchiveBoard `json:"double,omitempty"`
	Final      *JarchiveFinal `json:"final,omitempty"`
	TieBreaker *JarchiveFinal `json:"tiebreaker,omitempty"`
}

type JarchiveEpisodeIndex struct {
	schema.EpisodeIndex
	lock sync.RWMutex
}
