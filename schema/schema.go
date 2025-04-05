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
// github:kevindamm/q-party/schema/schema.go

package schema

/*
 * CHALLENGES
 */

type ChallengeID uint64

const UNKNOWN_CHALLENGE = ChallengeID(0)

type Value int

type Wager struct {
	Value
}

type ChallengeMetadata[T Wager | Value] struct {
	ChallengeID `json:"id"`
	Value       T `json:"value,omitempty"`
}

type Challenge struct {
	ChallengeMetadata[Value]
	Clue string `json:"clue"`

	Media    []MediaClue `json:"media"`
	Category string      `json:"category"`
	Comments string      `json:"comments,omitempty"`
}

type MediaClue struct {
	MimeType string `json:"mime"`
	URL      string `json:"url"`
}

type HostChallenge struct {
	Challenge
	Correct []string `json:"correct"`
}

type PlayerWager struct {
	BoardSelection[Wager]
}

type PlayerResponse struct {
	ChallengeMetadata[Value]
	Response string `json:"response,omitempty"`
}

/*
 * ROUNDS
 */

type RoundEnum int

const (
	UNKNOWN_ROUND    RoundEnum = 0
	ROUND_SINGLE     RoundEnum = 1
	ROUND_DOUBLE     RoundEnum = 2
	ROUND_FINAL      RoundEnum = 3
	ROUND_TIEBREAKER RoundEnum = 4
	ROUND_OFFLINE    RoundEnum = 5
)

type Board struct {
	EpisodeID ShowNumber              `json:"episode"`
	Round     RoundEnum               `json:"round"`
	Columns   []Category              `json:"columns"`
	Missing   []BoardPosition         `json:"missing,omitempty"`
	History   []BoardSelection[Value] `json:"history,omitempty"`
}

type BoardPosition struct {
	Column uint `json:"column"`
	Index  uint `json:"index"`
}

type BoardSelection[T Value | Wager] struct {
	BoardPosition
	ChallengeMetadata[T]
}

type Category struct {
	Title      string                     `json:"title"`
	Comments   string                     `json:"comments,omitempty"`
	Challenges []ChallengeMetadata[Value] `json:"challenges"`
}

/*
 * EPISODES
 */

type ShowNumber uint64

type ShowDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type EpisodeIndex struct {
	Episodes map[ShowNumber]EpisodeMetadata `json:"episodes"`
}

type EpisodeMetadata struct {
	ShowNumber `json:"show_number"`
	ShowTitle  string   `json:"show_title"`
	Season     SeasonID `json:"season,omitempty"`
	AiredDate  ShowDate `json:"aired,omitempty"`
	TapedDate  ShowDate `json:"taped,omitempty"`

	Contestants []ContestantID `json:"contestant_ids,omitempty"`
	Comments    string         `json:"comments,omitempty"`
	Media       []MediaClue    `json:"media"`
}

type EpisodeStats struct {
	SingleCount    int `json:"single_count,omitempty"`
	DoubleCount    int `json:"double_count,omitempty"`
	TripleStumpers int `json:"triple_stumpers,omitempty"`
}

/*
 * CONTESTANTS
 */

type ContestantID struct {
	PK   uint64 `json:"id"`
	Name string `json:"name,omitempty"`
}

type Contestant struct {
	ContestantID
	Occupation string `json:"occupation"`
	Residence  string `json:"residence"`
	Notes      string `json:"notes"`
}

type Appearance struct {
	ContestantID
	Episodes ShowNumber
}

type Career struct {
	ContestantID
	Episodes []ShowNumber `json:"episodes"`
	Winnings Value        `json:"winnings"`
}

/*
 * SEASONS
 */

type SeasonID string

type SeasonIndex struct {
	SemVer  []uint                      `json:"version"`
	Seasons map[SeasonID]SeasonMetadata `json:"seasons"`
}

type SeasonMetadata struct {
	Season SeasonID `json:"season"`
	Name   string   `json:"name"`
	Aired  struct {
		From  ShowDate `json:"from"`
		Until ShowDate `json:"until"`
	}

	EpisodeCount   int `json:"episode_count,omitempty"`
	ChallengeCount int `json:"challenge_count,omitempty"`
	TripStumpCount int `json:"tripstump_count,omitempty"`
}

/*
 * EMBEDDED SCHEMA
 */

// go:embed challenges.cue
var schemaChallenges string

// go:embed rounds.cue
var schemaRounds string

// go:embed episodes.cue
var schemaEpisodes string

// go:embed seasons.cue
var schemaSeasons string

// go:embed contestants.cue
var schemaContestants string
