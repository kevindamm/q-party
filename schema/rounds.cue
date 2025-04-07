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
// github:kevindamm/q-party/cue/rounds.cue

package schema

#RoundID: {
  episode!: #MatchNumber
  round!: int & >=0 & <len(_round_names)
  round_name?: _round_names[round]
  ...
}

// Board representation includes the minimum needed information for starting play.
#Board: #RoundID & {
  columns: [...#Category]
  missing?: [...#BoardPosition]
  ...
}

// Board state includes the player selection
#BoardState: #Board & {
  history: [...#BoardSelection]
}

// A board position is identified by its column and (row) index.
#BoardPosition: {
  column!: uint & <6
  index!: uint & <5
}

// Represents the board position and challenge, without contestant performance.
#BoardSelection: #ContestantID & #ChallengeMetadata & #BoardPosition

#CategoryMetadata: {
  catID: string
  title!: string
  ...
}

// A category instance must have a title and
// may have any number of challenges (typically five).
#Category: #CategoryMetadata & {
  comments?: string
  challenges: [...#ChallengeMetadata]
}

#CategoryAired: #CategoryMetadata & {
  aired: #ShowDate
}

// Display strings for the different rounds.
_round_names: [...string] & [
  "[UNKNOWN]",
	"Single!",
	"Double!",
	"Final!",
	"Tiebreaker!!",
	"[other]",
]

#CategoryTheme: string

_cat_themes: [...#CategoryTheme] & [
	"",
	"Geography",
	"Entertainment",
	"History & Royalty",
	"Art & Literature",
	"Science & Nature",
	"Sports & Leisure",
]
