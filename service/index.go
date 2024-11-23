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
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Contains the landing page route and JSON representations of the metadata,
// including seasons, episodes and categories.  These are initially pulled from
// a bespoke flat file instance and can evolve to pull from a database with new
// (custom and updated) challenges and episodes.

// Serves the root page, redirecting to an in-progress game or previous room,
// if a session is present.
func (server *Server) RouteLandingPage() func(ctx echo.Context) error {
	if server.staticFS == nil {
		log.Fatal("embedded filesystem absent when setting up route for landing page")
	}

	// Load the bytes for the favicon during server startup.
	reader, err := server.staticFS.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	homepage, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	return func(ctx echo.Context) error {
		// TODO check if user has logged in?
		// TODO check session for recently being in a room?
		// redirects in either case, otherwise shows the plain index.html

		return ctx.Blob(http.StatusOK, "text/html", homepage)
	}
}

func (server *Server) RouteIndexJSON(index_name string) func(echo.Context) error {
	// Assumes the index will not change, preload the bytes to deliver.
	var bytes []byte
	var err error
	switch index_name {
	case "seasons":
		bytes, err = json.Marshal(server.jarchive.Seasons)
		if err != nil {
			log.Fatal(err)
		}
	case "episodes":
		bytes, err = json.Marshal(server.jarchive.Episodes)
		if err != nil {
			log.Fatal(err)
		}
	case "categories":
		bytes, err = json.Marshal(server.jarchive.Categories)
		if err != nil {
			log.Fatal(err)
		}
	}
	// TODO cache updates to the index here or retrieve from a database on demand.

	return func(ctx echo.Context) error {
		return ctx.Blob(http.StatusOK, "application/json", bytes)
	}
}

func (server *Server) AboutPage() func(echo.Context) error {
	if server.staticFS == nil {
		log.Fatal("embedded filesystem absent when setting up route for landing page")
	}

	// Load the bytes for the favicon during server startup.
	reader, err := server.staticFS.Open("about.html")
	if err != nil {
		log.Fatal(err)
	}
	about_html, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	return func(ctx echo.Context) error {
		return ctx.Blob(http.StatusOK, "text/html", about_html)
	}
}
