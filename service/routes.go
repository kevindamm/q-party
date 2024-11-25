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
	"io"
	"io/fs"
	"log"
	"net/http"

	qparty "github.com/kevindamm/q-party"
	"github.com/labstack/echo/v4"
)

func (server *Server) RegisterRoutes() *http.Handler {
	//	e.POST("/join", server.RouteJoinRoom())
	//
	//	e.GET("/play/:roomid", server.RoutePlayRoom())
	//
	//	// Generate a random board of six categories and send response as JSON.
	//	e.GET("/play/:roomid/random/episode", server.RouteRandomEpisode())
	//	e.GET("/play/:roomid/random/categories", server.RouteRandomCategories())
	//	e.GET("/play/:roomid/random/challenges", server.RouteRandomChallenges())
	//
	//	spa := e.Group("/play")
	//	spa.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	//		HTML5:      true,
	//		Root:       "play",
	//		IgnoreBase: false,
	//		Filesystem: http.Dir("app/dist"),
	//	}))
	//
	//	// TODO websockets endpoints
	//	// /join/:room_id (ask to join room, get host/contestant/audience assignment)
	//	// /play/:room_id (contestant interface) .Group( player_auth )
	//	// /host/:room_id (host interface)       .Group( host_auth )
	//
	//	// Vue3 app is hosted here, see /app/* within this repo for implementation.
	//	// It's effectively static until it routes to lobby or an in-progress game.
	//	e.Static("/app", "app/dist/").Name = "vue3-root"
	//	// image, audio and video media for challenges are also under a static path.
	//	e.Static("/media", "media").Name = "media-root"
	//
	//	// TODO other request handlers
	//	// TODO /view/:room_id (audience interface via SSE)
	//
	//	return e
	return nil
}

func (server *Server) RouteStaticFiles(staticFS fs.FS) {
	if staticFS == nil {
		log.Fatal("embedded filesystem absent when setting up routes")
	}
	// Serve files embedded in staticFS for the root path.
	server.echo.GET("/",
		SingleFileHandler(staticFS, "static/index.html", "text/html"))
	server.echo.GET("/about",
		SingleFileHandler(staticFS, "static/about.html", "text/html"))
	server.echo.GET("/*",
		echo.WrapHandler(EmbedFSHandler(staticFS)))

	// Retrieve the season, episode and category indices as subsets of the index.
	// Served separately because they are often used independently and can be
	// delivered in smaller pieces this way.
	server.echo.GET("/seasons", server.RouteIndexJSON("seasons"))
	server.echo.GET("/episodes", server.RouteIndexJSON("episodes"))
	server.echo.GET("/categories", server.RouteIndexJSON("categories"))
}

func SingleFileHandler(staticFS fs.FS, path, filetype string) echo.HandlerFunc {
	if staticFS == nil {
		log.Fatal("empty filesystem")
	}
	file, err := staticFS.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return func(ctx echo.Context) error {
		return ctx.Blob(http.StatusOK, filetype, bytes)
	}
}

func EmbedFSHandler(staticFS fs.FS) http.Handler {
	fsys, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(fsys))
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
