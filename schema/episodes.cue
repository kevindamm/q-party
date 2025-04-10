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
// github:kevindamm/q-party/schema/episodes.cue

package schema

// Unique identifier for an episode.
#MatchID: {
  season?: #SeasonName
  match: #MatchNumber
  show_title?: string
}

#MatchNumber: uint64 & >0

#EpisodeIndex: [#MatchNumber]: #EpisodeMetadata

// Identifiers and statistics for each episode.
#EpisodeMetadata: #MatchID & {
  jaid?: uint

  aired?: #ShowDate
  taped?: #ShowDate

  contestants?: [...#ContestantID]
  media?: [...#MediaRef]
  comments?: string
  ...
}

// A minified 
#BoardLayout: {
  cat_bitmap: [...uint]
}

#EpisodeStats: #EpisodeMetadata & {
  single_count?: int
  double_count?: int
  triple_stumpers?: [...#BoardPosition]
}

// Represents a (year, month, day) when a show was aired or taped.
#ShowDate: {
  year: int & >1980
  month: int & >=1 & <=12
  day: int & >=1 & <=31
}

#ShowDateRange: {
  from?: #ShowDate
  until?: #ShowDate
}
