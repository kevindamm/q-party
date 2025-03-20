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
// github:kevindamm/q-party/service/updates.go

package service

import (
	"github.com/labstack/echo/v4"
)

// Notifies clients of changes using SSE (server-sent events).  This could also
// be done with websockets, and perhaps that will be necessary for STT later.
// Even with websocket handling, SSE is still better for audience members.
//
// The client app (e.g. a SPA in Vue) can listen to these events and update its
// local path structure to match updates in game state (e.g. view categories,
// peer buzz-in, challenge selection, etc.) that were not initiated by a player.
func (server *Server) RouteInitiateUpdates() func(echo.Context) error {

	return func(ctx echo.Context) error {
		// get room state
		// if not exists, create
		// if empty, add this user session as host
		// else user starts as a player
		//   (or, by request, as audience)

		return nil
	}
}
