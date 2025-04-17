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
	"path"

	"github.com/kevindamm/q-party/schema"
	"github.com/kevindamm/q-party/selfhost/cmd/fetch"
)

type EpisodeID int

type EpisodeMatchNumber map[EpisodeID]schema.MatchNumber
type MatchNumberEpisode map[schema.MatchNumber]EpisodeID

// All details of the episode, including correct answers & the contestants' bios.
type JarchiveEpisode interface {
	fetch.Fetchable
}

func NewEpisode(id EpisodeID) JarchiveEpisode {
	episode := episode{
		EpisodeID: id,
	}

	// TODO
	return &episode
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

func ParseEpisodeHtml(episode_html []byte) (JarchiveEpisode, error) {
	episode := episode{}
	// TODO

	return &episode, nil
}

func (episode *episode) String() string {
	// TODO
	return "episode ..."
}

func (episode *episode) URL() string {
	const FULL_EPISODE_FMT = "https://j-archive.com/showgame.php?game_id=%d"
	return fmt.Sprintf(FULL_EPISODE_FMT, episode.EpisodeID)
}

func (episode *episode) FilepathHTML() string {
	return path.Join("episode", fmt.Sprintf("%d.html", episode.EpisodeID))
}

func (episode *episode) FilepathJSON() string {
	return path.Join("json", "episode", fmt.Sprintf("%d.html", episode.MatchID))
}

func (episode *episode) ParseHTML(html_bytes []byte) error {

	// TODO
	return nil
}

func (episode *episode) WriteJSON(output io.WriteCloser) error {
	defer output.Close()
	// TODO
	return nil
}
func (episode *episode) LoadJSON(input io.ReadCloser) error {
	defer input.Close()
	// TODO
	return nil
}
