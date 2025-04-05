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
// github:kevindamm/q-party/cue/seasons.cue

package schema

// Schema for [seasons.json] containing season and episode metadata.
#SeasonIndex: {
  version?: [...uint]
  seasons: [#SeasonID]: #SeasonMetadata
}

// Unique identifier for the season and its episodes.
#SeasonID: string

// Metadata for a single season, has identity and some statistics.
#SeasonMetadata: {
  season: #SeasonID
  name: string
  aired: {
    from: #ShowDate
    until: #ShowDate
  }

  episode_count?: int
  challenge_count?: int
  tripstump_count?: int
}
