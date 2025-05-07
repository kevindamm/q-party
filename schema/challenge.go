// Copyright (c) 2025 Kevin Damm
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
// github:kevindamm/q-party/schema/challenge.go

package schema

import _ "embed"

// go:embed challenge.cue
var schemaChallenge string

type ChallengeID uint64

const UNKNOWN_CHALLENGE_ID = ChallengeID(0)

// Value and Wager are distinct operationally.
type Value int
type Wager int

type ChallengeMetadata struct {
	ChallengeID `json:"qid"`
	Value       `json:"value,omitempty"`
}

type ChallengeData struct {
	Clue string `json:"clue"`

	Media    []MediaRef   `json:"media,omitempty"`
	Category CategoryName `json:"category,omitempty"`
	Comments string       `json:"comments,omitempty"`
}

type Challenge struct {
	ChallengeMetadata `json:",inline"`
	ChallengeData     `json:",inline"`
	Value             Value `json:"value"`
}

type BiddingChallenge struct {
	ChallengeMetadata `json:",inline"`
	ChallengeData     `json:",inline"`
	Wager             Wager `json:"wager"`
}

func UnknownChallenge() Challenge {
	return Challenge{
		ChallengeMetadata: ChallengeMetadata{
			ChallengeID: UNKNOWN_CHALLENGE_ID}}
}

// Challenges may have zero or more media clues (image, audio, video).  Each is
// represented by its own MediaClue instance.  MediaURL is relative a base URL.
type MediaRef struct {
	MimeType `json:"mime,omitempty"`
	MediaURL string `json:"url"`
}

// This enumeration over available media types is modeled after its equivalent
// MIME type such as image/jpeg, image/png, audio/mpeg, etc.  The default (its
// zero value) is an empty string which implicitly represents an unknown media.
type MimeType string

const (
	MediaImageJPG MimeType = "image/jpeg"
	MediaImagePNG MimeType = "image/png"
	MediaImageSVG MimeType = "image/svg+xml"
	MediaAudioMP3 MimeType = "audio/mpeg"
	MediaVideoMP4 MimeType = "video/mp4"
	MediaVideoMOV MimeType = "video/quicktime"
)

type HostChallenge struct {
	Challenge `json:",inline"`
	Value     Value `json:"value,omitempty"`
	Wager     Wager `json:"wager,omitempty"`

	Correct []string `json:"correct"`
}

type PlayerWager struct {
	ContestantID      `json:",inline"`
	ChallengeMetadata `json:",inline"`

	Wager    Wager  `json:"wager"`
	Comments string `json:"comments,omitempty"`
}

type PlayerResponse struct {
	ContestantID      `json:",inline"`
	ChallengeMetadata `json:",inline"`

	Response string `json:"response,omitempty"`
}
