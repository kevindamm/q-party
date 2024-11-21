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
// github:kevindamm/q-party/category.go

package qparty

type CategoryMetadata struct {
	Title    string          `json:"title"`
	Theme    CategoryTheme   `json:"theme,omitempty"`
	Episodes []CategoryAired `json:"episodes"`
}

type CategoryAired struct {
	EpisodeID `json:"episode_id"`
	ShowDate  `json:",inline"`
}

type Category struct {
	Metadata   CategoryMetadata    `json:"metadata,omitempty,inline"`
	Challenges []ChallengeMetadata `json:"challenges,omitempty"`
}

type FullCategory struct {
	Title      string          `json:"title"`
	Comments   string          `json:"comments,omitempty"`
	Challenges []FullChallenge `json:"challenges"`
}

func (category FullCategory) Complete() bool {
	for _, challenge := range category.Challenges {
		if challenge.ChallengeID == 0 {
			return false
		}
	}
	return true
}

// Proposal for category breakdown based on Trivial Pursuit classic categories.
type CategoryTheme string

const (
	ThemeUnknown        CategoryTheme = ""
	ThemeGeography      CategoryTheme = "Geography"
	ThemeEntertainment  CategoryTheme = "Entertainment"
	ThemeHistoryRoyalty CategoryTheme = "History & Royalty"
	ThemeArtLiterature  CategoryTheme = "Art & Literature"
	ThemeScienceNature  CategoryTheme = "Science & Nature"
	ThemeSportsLeisure  CategoryTheme = "Sports & Leisure"
)
