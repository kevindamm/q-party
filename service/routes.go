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
	"log"
	"net/http"

	qparty "github.com/kevindamm/q-party"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (server *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.HTTPSRedirect())
	e.Renderer = NewRenderer()

	if server.jsonFS == nil || server.staticFS == nil {
		log.Fatal("embedded filesystem absent when setting up routes")
	}

	// Retrieve the season, episode and category indices as subsets of the index.
	// Served separately because they are often used independently and can be
	// delivered in smaller pieces this way.
	e.GET("/seasons", server.RouteIndexJSON("seasons"))
	e.GET("/episodes", server.RouteIndexJSON("episodes"))
	e.GET("/categories", server.RouteIndexJSON("categories"))

	e.POST("/join", server.RouteJoinRoom())

	// Serve only files embedded in staticFS for the root path.
	static := e.Group("/")
	static.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(server.staticFS)}))

	e.GET("/play/:roomid", server.RoutePlayRoom())

	// Generate a random board of six categories and send response as JSON.
	e.GET("/play/:roomid/random/episode", server.RouteRandomEpisode())
	e.GET("/play/:roomid/random/categories", server.RouteRandomCategories())
	e.GET("/play/:roomid/random/challenges", server.RouteRandomChallenges())

	spa := e.Group("/play")
	spa.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Root:       "play",
		IgnoreBase: false,
		Filesystem: http.Dir("app/dist"),
	}))

	// TODO websockets endpoints
	// /join/:room_id (ask to join room, get host/contestant/audience assignment)
	// /play/:room_id (contestant interface) .Group( player_auth )
	// /host/:room_id (host interface)       .Group( host_auth )

	// Vue3 app is hosted here, see /app/* within this repo for implementation.
	// It's effectively static until it routes to lobby or an in-progress game.
	e.Static("/app", "app/dist/").Name = "vue3-root"
	// image, audio and video media for challenges are also under a static path.
	e.Static("/media", "media").Name = "media-root"

	// TODO other request handlers
	// TODO /view/:room_id (audience interface via SSE)

	return e
}

func (server *Server) ListCategoriesByYear(ctx echo.Context) error {
	cat_years := make(map[int][]qparty.CategoryMetadata)
	for _, category := range server.jarchive.Categories {
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
