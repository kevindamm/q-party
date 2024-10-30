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
// github:kevindamm/q-party/cmd/jarchive/episode_round.go

package main

// enum representation for
type EpisodeRound int

const (
	ROUND_UNKNOWN EpisodeRound = iota
	ROUND_SINGLE_JEOPARDY
	ROUND_DOUBLE_JEOPARDY
	ROUND_FINAL_JEOPARDY
	ROUND_TIE_BREAKER
	ROUND_PRINTED_MEDIA
)

var round_strings = map[EpisodeRound]string{
	ROUND_UNKNOWN:         "[UNKNOWN]",
	ROUND_SINGLE_JEOPARDY: "Jeopardy!",
	ROUND_DOUBLE_JEOPARDY: "Double Jeopardy!",
	ROUND_FINAL_JEOPARDY:  "Final Jeopardy!",
	ROUND_TIE_BREAKER:     "Tiebreaker",
	ROUND_PRINTED_MEDIA:   "[printed media]",
}

func (round EpisodeRound) String() string {
	printed := round_strings[round]
	if printed == "" {
		printed = round_strings[ROUND_UNKNOWN]
	}
	return printed
}

func ParseString(round string) EpisodeRound {
	for k, v := range round_strings {
		if v == round {
			return k
		}
	}
	return ROUND_UNKNOWN
}
