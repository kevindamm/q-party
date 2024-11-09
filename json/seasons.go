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
// github:kevindamm/q-party/json/seasons.go

package json

type SeasonIndex struct {
	Version  []int             `json:"version,omitempty"`
	Seasons  []SeasonMetadata  `json:"seasons"`
	Episodes []EpisodeMetadata `json:"episodes"`
}

type SeasonID struct {
	JSID `json:"id"`
	Name string `json:"name,omitempty"`
}

type JSID string

type SeasonMetadata struct {
	SeasonID `json:",inline"`
	Aired    ShowDateRange `json:"aired"`

	EpisodeCount   int `json:"episode_count"`
	ChallengeCount int `json:"challenge_count"`
}

type ShowDateRange struct {
	From  ShowDate `json:"from"`
	Until ShowDate `json:"until"`
}
