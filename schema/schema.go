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
	Value `json:"value"`
}

type ChallengeMetadata[T Wager | Value] struct {
	ChallengeID `json:"id"`
	Value       T `json:"value,omitempty"`
}

type Challenge struct {
	ChallengeMetadata[Value] `json:",inline"`
	Clue                     string `json:"clue"`

	Media    []MediaClue `json:"media"`
	Category string      `json:"category"`
	Comments string      `json:"comments,omitempty"`
}

type MediaClue struct {
	MimeType string `json:"mime"`
	URL      string `json:"url"`
}

type HostChallenge struct {
	Challenge `json:",inline"`
	Correct   []string `json:"correct"`
}

type PlayerWager struct {
	BoardSelection[Wager] `json:",inline"`
}

type PlayerResponse struct {
	ChallengeMetadata[Value] `json:",inline"`
	Contestant               ContestantID `json:"contestant"`
	Response                 string       `json:"response,omitempty"`
}

/*
 * ROUNDS
 */

type RoundID struct {
	Episode MatchNumber `json:"episode,omitempty"`
	Round   RoundEnum   `json:"round,omitempty"`
}

type Board struct {
	RoundID `json:",inline"`
	Columns []Category      `json:"columns"`
	Missing []BoardPosition `json:"missing,omitempty"`
}

type BoardState struct {
	Board   `json:",inline"`
	History []BoardSelection[Value] `json:"history"`
}

// An enum-like value for the different rounds.
type RoundEnum int

const (
	ROUND_UNKNOWN RoundEnum = iota
	ROUND_SINGLE
	ROUND_DOUBLE
	ROUND_FINAL
	ROUND_TIEBREAKER
	PRINTED_MEDIA
	MaxRoundEnum
)

func (round RoundEnum) String() string {
	if round >= MaxRoundEnum {
		return round_names[0]
	}
	return round_names[round]
}

func (round RoundID) RoundName() string {
	return round_names[round.Round]
}

var round_names = [6]string{
	"[UNKNOWN]",
	"Single!",
	"Double!",
	"Final!",
	"Tiebreaker!!",
	"[printed media]"}

type BoardPosition struct {
	Column uint `json:"column"`
	Index  uint `json:"index"`
}

type BoardSelection[T Value | Wager] struct {
	BoardPosition        `json:",inline"`
	ChallengeMetadata[T] `json:",inline"`
}

type CategoryID uint64

type Category struct {
	CategoryID   `json:"id"`
	Title        string        `json:"title"`
	Comments     string        `json:"comments,omitempty"`
	ChallengeIDs []ChallengeID `json:"challenges,omitempty"`
}

/*
 * EPISODES
 */

// Shows are numbered sequentially based on air date.
// These are historic contests obtained piecemeal from jarchive.com
type ShowIndex struct {
	Season    SeasonID    `json:"season,omitempty"`
	Episode   MatchNumber `json:"episode"`
	ShowTitle string      `json:"show_title,omitempty"`
}

// A match identifier refers to the unique identifier of the ?-Party database.
// These are not universally unique, they are only certain to be locally unique.
type MatchNumber uint64

type ShowDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type EpisodeIndex struct {
	Episodes map[MatchNumber]EpisodeMetadata `json:"episodes"`
}

type EpisodeMetadata struct {
	ShowIndex `json:"show"`
	Season    SeasonID `json:"season,omitempty"`
	AiredDate ShowDate `json:"aired,omitempty"`
	TapedDate ShowDate `json:"taped,omitempty"`

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
	ContestantID `json:",inline"`
	Occupation   string `json:"occupation"`
	Residence    string `json:"residence"`
	Notes        string `json:"notes"`
}

type Appearance struct {
	ContestantID `json:",inline"`
	Episode      ShowIndex
}

type Career struct {
	ContestantID `json:",inline"`
	Appearances  []ShowIndex `json:"appearances"`
	Winnings     Value       `json:"winnings"`
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
