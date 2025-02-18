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
// github:kevindamm/q-party/cue/episodes.cue

package qparty

// Identifiers and statistics for each episode.
#EpisodeMetadata: {
  show_number!: #ShowNumber
  show_title: string
  season?: #SeasonID
  aired?: #ShowDate
  taped?: #ShowDate

  contestant_ids?: [#ContestantID, #ContestantID, #ContestantID]
  comments?: string
  media?: [...#Media]
  ...
}

#EpisodeStats: #EpisodeMetadata & {
  single_count?: int
  double_count?: int
  triple_stumpers?: int
}

// Unique identifier for an episode.
#ShowNumber: uint

// Represents a (year, month, day) when a show was aired or taped.
#ShowDate: {
  year: int & >1980
  month: int & >=1 & <=12
  day: int & >=1 & <=31
}
