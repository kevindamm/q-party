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
// github:kevindamm/q-party/cue/challenges.cue

package schema

#Value: int

#Wager: {
  wager: #Value
}

// A unique identifier for the challenge and (optionally) its ID.
// If value is undefined it did not have an associated monetary value.
#ChallengeMetadata: {
  qid!: uint & >=0
  value?: #Value | #Wager
  ...
}

// Sentinel representation for a blank board cell.
UnknownChallenge: #ChallengeMetadata & { qid: 0 }

// The challenge details, except the correct answer(s).
#Challenge: #ChallengeMetadata & {
  value!: #Value
  clue!: string

  media?: [...#MediaClue]
  category?: string
  comments?: string
}

// A link to the media accompaniment for a challenge
// (or for the commentary of an episode).
#MediaClue: {
  mime: string & #MimeType
  url: string
}

// The allowed mime types for media assets (plus text/plain;encoding=UTF-8).
#MimeType: "image/jpeg" | "audio/mpeg" | "video/mp4"

// The host may see the correct answer while the contestants cannot.
#HostChallenge: #Challenge & {
  correct: [...string] // excluding "what is..." preface
}

// Before answering, sometimes a player must provide a wager value first.
#PlayerWager: #ContestantID & #ChallengeMetadata & {
  value!: #Wager
  comments?: string
}

// The player's response for a challenge.
#PlayerResponse: #ContestantID & #ChallengeMetadata & {
  response?: string
}
