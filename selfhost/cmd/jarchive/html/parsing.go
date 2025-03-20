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
// github:kevindamm/q-party/cmd/jarchive/html/parsing.go

package html

// Main entry points for parsing entire HTML files and some utility funcstions
// for navigating the document after parsing it with net/html.
// Most of the detailed parsing routines are alongside, typically also methods
// on, the other types in this module (episode, board, category, challenge).

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	qparty "github.com/kevindamm/q-party"
	"golang.org/x/net/html"
)

func LoadEpisodeHTML(html_path string) *qparty.FullEpisode {
	// Parse the episode's HTML to get the show and challenge details.
	reader, err := os.Open(html_path)
	if err != nil {
		// Unlikely, we know the file exists at this point, but state changes...
		log.Print("ERROR: ", err)
		return nil
	}
	defer reader.Close()

	episode := ParseEpisode(reader)
	return episode
}

func ParseEpisode(html_reader io.Reader) *qparty.FullEpisode {
	episode := new(qparty.FullEpisode)
	doc, err := html.Parse(html_reader)
	if err != nil {
		log.Fatal("error parsing HTML\n", err)
	}

	child := doc.FirstChild
	for child != nil {
		if child.Type == html.DocumentNode ||
			(child.Type == html.ElementNode && child.Data == "html") ||
			(child.Type == html.ElementNode && child.Data == "body") {
			child = child.FirstChild
			continue
		}
		if divWithID(child) == "content" {
			parseContent(child, episode)
			break
		}
		child = child.NextSibling
	}

	return episode
}

// Returns a list of the direct children elements that have the indicated class.
// If elType is not the empty string, only returns elements of that type.  Only
// elements are returned; text nodes and other non-element children are ignored.
func childrenWithClass(node *html.Node, elType string, elClass string) []*html.Node {
	matching_children := make([]*html.Node, 0)
	child := node.FirstChild

	for child != nil {
		if child.Type == html.ElementNode &&
			(elType == "" || child.Data == elType) {
			if elClass == "" {
				matching_children = append(matching_children, child)
			} else {
				for _, attr := range child.Attr {
					if attr.Key == "class" {
						if hasClass(attr.Val, elClass) {
							matching_children = append(matching_children, child)
						}
						break
					}
				}
			}
		}
		child = child.NextSibling
	}
	return matching_children
}

// Searches recursively through descendents (DFS) looking for the next element
// with the indicated type and class.  If class=="" then any (or no) class will
// satisfy the search.  It returns the first matching subelement, depth first.
// If there is no matching element (and class) then nil is returned instead.
func nextDescendantWithClass(node *html.Node, elType string, elClass string) *html.Node {
	var found *html.Node = nil
	var recursiveFind func(*html.Node, string, string)
	recursiveFind = func(next *html.Node, elType string, elClass string) {
		child := next.FirstChild
		for child != nil {
			if child.Type == html.ElementNode &&
				(elType == "" || child.Data == elType) {
				if elClass == "" {
					found = child
					return
				}
				for _, attr := range child.Attr {
					if attr.Key == "class" {
						if hasClass(attr.Val, elClass) {
							found = child
							return
						}
						break
					}
				}
			}
			if child.FirstChild != nil {
				recursiveFind(child, elType, elClass)
				if found != nil {
					return
				}
			}
			child = child.NextSibling
		}
	}
	recursiveFind(node, elType, elClass)
	return found
}

// Similar to [nextDescendantWithClass] but only looks at adjacent siblings to
// the element being passed as the first argument.  Similarly, elClass can be
// an empty string, matching against any element of the indicated type, without
// a class defined on it, or with any arbitrary class or classes.
func nextSiblingWithClass(el *html.Node, elType string, elClass string) *html.Node {
	nextEl := el.NextSibling
	for nextEl != nil {
		if nextEl.Type != html.ElementNode || nextEl.Data != elType {
			nextEl = nextEl.NextSibling
			continue
		}
		if elClass == "" {
			return nextEl
		}
		for _, attr := range nextEl.Attr {
			if attr.Key == "class" {
				if hasClass(attr.Val, elClass) {
					return nextEl
				}
				break
			}
		}
		nextEl = nextEl.NextSibling
	}
	return nil
}

// Handles the situation where an element has multiple classes.
func hasClass(haystack string, needle string) bool {
	for _, class := range strings.Split(haystack, " ") {
		if class == needle {
			return true
		}
	}
	return false
}

// Returns the ID of the node if it is a <div> element,
// otherwise returns the empty string.
func divWithID(node *html.Node) string {
	if node.Type != html.ElementNode || node.Data != "div" {
		// Not a <div>.
		return ""
	}
	for _, attr := range node.Attr {
		if attr.Key == "id" {
			return attr.Val
		}
	}

	// Is a <div> but has no ID.
	return ""
}

func innerText(node *html.Node) string {
	text := make([]string, 0)
	var recursiveFind func(*html.Node)

	recursiveFind = func(node *html.Node) {
		if node == nil {
			return
		}
		child := node.FirstChild
		for child != nil {
			if child.Type == html.TextNode {
				text = append(text, strings.ReplaceAll(child.Data, "\n", " "))
			}
			if child.FirstChild != nil {
				recursiveFind(child)
			}
			child = child.NextSibling
		}
	}
	recursiveFind(node)
	flattened := strings.ReplaceAll(strings.Join(text, " "), "  ", " ")
	return strings.Trim(flattened, " \t\r\n")
}

func parseIntoMarkdown(root *html.Node) (string, []qparty.Media) {
	prompt := ""
	media := make([]qparty.Media, 0)
	var recursiveGather func(*html.Node)

	recursiveGather = func(root *html.Node) {
		child := root.FirstChild
		for child != nil {
			if child.Type == html.TextNode {
				// Concatenate all immediate children that are text nodes.
				prompt += child.Data
			} else if child.Type == html.ElementNode {
				if child.Data == "a" {
					href, text := "", ""
					// Parse link destination.
					for _, attr := range child.Attr {
						if attr.Key == "href" {
							href = attr.Val
							break
						}
					}
					// Parse link text (<a>...</a> contents).
					if child.FirstChild != nil && child.FirstChild.Type == html.TextNode {
						text = child.FirstChild.Data
					}

					// represent as Media and as markdown within the prompt.
					media_asset := MakeMedia(href)
					prompt += fmt.Sprintf("%s [%d] ", text, len(media))
					media = append(media, media_asset)
				} else if child.Data == "u" {
					// Recursively collect the text and media of the prompt.
					prompt += " _"
					recursiveGather(child)
					prompt += "_ "
				} else if child.Data == "i" {
					// Recursively collect the text and media of the prompt.
					prompt += " *"
					recursiveGather(child)
					prompt += "* "
				} else if child.Data == "b" {
					// Recursively collect the text and media of the prompt.
					prompt += " **"
					recursiveGather(child)
					prompt += "** "
				} else if child.Data == "del" {
					prompt += " ~~"
					recursiveGather(child)
					prompt += "~~ "
				} else if child.Data == "span" {
					recursiveGather(child)
				} else if child.Data == "big" {
					prompt += "<big>"
					recursiveGather(child)
					prompt += "</big>"
				} else if child.Data == "small" {
					prompt += "<small>"
					recursiveGather(child)
					prompt += "</small>"
				} else if child.Data == "sub" {
					prompt += "<sub>"
					recursiveGather(child)
					prompt += "</sub>"
				} else if child.Data == "sup" {
					prompt += "<sup>"
					recursiveGather(child)
					prompt += "</sup>"
				} else if child.Data == "br" {
					// pass, safe to ignore; insert a newline if it's a double-<br/>.
					if child.NextSibling != nil &&
						child.NextSibling.Type == html.ElementNode &&
						child.NextSibling.Data == "br" {
						child = child.NextSibling
						prompt = strings.TrimRight(prompt, " ")
						prompt += "\n\n"
					} else {
						prompt += " "
					}
				} else {
					log.Fatalf("unexpected element type (%s) in Prompt", child.Data)
				}
			} else {
				log.Fatalf("unexpected node (%d %s) in Prompt",
					child.Type, child.Data)
			}
			// Visit all immediate children.
			child = child.NextSibling
		}
	}
	recursiveGather(root)

	prompt = strings.Trim(strings.ReplaceAll(prompt, "  ", " "), "\t \r\n")
	return prompt, media
}
