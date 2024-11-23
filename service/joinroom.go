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
	"log"

	"github.com/labstack/echo/v4"
)

type RoomMetadata struct {
	Members []QPartySecret
}

type QPartyPlayer struct {
	Name string `json:"name"`
}

type QPartySecret struct {
	QPartyPlayer
	Token string
}

func (server *Server) RouteJoinRoom() func(echo.Context) error {
	// Setup fixed set of available rooms.
	// (a future version may allow this to be dynamic via a database)
	var rooms map[string]*RoomMetadata
	if server.jsonFS == nil {
		log.Fatal("missing embedded json files")
	}
	reader, err := server.jsonFS.Open("rooms.json")
	if err != nil {
		log.Fatal(err)
	}
	json_bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(json_bytes, &rooms)

	return func(ctx echo.Context) error {
		room_id := ctx.Param("room_id")
		room, ok := rooms[room_id]
		if !ok {
			return fmt.Errorf("room not found")
		}
		if len(room.Members) == 0 {
			// First to join the room becomes host.

		}

		return nil
	}
}

func NewRoomMetadata() *RoomMetadata {
	metadata := new(RoomMetadata)
	metadata.Members = make([]QPartySecret, 0)
	return metadata
}
