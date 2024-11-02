package main

import (
	"fmt"

	"golang.org/x/net/html"
)

type JArchiveContestant struct {
	JCID    int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (episode *JArchiveEpisode) parseContestants(div *html.Node) {
	// not necessary but could be nice for tracking a contestant's career
	fmt.Println("TODO parse contestants")
}
