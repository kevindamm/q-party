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
// github:kevindamm/q-party/challenges.go

package qparty

type ChallengeMetadata struct {
	ChallengeID uint        `json:"id"`
	Value       DollarValue `json:"value,omitempty"`

	TripleStumper bool `json:"stumper,omitempty"`
}

// Sentinel value for board entries that are missing/blank.
var UnknownChallenge = ChallengeMetadata{0, 0, false}

// Challenge data (without the answer), for when a board position is selected.
type Challenge struct {
	ChallengeMetadata `json:",inline"`
	Clue              string `json:"clue"`

	Media    []Media `json:"media,omitempty"`
	Category string  `json:"category,omitempty"`
	Comments string  `json:"comments,omitempty"`
}

// Host view of the challenge, includes the correct response.
type HostChallenge struct {
	Challenge `json:",inline"`
	Correct   string `json:"correct"` // excluding "what is..." preface
}

type PlayerWager struct {
	ChallengeMetadata `json:",inline"`
	Comments          string `json:"comments,omitempty"`
}

type PlayerResponse struct {
	ChallengeMetadata `json:",inline"`
	Response          string `json:"response,omitempty"`
}
