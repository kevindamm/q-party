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
// github:kevindamm/q-party/cmd/jarchive/parsing.go

package main

// Main entry points for parsing entire HTML files and some utility funcstions
// for navigating the document after parsing it with net/html.
// Most of the detailed parsing routines are alongside, typically also methods
// on, the other types in this module (episode, board, category, challenge).

import (
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func ParseEpisode(jeid JEID, html_reader io.Reader) *JArchiveEpisode {
	episode := new(JArchiveEpisode)
	episode.JEID = jeid
	doc, err := html.Parse(html_reader)
	if err != nil {
		log.Fatalf("error parsing HTML of %s\n\n%s", jeid, err)
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
			episode.parseContent(child)
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
			for _, attr := range child.Attr {
				if attr.Key == "class" {
					if hasClass(attr.Val, elClass) {
						matching_children = append(matching_children, child)
						break
					}
					break
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
				text = append(text, child.Data)
			}
			if child.FirstChild != nil {
				recursiveFind(child)
			}
			child = child.NextSibling
		}
	}
	recursiveFind(node)

	return strings.ReplaceAll(strings.Join(text, " "), "  ", " ")
}

func parseTimeYYYYMMDD(yyyy, mm, dd []byte) time.Time {
	year, err := strconv.Atoi(string(yyyy))
	if err != nil {
		log.Fatal(yyyy, err)
	}
	month, err := strconv.Atoi(string(mm))
	if err != nil {
		log.Fatal(mm, err)
	}
	day, err := strconv.Atoi(string(dd))
	if err != nil {
		log.Fatal(dd, err)
	}
	return time.Date(year, time.Month(month), day, 10, 8, 0, 0, time.UTC)
}
