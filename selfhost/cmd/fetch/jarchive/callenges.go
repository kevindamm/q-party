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
// github:kevindamm/q-party/selfhost/cmd/fetch/jarchive/challenges.go

package main

import "github.com/kevindamm/q-party/schema"

// Full representation of the challenge, including the correct response.
type JarchiveChallenge struct {
	schema.Challenge
	Responses []string
	Correct   []string
	Stumped   int // 0..3 how many incorrect responses were given (not counting silence)
}

// Full representation, as with the above, as well as each player's wager.
type JarchiveFinal struct {
	schema.Challenge
	Wagers    []schema.PlayerWager
	Responses []string
	Correct   []string
	Stumped   int // 0..|players|
}

// Sentinel value for board entries that are missing/blank.
func UnknownChallenge() JarchiveChallenge {
	return JarchiveChallenge{
		schema.UnknownChallenge(),
		[]string{},
		[]string{},
		0}
}

// There may not be a final jeopardy (or it may not have been entered yet).
func UnknownFinal() JarchiveFinal {
	return JarchiveFinal{
		schema.UnknownChallenge(),
		[]schema.PlayerWager{},
		[]string{},
		[]string{},
		0}
}
