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
// github:kevindamm/q-party/cmd/fetch/jarchive/challenges_test.go

package main_test

import (
	"fmt"
	"strings"
	"testing"

	jarchive "github.com/kevindamm/q-party/selfhost/cmd/fetch/jarchive"
	"golang.org/x/net/html"
)

func TestParseChallenge(t *testing.T) {
	html_string := `<table><tbody><tr><td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_5_2', 'clue_J_5_2_r', 'clue_J_5_2_stuck')" onmouseout="toggle('clue_J_5_2_r', 'clue_J_5_2', 'clue_J_5_2_stuck')" onclick="togglestick('clue_J_5_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_5_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=164163" title="Suggest a correction for this clue" rel="nofollow">22</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_5_2" class="clue_text">Style that got its name from 1925's Exposition Internationale des Arts Decoratifs</td>
    <td id="clue_J_5_2_r" class="clue_text" style="display:none;"><em class="correct_response">art deco</em><br /><br /><table width="100%"><tr><td class="right">Stefanie</td></tr></table></td>
  </tr>
</table>
    </td></tr></tbody></table>`
	html_reader := strings.NewReader(html_string)
	dom, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("??? failed to parse test input\n%s\n\n%s\n",
			err, html_string)
	}
	clue_td := jarchive.NextDescendantWithClass(dom, "td", "clue")
	cat_name := "CATEGORY"
	expected := jarchive.NewChallenge(cat_name, 164163)
	expected.Value = 400
	expected.Clue = "Style that got its name from 1925's Exposition Internationale des Arts Decoratifs"
	expected.Correct = []string{"art deco"}

	parsed, err := jarchive.ParseChallenge(clue_td, cat_name)
	if err != nil {
		t.Error(err)
	}

	if err := equalChallenge(parsed, expected); err != nil {
		t.Error(err)
	}
}

func TestParseEmptyChallenge(t *testing.T) {
	html_string := `<table><tr><td class="clue">
    </td></tr></table>`
	html_reader := strings.NewReader(html_string)
	dom, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("??? failed to parse empty test input\n%s\n", err)
	}

	cat_name := "CATEGORY"
	clue_td := jarchive.NextDescendantWithClass(dom, "td", "clue")
	parsed, err := jarchive.ParseChallenge(clue_td, cat_name)
	if err != nil {
		t.Fatal(err)
	}

	if parsed.Category != "" || parsed.Comments != "" ||
		parsed.Clue != "" || len(parsed.Correct) != 0 {
		t.Error("incorrect (non-zero) result from empty challenge")
	}
	if parsed.ChallengeID != 0 {
		t.Error("incorrect (non-zero) result from empty challenge")
	}
}

func equalChallenge(have, expect *jarchive.JarchiveChallenge) error {
	if have.Category != expect.Category {
		if expect.Category == "" {
			return fmt.Errorf("category name was set %s not expected to be set", have.Category)
		} else {
			return fmt.Errorf("category has unexpected value %s != %s", have.Category, expect.Category)
		}
	}
	if have.Comments != expect.Comments {
		if expect.Comments == "" {
			return fmt.Errorf("comments were set %s not expected to be set", have.Comments)
		} else {
			return fmt.Errorf("comments has unexpected value %s != %s", have.Comments, expect.Comments)
		}
	}

	if have.ChallengeID != expect.ChallengeID {
		if expect.ChallengeID == 0 {
			return fmt.Errorf("challenge ID has a value %d not expecting any value", have.ChallengeID)
		} else {
			return fmt.Errorf("challenge ID is unexpected value %d != %d", have.ChallengeID, expect.ChallengeID)
		}
	}

	if have.Value != expect.Value {
		if expect.Value == 0 {
			return fmt.Errorf("challenge value is %d not expecting any value", have.Value)
		} else {
			return fmt.Errorf("challenge value is unexpected %d != %d", have.Value, expect.Value)
		}
	}

	if have.Clue != expect.Clue {
		if expect.Clue == "" {
			return fmt.Errorf("challenge clue has a value %s not expecting a value", have.Clue)
		} else {
			return fmt.Errorf("challenge clue is unexpected %s != %s", have.Clue, expect.Clue)
		}
	}

	if len(have.Correct) != len(expect.Correct) {
		return fmt.Errorf("Different correct/accepted answers\n%v\n%v",
			have.Correct, expect.Correct)
	}
	for i := 0; i < len(have.Correct); i++ {
		if have.Correct[i] != expect.Correct[i] {
			return fmt.Errorf("different correct/accepted answers %s != %s", have.Correct[i], expect.Correct[i])
		}
	}

	return nil
}
