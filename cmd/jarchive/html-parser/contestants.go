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
// github:kevindamm/q-party/cmd/jarchive/html/contestants.go

package html

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	qparty "github.com/kevindamm/q-party"
	"golang.org/x/net/html"
)

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
	if len(contestants) > 3 {
		// Some all-star games have 9 contestants in rotation of threes; we can skip
		// parsing the contestants for those, contestant details are optional.
		log.Print("SKIPPING CONTESTANTS this is one of the weird ones.")
		return
	}
	for i, contestant := range contestants {
		err := parseContestant(contestant, &episode.Contestants[i])
		if err != nil {
			log.Fatal("failed to parse contestant", err)
		}
	}
}

// Parse a single <p class="contestants"> subtree into a [JArchiveContestant].
func parseContestant(root *html.Node, contestant *qparty.Contestant) error {
	link := nextDescendantWithClass(root, "a", "")
	contestant.Name = innerText(link)
	for _, attr := range link.Attr {
		if attr.Key == "href" {
			jcid, err := strconv.Atoi(strings.Split(attr.Val, "=")[1])
			if err != nil {
				return err
			}
			contestant.UCID = qparty.UCID(jcid)
		}
	}
	textNode := link.NextSibling
	if textNode.Type == html.TextNode {
		contestant.Biography = textNode.Data[2:]
	}
	return nil
}
