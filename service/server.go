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
// github:kevindamm/q-party/service/server.go

package service

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	qparty "github.com/kevindamm/q-party"
	"github.com/labstack/echo/v4"
)

type Server struct {
	*http.Server

	port     int
	jarchive *qparty.JArchiveIndex
	jsonFS   fs.FS
	staticFS fs.FS
}

func NewServer(jarchive *qparty.JArchiveIndex, jsonFS fs.FS, staticFS fs.FS) *Server {
	port, err := strconv.Atoi(os.Getenv("QPARTY_PORT"))
	if err != nil {
		port = 80
	}
	qps := new(Server)
	qps.port = port
	log.Printf("Listening on port %d", port)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", qps.port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	qps.Server = &server
	qps.jarchive = jarchive
	qps.jsonFS = jsonFS
	qps.staticFS = staticFS
	qps.Handler = qps.RegisterRoutes()

	return qps
}

func (server *Server) Favicon() func(echo.Context) error {
	if server.staticFS == nil {
		log.Fatal("embedded filesystem absent when setting up favicon route")
	}
	favicon_bytes := must_load_file(server.staticFS, "favicon.ico")

	return func(ctx echo.Context) error {
		return ctx.Blob(http.StatusOK, "image/x-icon", favicon_bytes)
	}
}

func (server *Server) StyleCSS() func(echo.Context) error {
	if server.staticFS == nil {
		log.Fatal("embedded filesystem absent when setting up favicon route")
	}
	stylecss_bytes := must_load_file(server.staticFS, "style.css")

	return func(ctx echo.Context) error {
		return ctx.Blob(http.StatusOK, "text/css", stylecss_bytes)
	}
}

func must_load_file(fs fs.FS, filename string) []byte {
	reader, err := fs.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
