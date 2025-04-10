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
// github:kevindamm/q-party/schema/challenge.cue

package schema

#Value: int
#Wager: int & >0

// A unique identifier for the challenge and (optionally) its ID.
// If value is undefined it did not have an associated monetary value.
#ChallengeMetadata: {
  qid!: uint & >=0
  value?: #Value
  ...
}

// Sentinel representation for a blank board cell.
UnknownChallenge: #ChallengeMetadata & { qid: 0 }

// The challenge details, except the correct answer(s).
#ChallengeData: {
  clue!: string // markdown format, including media references

  media?: [...#MediaRef]
  category?: string
  comments?: string
  ...
}

#Challenge: #ChallengeMetadata & #ChallengeData & {
  value: #Value
}

#BiddingChallenge: #ChallengeMetadata & #ChallengeData & {
  wager: #Wager
}

// A link to the media accompaniment for a challenge
// (or for the commentary of an episode).
#MediaRef: {
  mime: string & #MimeType
  url: string
}

// The allowed mime types for media assets (plus text/plain;encoding=UTF-8).
#MimeType: ( "image/jpeg" |
              "image/png" |
             "image/jpeg" |
          "image/svg+xml" |
             "audio/mpeg" |
              "video/mp4" |
        "video/quicktime" )

// The host may see the correct answer while the contestants cannot.
#HostChallenge: #Challenge & {
  correct: [...string] // excluding "what is..." preface
  value?: #Value
  wager?: #Wager
  ...
}

// Before answering, sometimes a player must provide a wager value first.
#PlayerWager: #ContestantID & #ChallengeMetadata & {
  wager!: #Wager
  comments?: string
}

// The player's response for a challenge, if entered as plain text.
// This may instead be an audio file stored as multi-part form attachment.
#PlayerResponse: #ContestantID & #ChallengeMetadata & {
  response?: string
}
