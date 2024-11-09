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
// github:kevindamm/q-party/json/contestants.go

package json

// Struct for embedding in entities that extend the contestant role.
type ContestantID struct {
	UCID `json:"id" cue:">=0"`
	Name string `json:"name,omitempty"`
}

// Unique identifier for contestants.
type UCID uint

type Contestant struct {
	ContestantID `json:",inline"`
	Biography    string `json:"bio"`
}

type Appearance struct {
	ContestantID `json:",inline"`
	Episode      ShowNumber `json:"episode" cue:">0"`
}

type Career struct {
	ContestantID `json:",inline"`
	Episodes     []ShowNumber `json:"episodes"`
	Winnings     int          `json:"winnings"`
}
