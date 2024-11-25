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
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*http.Server
	Rooms map[RoomID]*RoomState

	port     int
	jarchive *JArchiveIndex

	echo *echo.Echo
	lock sync.Mutex
}

func NewServer(port int) *Server {
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

	qps.echo = echo.New()
	qps.echo.Use(middleware.Logger())
	qps.echo.Use(middleware.Recover())
	//qps.echo.Pre(middleware.HTTPSRedirect())
	//qps.echo.Pre(middleware.RemoveTrailingSlash())

	qps.echo.Renderer = NewRenderer()
	qps.Handler = qps.echo

	return qps
}

func (server *Server) Serve(port int) error {
	url := fmt.Sprintf(":%d", port)
	return server.echo.Start(url)
}

func (server *Server) ServeTLS(crt_path, key_path string) error {
	url := "q-party.kevindamm.com"
	return server.echo.StartTLS(url, crt_path, key_path)
}
