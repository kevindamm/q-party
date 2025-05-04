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
// github:kevindamm/q-party/schema/round.cue

package schema

//
// ROUNDS = BOARDS | FINAL
//

#RoundID: {
  episode!: #MatchNumber
  round!: int & >=0 & <len(_round_names)
  round_name?: _round_names[round]
  ...
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

// Board representation includes the minimum needed information for starting play.
#Board: #RoundID & {
  columns: [...#CategoryMetadata]
  missing?: [...#BoardPosition]
  ...
}

// Board state includes the player selection and a minified representation of
// which positions are still available, organized by category as a bitmap of
// those (0) missing and (1) present, with least significant bit at the top.
#BoardState: #Board & {
  cat_bitmap: [...uint]
  history: [...#SelectionOutcome]
}

// A board position is identified by its column and (row) index.  These value
// ranges may not need to be hard-coded, though the non-zero positive part will.
#BoardPosition: {
  column!: int & >0 // typically 1..6
  index!: int & >0 // typically 1..5
  ...
}

// Represents the board position and challenge, without contestant performance.
#BoardSelection: #ContestantID & #ChallengeMetadata & #BoardPosition

#SelectionOutcome: #BoardSelection & {
  correct: bool
  delta: #Value
}
