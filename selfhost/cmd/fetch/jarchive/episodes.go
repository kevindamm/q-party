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
	"io"
	"log"
	"path"

	"github.com/kevindamm/q-party/schema"
	"github.com/kevindamm/q-party/selfhost/cmd/fetch"
)

type EpisodeID int

// All details of the episode, including correct answers & the contestants' bios.
type JarchiveEpisode interface {
	fetch.Fetchable
	// Properties are set via ParseHTML or LoadJSON (of Fetchable interface)

	// Property getters
	MatchNumber() schema.MatchNumber // may be 0 for unknown

}

func NewEpisode(id EpisodeID) JarchiveEpisode {
	if id == 0 {
		log.Fatal("zero value invalid for episode ID (equiv to UNKNOWN episode)")
	}
	return &episode{
		EpisodeID: id,
	}
}

func UnknownEpisode() JarchiveEpisode {
	return &episode{EpisodeID: 0}
}

type episode struct {
	schema.EpisodeMetadata `json:",inline"`

	EpisodeID EpisodeID         `json:"episode_id,omitempty"`
	Comments  string            `json:"comments,omitempty"`
	Media     []schema.MediaRef `json:"media,omitempty"`

	// Due to absence of archival evidence,
	// not every episode has both single & double rounds.
	Single     *JarchiveBoard `json:"single,omitempty"`
	Double     *JarchiveBoard `json:"double,omitempty"`
	Final      *JarchiveFinal `json:"final,omitempty"`
	TieBreaker *JarchiveFinal `json:"tiebreaker,omitempty"`
}

func (episode *episode) String() string {
	return fmt.Sprintf("Match %d (episode #%d)",
		episode.MatchID.Match,
		episode.EpisodeID)
}

func (episode *episode) MatchNumber() schema.MatchNumber {
	return episode.MatchID.Match
}

func (episode *episode) URL() string {
	const FULL_EPISODE_FMT = "https://j-archive.com/showgame.php?game_id=%d"
	return fmt.Sprintf(FULL_EPISODE_FMT, episode.EpisodeID)
}

func (episode *episode) FilepathHTML() string {
	return path.Join("jarchive", "episode",
		fmt.Sprintf("%d.html", episode.EpisodeID))
}

func (episode *episode) FilepathJSON() string {
	return path.Join("json", "episode",
		fmt.Sprintf("%d.html", episode.MatchID.Match))
}

func (episode *episode) ParseHTML(html_bytes []byte) error {
	// TODO
	// TODO
	// TODO populate this instance with parsed contents
	return nil
}

func (episode *episode) WriteJSON(output io.WriteCloser) error {
	defer output.Close()
	// TODO json.Unmarshal
	return nil
}
func (episode *episode) LoadJSON(input io.ReadCloser) error {
	defer input.Close()
	// TODO json.Marshal
	return nil
}
