// Copyright (c) 2025 Kevin Damm
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
// github:kevindamm/q-party/cmd/fetch/jarchive/final_test.go

package main_test

import (
	"strings"
	"testing"

	jarchive "github.com/kevindamm/q-party/selfhost/cmd/fetch/jarchive"
	"golang.org/x/net/html"
)

func TestParseFinalChallenge(t *testing.T) {
	html_string := `
	`
	html_reader := strings.NewReader(html_string)
	dom, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("??? failed to parse test final challenge\n%s\n\n%s\n",
			err, html_string)
	}
	clue_td := jarchive.NextDescendantWithClass(dom, "td", "clue")

	expected := &jarchive.JarchiveFinalChallenge{}
	//expected.Value = ...

	parsed, err := jarchive.ParseFinalChallenge(clue_td)
	if err != nil {
		t.Error(err)
	}

	if !equalFinal(parsed, expected) {
		t.Error("final challenge values did not match")
	}
}

func TestParseTiebreakerChallenge(t *testing.T) {
	html_string := ` `
	html_reader := strings.NewReader(html_string)
	dom, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("??? failed to parse Tiebreaker test input\n%s\n\n%s\n",
			err, html_string)
	}
	expected := jarchive.JarchiveTiebreaker{}
	//TODO

	parsed, err := jarchive.ParseTiebreakerChallenge(dom)
	if err != nil {
		t.Error(err)
	}
	if !equalTB(parsed, &expected) {
		t.Logf("parsed results mismatch\nexpected:\n%v\n\nparsed:\n%v\n",
			expected, parsed)
		t.FailNow()
	}
}

func equalFinal(have, expect *jarchive.JarchiveFinalChallenge) bool {
	// TODO
	return false
}

func equalTB(have, expect *jarchive.JarchiveTiebreaker) bool {
	// TODO
	return false
}
