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
// github:kevindamm/q-party/service/index.go

package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	qparty "github.com/kevindamm/q-party"
	"github.com/labstack/echo/v4"
)

type JArchiveIndex struct {
	Version    []uint                                    `json:"version,omitempty"`
	Seasons    map[qparty.SeasonID]qparty.SeasonMetadata `json:"seasons"`
	Categories map[string]qparty.CategoryMetadata        `json:"categories"`
	Episodes   map[qparty.EpisodeID]qparty.EpisodeStats  `json:"episodes"`
}

func (server *Server) LoadJArchiveIndex(jarchive_json []byte) error {
	index := new(JArchiveIndex)
	err := json.Unmarshal(jarchive_json, index)
	if err != nil {
		return err
	}
	// TODO create additional (in-memory) indexes for easier retrieval
	// TODO also cache recently loaded episodes?  that can be done separately in main, though
	server.jarchive = index
	return nil
}

func (all_seasons JArchiveIndex) WriteSeasonIndexJSON(json_path string) error {
	writer, err := os.Create(json_path)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(all_seasons)
	if err != nil {
		return fmt.Errorf("failed to marshal seasons to JSON bytes\n%s", err)
	}
	nbytes, err := writer.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write%s\n%s", json_path, err)
	} else {
		log.Printf("Wrote seasons.json, %d bytes", nbytes)
	}

	return nil
}

func (index JArchiveIndex) EpisodesBySeason(jsid qparty.SeasonID) []qparty.EpisodeMetadata {
	episodes := make([]qparty.EpisodeMetadata, 0)
	// TODO

	return episodes
}

func (index JArchiveIndex) GetShowEpisode(show qparty.ShowNumber) qparty.EpisodeMetadata {
	// TODO
	return qparty.EpisodeMetadata{}
}

// Handler for static JSON files representing parts of the index.
func (server *Server) RouteIndexJSON(index_name string) func(echo.Context) error {
	// Assumes the index will not change, preload the bytes to deliver.
	var bytes []byte
	var err error
	switch index_name {
	case "seasons":
		bytes, err = json.Marshal(server.jarchive.Seasons)
		if err != nil {
			log.Print("ERROR marshaling seasons into JSON" + err.Error())
		}
	case "episodes":
		bytes, err = json.Marshal(server.jarchive.Episodes)
		if err != nil {
			log.Print("ERROR marshaling episodes into JSON" + err.Error())
		}
	case "categories":
		bytes, err = json.Marshal(server.jarchive.Categories)
		if err != nil {
			log.Print("ERROR marshaling categories into JSON" + err.Error())
		}
	}

	return func(ctx echo.Context) error {
		return ctx.Blob(http.StatusOK, "application/json", bytes)
	}
}
