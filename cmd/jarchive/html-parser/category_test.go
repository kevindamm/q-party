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
// github:kevindamm/q-party/cmd/jarchive/html/category_test.go

package html

import (
	"strings"
	"testing"

	qparty "github.com/kevindamm/q-party"
	"golang.org/x/net/html"
)

func TestParseCategoryHeader(t *testing.T) {
	html_raw := `<table><tr><td class="category">
<table>
  <tr><td class="category_name">HANS CHRISTIAN ANDERSEN</td></tr>
  <tr><td class="category_comments">(Alex: We're celebrating his bicentennial this year!)</td></tr>
</table>
    </td></tr></table>`
	category := new(qparty.FullCategory)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatal("failed to parse (html_reader)\n", err)
	}
	td := nextDescendantWithClass(doc, "td", "category")
	err = parseCategoryHeader(td, category)
	if err != nil {
		t.Fatal("failed to parse the category names", err)
	}

	if category.Title != "HANS CHRISTIAN ANDERSEN" {
		t.Error("unexpected category title", category.Title)
	}
	if category.Comments != "(Alex: We're celebrating his bicentennial this year!)" {
		t.Error("unexpected category comment", category.Comments)
	}
}

func TestParseCategoryChallenge(t *testing.T) {
	html_raw := `<table><tr><td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_1_1', 'clue_J_1_1_r', 'clue_J_1_1_stuck')" onmouseout="toggle('clue_J_1_1_r', 'clue_J_1_1', 'clue_J_1_1_stuck')" onclick="togglestick('clue_J_1_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_1_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32559" title="Suggest a correction for this clue" rel="nofollow">12</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_1_1" class="clue_text">expected prompt</td>
    <td id="clue_J_1_1_r" class="clue_text" style="display:none;"><em class="correct_response">expected correct response</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
	</td></tr></table>`
	category := new(qparty.FullCategory)
	category.Title = "categoryname"

	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatal("failed to parse (html_reader)\n", err)
	}
	td := nextDescendantWithClass(doc, "td", "clue")
	err = parseCategoryChallenge(td, category)
	if err != nil {
		t.Fatal("failed to parse contents of <td class='clue'>",
			"\n", err)
	}

	if category.Challenges[0].Category != string(category.Title) {
		t.Error("category mismatch",
			category.Challenges[0].Category, "!=", category.Title)
	}
	if category.Challenges[0].Clue != "expected prompt" {
		t.Error("prompt mismatch", category.Challenges[0].Clue, "expected prompt")
	}
	if category.Challenges[0].Correct != "expected correct response" {
		t.Error("response mismatch", category.Challenges[0].Correct, "expected correct response")
	}
}
