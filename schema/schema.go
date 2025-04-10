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

import "fmt"

/*
 * CHALLENGES
 */

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

	Media    []MediaRef `json:"media,omitempty"`
	Category string     `json:"category,omitempty"`
	Comments string     `json:"comments,omitempty"`
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

type MediaRef struct {
	MimeType string `json:"mime"`
	URL      string `json:"url"`
}

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

/*
 * ROUNDS
 */

type RoundID struct {
	Episode MatchNumber `json:"episode,omitempty"`
	Round   RoundEnum   `json:"round,omitempty"`
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

func (round RoundID) String() string {
	round_name := round_names[0]
	if round.Round < MaxRoundEnum {
		round_name = round_names[round.Round]
	}
	return fmt.Sprintf("#%05d: %s", round.Episode, round_name)
}

func (round RoundID) RoundName() string {
	if round.Round < 0 || round.Round >= MaxRoundEnum {
		return round_names[0]
	}
	return round_names[round.Round]
}

var round_names = [6]string{
	"[UNKNOWN]",
	"Single!",
	"Double!",
	"Final!",
	"Tiebreaker!!",
	"[printed media]"}

type Board struct {
	RoundID `json:",inline"`
	Columns []CategoryMetadata `json:"columns"`
	Missing []BoardPosition    `json:"missing,omitempty"`
}

type BoardState struct {
	Board   `json:",inline"`
	History []SelectionOutcome `json:"history"`
}

type BoardPosition struct {
	Column uint `json:"column"`
	Index  uint `json:"index"`
}

type BoardSelection struct {
	ChallengeMetadata `json:",inline"`
	BoardPosition     `json:",inline"`
}

type SelectionOutcome struct {
	BoardSelection `json:",inline"`
	Correct        bool  `json:"correct"`
	Delta          Value `json:"delta"`
}

/*
 * CATEGORIES
 */

type CategoryName string

type CategoryIndex map[CategoryName]*CategoryAired

type CategoryMetadata struct {
	Name       CategoryName `json:"title"`
	CategoryID uint64       `json:"catID"`
}

type Category struct {
	CategoryMetadata `json:",inline"`

	ChallengeIDs []ChallengeID `json:"challenges"`
	Media        []MediaRef    `json:"media,omitempty"`
	Comments     string        `json:"comments,omitempty"`
}

type CategoryAired struct {
	CategoryMetadata `json:",inline"`

	Aired ShowDate `json:"aired"`
}

type CategoryThemeEnum int
type CategoryTheme string

const (
	UNKNOWN_CATEGORY CategoryThemeEnum = iota
	CATEGORY_GEOGRAPHY
	CATEGORY_ENTERTAINMENT
	CATEGORY_HISTORY_ROYALTY
	CATEGORY_ART_LITERATURE
	CATEGORY_SCIENCE_NATURE
	CATEGORY_SPORTS_LEISURE
)

/*
 * EPISODES
 */

// A match identifier refers to the unique identifier of the ?-Party database.
// These are not universally unique, they are only certain to be locally unique.
type MatchNumber uint64

// Shows are numbered sequentially based on air date.
// These are historic contests obtained piecemeal from jarchive.com
type MatchID struct {
	Match MatchNumber `json:"match"`

	SeasonStub string `json:"season,omitempty"`
	ShowTitle  string `json:"show_title,omitempty"`
}

type EpisodeIndex map[MatchNumber]*EpisodeMetadata

type EpisodeMetadata struct {
	MatchID   `json:",inline"`
	EpisodeID uint `json:"jaid,omitempty"`

	AiredDate ShowDate `json:"aired,omitempty"`
	TapedDate ShowDate `json:"taped,omitempty"`

	Contestants []ContestantID `json:"contestants,omitempty"`
	Media       []MediaRef     `json:"media,omitempty"`
	Comments    string         `json:"comments,omitempty"`
}

type BoardLayout struct {
	CategoryBitmaps []uint `json:"cat_bitmap"`
}

type EpisodeStats struct {
	SingleCount int `json:"single_count,omitempty"`
	DoubleCount int `json:"double_count,omitempty"`

	TripleStumpers []BoardPosition `json:"triple_stumpers,omitempty"`
}

type ShowDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type ShowDateRange struct {
	From  ShowDate `json:"from,omitempty"`
	Until ShowDate `json:"until,omitempty"`
}

/*
 * CONTESTANTS
 */

type ContestantID struct {
	PK   uint64 `json:"cid"`
	Name string `json:"name,omitempty"`
}

type Contestant struct {
	ContestantID `json:",inline"`

	Name       string     `json:"name"`
	Occupation string     `json:"occupation,omitempty"`
	Residence  string     `json:"residence,omitempty"`
	Notes      string     `json:"notes,omitempty"`
	Media      []MediaRef `json:"media,omitempty"`
}

type Appearance struct {
	ContestantID `json:",inline"`
	Episode      MatchID `json:"episode"`
}

type Career struct {
	ContestantID `json:",inline"`
	Appearances  []MatchID `json:"appearances"`
	Winnings     Value     `json:"winnings"`
}

/*
 * SEASONS
 */

type SeasonStub string

type SeasonID struct {
	Stub  SeasonStub `json:"stub"`
	Title string     `json:"title"`
}

type SeasonIndex map[SeasonStub]*SeasonMetadata

type SeasonMetadata struct {
	Season SeasonID      `json:"season"`
	Title  string        `json:"title,omitempty"`
	Aired  ShowDateRange `json:"aired,omitempty"`

	EpisodeCount   int `json:"episode_count,omitempty"`
	ChallengeCount int `json:"challenge_count,omitempty"`
	TripStumpCount int `json:"tripstump_count,omitempty"`
}

type Season struct {
	SeasonMetadata `json:",inline"`
	EpisodeIndex   `json:",inline"`
	CategoryIndex  `json:",inline"`
}

/*
 * DATA QUALITY
 */

type DataQualityEnum uint8

type DataQuality struct {
	QualityID   DataQualityEnum `json:"dqID"`
	QualityName string          `json:"quality"`
}

type DataQualityJudgement struct {
	ChallengeMetadata `json:",inline"`

	Quality  DataQualityEnum `json:"quality"`
	Comments string          `json:"comments"`
}

/*
 * EMBEDDED SCHEMA AS CUE
 */

// go:embed challenges.cue
var schemaChallenges string

// go:embed rounds.cue
var schemaRounds string

// go:embed episodes.cue
var schemaEpisodes string

// go:embed contestants.cue
var schemaContestants string

// go:embed seasons.cue
var schemaSeasons string
