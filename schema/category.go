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
// github:kevindamm/q-party/schema/category.go

package schema

import _ "embed"

// go:embed category.cue
var schemaCategories string

type CategoryName string

type CategoryIndex map[CategoryName][]CategoryAired

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
