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
// github:kevindamm/q-party/schema/contestants.cue

package schema

// Uniquely identifies a contestant across episodes.
#ContestantID: {
  cid!: uint64
  name?: string
  ...
}

// Additional details about the contestant.
#Contestant: #ContestantID & {
  name!: string
  occupation?: string
  residence?: string
  notes?: string
  media?: [...#MediaRef]
}

// An appearance is the joining of a contestant and an episode.
#Appearance: #ContestantID & {
  episode: #MatchID
}

// The episodes that a contestant has appeared in and their total winnings.
#Career: #ContestantID & {
  episodes: [...#MatchID]
  winnings: #Value
}
