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
// github:kevindamm/q-party/json/challenges.go

package json

type ChallengeID struct {
	Ident int         `json:"id"`
	Value DollarValue `json:"value,omitempty"`
}

// Sentinel value for board entries that are missing/blank.
var UnknownChallenge = ChallengeID{0, 0}

type Challenge struct {
	ChallengeID `json:",inline"`
	CluePrompt  string  `json:"clue"`
	Media       []Media `json:"media,omitempty"`

	Category string `json:"category,omitempty"`
	Comments string `json:"comments,omitempty"`
}

type HostChallenge struct {
	ChallengeID `json:",inline"`
	Correct     string `json:"correct"`
}

type PlayerWager struct {
	ChallengeID `json:",inline"`
	Comments    string `json:"comments,omitempty"`
}

type PlayerResponse struct {
	ChallengeID `json:",inline"`
	Response    string `json:"response,omitempty"`
}
