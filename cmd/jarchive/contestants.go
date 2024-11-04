package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type JArchiveContestant struct {
	JCID `json:"id,omitempty"`
	Name string `json:"name"`
	Bio  string `json:"comment"`
}

// Unique numeric value for identifying customers in the archive.
type JCID int

func (id JCID) String() string {
	return fmt.Sprintf("%d", int(id))
}

// Parses the numeric value from a string.
// Fatal error if the value cannot be converted into a number.
func MustParseJCID(numeric string) JCID {
	id, err := strconv.Atoi(numeric)
	if err != nil {
		log.Fatalf("failed to parse JCID from string '%s'\n%s", numeric, err)
	}
	return JCID(id)
}

// Parse all three contestants, storing in the episode's metadata.
func (episode *JArchiveEpisode) parseContestants(root *html.Node) {
	td_parent := nextDescendantWithClass(root, "p", "contestants").Parent
	contestants := childrenWithClass(td_parent, "p", "contestants")
	for i, contestant := range contestants {
		err := parseContestant(contestant, &episode.Contestants[i])
		if err != nil {
			log.Fatal("failed to parse contestant", err)
		}
	}
}

// Parse a single <p class="contestants"> subtree into a [JArchiveContestant].
func parseContestant(root *html.Node, contestant *JArchiveContestant) error {
	link := nextDescendantWithClass(root, "a", "")
	contestant.Name = innerText(link)
	for _, attr := range link.Attr {
		if attr.Key == "href" {
			jcid, err := strconv.Atoi(strings.Split(attr.Val, "=")[1])
			if err != nil {
				return err
			}
			contestant.JCID = JCID(jcid)
		}
	}
	textNode := link.NextSibling
	if textNode.Type == html.TextNode {
		contestant.Bio = textNode.Data[2:]
	}
	return nil
}
