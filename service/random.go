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
// github:kevindamm/q-party/service/random.go

package service

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"

	qparty "github.com/kevindamm/q-party"
	"github.com/labstack/echo/v4"
)

func (server *Server) RouteRandomEpisode() func(echo.Context) error {
	// TODO

	return func(ctx echo.Context) error {
		// TODO
		return nil
	}
}

func (server *Server) RouteRandomCategories() func(echo.Context) error {
	// TODO

	return func(ctx echo.Context) error {
		// TODO

		return nil
	}
}

func (server *Server) RouteRandomChallenges() func(echo.Context) error {
	// TODO

	return func(ctx echo.Context) error {
		// TODO

		return nil
	}
}

func (server *Server) NewRandomBoard() func(ctx echo.Context) error {
	if server.jsonFS == nil {
		log.Fatal("embedded json files absent when setting up route for random categories")
	}

	cat_names := make([]string, 0, len(server.jarchive.Categories))
	i := 0
	for key := range server.jarchive.Categories {
		cat_names[i] = key
		i++
	}

	return func(ctx echo.Context) error {
		indices := make(map[int]any)
		categories := make([]qparty.CategoryMetadata, 0)
		for range 6 {
			index := rand.Int() % len(cat_names)
			if _, ok := indices[index]; ok {
				continue
			}
			indices[index] = struct{}{}
			categories = append(categories,
				server.jarchive.Categories[cat_names[index]])

		}

		// TODO export qparty.Board instead
		json_bytes, err := json.Marshal(categories)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, json_bytes)
	}
}
