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
// github:kevindamm/q-party/service/routes_test.go

package service

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"testing/fstest"

	"github.com/labstack/echo/v4"
)

func TestLanding(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	ctx := e.NewContext(request, response)
	server := new(Server)

	test_fs := fstest.MapFS{
		"index.html": {Data: []byte("<html><body><p>Hello, World!</p></body></html>")},
	}

	handler := server.LandingPage(test_fs)
	if err := handler(ctx); err != nil {
		t.Errorf("handler() error = %v", err)
		return
	}

	if response.Code != http.StatusOK {
		t.Errorf("handler() unexpected status code = %v", response.Code)
		return
	}

	reContentMatch := regexp.MustCompile(`Hello, World!`)
	if !reContentMatch.Match(response.Body.Bytes()) {
		t.Error("handler() unexpected body (missing the greeting)\n",
			response.Body.String())
		return
	}
}
