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
// github:kevindamm/q-party/service/routes.go

package service

import (
	"io/fs"
	"net/http"

	qparty "github.com/kevindamm/q-party"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (server *Server) RegisterRoutes(fs fs.FS) http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = NewRenderer()

	e.GET("/", server.LandingPage(fs)).Name = "home"

	e.GET("/favicon.ico", Favicon(fs)).Name = "favicon"

	// TEMPORARY just to view some stats
	e.GET("/categories", server.ListCategoriesByYear)

	// websockets endpoints
	// /join/:room_id (ask to join room, get host/contestant/audience assignment)
	// /play/:room_id (contestant interface) .Group( player_auth )
	// /host/:room_id (host interface)       .Group( host_auth )

	// Vue3 app is hosted here, see /app/* within this repo for implementation.
	// It's effectively static until it routes to lobby or an in-progress game.
	e.Static("/app", "app/dist/").Name = "vue3-root"
	// image, audio and video media for challenges are also under a static path.
	e.Static("/media", "media").Name = "media-root"
	// And some root-level static files that can be listed individually.
	e.Static("/jarchive.json", "public/jarchive.json")

	// Generate a random board of six categories and send response as JSON.
	e.GET("/randboard", server.NewRandomBoard(fs))

	// TODO other request handlers
	// /view/:room_id (audience interface via SSE)

	return e
}

func (server *Server) ListCategoriesByYear(ctx echo.Context) error {
	cat_years := make(map[int][]qparty.CategoryMetadata)
	for _, category := range server.Categories {
		for _, episode := range category.Episodes {
			year := episode.Year
			list, exists := cat_years[year]
			if exists {
				cat_years[year] = append(list, category)
			} else {
				cat_years[year] = []qparty.CategoryMetadata{category}
			}
		}
	}

	return ctx.HTML(http.StatusOK, "TODO template categories")
}
