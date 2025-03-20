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
// github:kevindamm/q-party/service/joinroom.go

package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/r3labs/sse"
)

type RoomID string

// Maintains independent state for each room being managed.
type RoomMetadata struct {
	Rooms map[RoomID]RoomConfig
}

type RoomConfig struct {
	Notes string `json:"notes,omitempty"`
}

type RoomState struct {
	Tokens     map[QPartySecret]bool
	Stream     *sse.Stream
	HostPlayer string

	lock sync.Mutex
}

type QPartyPlayer struct {
	Name string `json:"name"`
}

type QPartySecret struct {
	QPartyPlayer
	Token string
}

func (server *Server) RouteJoinRoom(jsonFS fs.FS) func(echo.Context) error {
	// Setup fixed set of available rooms.
	// (a future version may allow this to be dynamic via a database)
	var room_config RoomMetadata
	reader, err := jsonFS.Open("rooms.json")
	if err != nil {
		log.Fatal(err)
	}
	json_bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(json_bytes, &room_config)
	if err != nil {
		log.Fatal(err)
	}

	server.lock.Lock()
	server.Rooms = make(map[RoomID]*RoomState)
	for room_id, config := range room_config.Rooms {
		server.Rooms[room_id] = NewRoomState(config)
	}
	server.lock.Unlock()

	return func(ctx echo.Context) error {
		room_id := RoomID(ctx.Param("room_id"))
		room, ok := server.Rooms[room_id]
		if !ok {
			return fmt.Errorf("room not found")
		}
		token := QPartySecret{QPartyPlayer{Name: "youknowwhom"}, "abcd98765432"}
		if len(room.Tokens) == 0 {
			room.lock.Lock()
			// First to join the room becomes host.
			room.Tokens[token] = true
			room.lock.Unlock()
		}

		return nil
	}
}

func (server *Server) RoutePlayRoom() func(echo.Context) error {
	// TODO

	return func(ctx echo.Context) error {
		return nil
	}
}

func NewRoomState(config RoomConfig) *RoomState {
	state := new(RoomState)
	state.Tokens = make(map[QPartySecret]bool)
	sseServer := sse.New()
	sseServer.AutoReplay = true
	state.Stream = sseServer.CreateStream("updates")

	return state
}
