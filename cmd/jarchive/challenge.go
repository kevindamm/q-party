package main

import "golang.org/x/net/html"

// The de-normed representation as found in some datasets, e.g. on Kaggle.
type JArchiveChallenge struct {
	Category string            `json:"category"`
	AirDate  `json:"air_date"` // YYYY-MM-DD

	Value    DollarValue  `json:"value"` // /$(\d+)/
	Question string       `json:"prompt"`
	Answer   string       `json:"correct"` // excluding "what is..." preface
	Accept   []string     `json:"accept,omitempty"`
	Round    EpisodeRound `json:"round"`
}

type JArchiveFinalChallenge JArchiveChallenge

func (challenge *JArchiveChallenge) parseChallenge(node *html.Node) {
	// TODO
}
