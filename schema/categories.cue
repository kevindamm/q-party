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
// github:kevindamm/q-party/schema/categories.cue

package schema

#CategoryID: string

#CategoryIndex: [=~#CategoryID]: #CategoryAired

#CategoryMetadata: {
  catID: #CategoryID
  title!: string
  ...
}

// A category instance must have a title and
// may have any number of challenges (typically five).
#Category: #CategoryMetadata & {
  challenges: [...#ChallengeMetadata]

  media?: [...#MediaRef]
  comments?: string
}

#CategoryAired: #CategoryMetadata & {
  aired: #ShowDate
}

#CategoryThemeEnum: int & >=0
#CategoryTheme: string

_cat_themes: [...#CategoryTheme] & [
	"[UNKNOWN]",
	"Geography",
	"Entertainment",
	"History & Royalty",
	"Art & Literature",
	"Science & Nature",
	"Sports & Leisure",
]
