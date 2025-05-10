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

	"github.com/kevindamm/q-party/schema"
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

func TestParseImageChallenge(t *testing.T) {
	html_string := `  <table><tr><td class="clue">
<table>
  <tr>
    <td>
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_2_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=146846" title="Suggest a correction for this clue" rel="nofollow">4</a></td>
          </tr>
        </table>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_2_4" class="clue_text">A Veronica is a movement done in <a href="http://www.j-archive.com/media/2004-07-19_DJ_11.jpg" target="_blank">this</a> sport, popular in Mexico</td>
    <td id="clue_DJ_2_4_r" class="clue_text" style="display:none;"><em class="correct_response">bullfighting</em><br /><br /><table width="100%"><tr><td class="right">Tim</td></tr></table></td>
  </tr>
</table>
    </td></tr></table>`
	html_reader := strings.NewReader(html_string)
	dom, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("??? failed to parse test input\n%s\n\n%s\n",
			err, html_string)
	}
	clue_td := jarchive.NextDescendantWithClass(dom, "td", "clue")

	cat_name := "CATEGORY"
	expected := jarchive.NewChallenge(cat_name, 146846)
	expected.Clue = "A Veronica is a movement done in this [0] sport, popular in Mexico"
	expected.Value = 1600
	expected.Media = []schema.MediaRef{
		{MimeType: schema.MediaImageJPG, MediaURL: "2004-07-19_DJ_11.jpg"},
	}

	parsed, err := jarchive.ParseChallenge(clue_td, cat_name)
	if err != nil {
		t.Error(err)
	}

	if err := equalChallenge(parsed, expected); err != nil {
		t.Error(err)
	}
}

func TestParseAudioWageringChallenge(t *testing.T) {
	html_string := `<table><tr><td class="clue">
<table>
  <tr>
    <td>
        <table class="clue_header">
          <tr>
            <td class="clue_value_daily_double">DD: $1,900</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=19078" title="Suggest a correction for this clue" rel="nofollow">9</a></td>
          </tr>
        </table>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_3_4" class="clue_text"><i>"I dream of Jeanie with the light brown hair /<br />Borne, like a vapor..."</i><br /><br />What we call <a href="http://www.j-archive.com/media/1984-10-02_DJ_09.mp3">this</a> singing <u>sans</u> accompaniment:</td>
    <td id="clue_DJ_3_4_r" class="clue_text" style="display:none;"><em class="correct_response">a capella</em><br /><br /><table width="100%"><tr><td class="right">Mary</td></tr></table></td>
  </tr>
</table>
    </td></tr></table>`

	html_reader := strings.NewReader(html_string)
	dom, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("??? failed to parse test input\n%s\n\n%s\n",
			err, html_string)
	}
	clue_td := jarchive.NextDescendantWithClass(dom, "td", "clue")

	cat_name := "CATEGORY_NAME"
	expected := jarchive.NewChallenge(cat_name, 19078)
	expected.Wager = 1900
	expected.Clue = "*\"I dream of Jeanie with the light brown hair / Borne, like a vapor...\"*\n\nWhat we call this [0] singing _sans_ accompaniment:"
	expected.Media = []schema.MediaRef{{
		MimeType: schema.MediaAudioMP3,
		MediaURL: "1984-10-02_DJ_09.mp3"}}
	expected.Correct = []string{"a capella"}

	parsed, err := jarchive.ParseChallenge(clue_td, cat_name)
	if err != nil {
		t.Error(err)
	}

	if err := equalChallenge(parsed, expected); err != nil {
		t.Error(err)
	}
}

func TestParseMultiVideoChallenge(t *testing.T) {
	html_string := `<table><tr><td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_3_4', 'clue_DJ_3_4_r', 'clue_DJ_3_4_stuck')" onmouseout="toggle('clue_DJ_3_4_r', 'clue_DJ_3_4', 'clue_DJ_3_4_stuck')" onclick="togglestick('clue_DJ_3_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_3_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=146851" title="Suggest a correction for this clue" rel="nofollow">11</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_3_4" class="clue_text">(<a href="http://www.j-archive.com/media/2004-07-19_DJ_11.mp4">Cheryl of the Clue Crew drags a mustache back onto Alex.</a>) <a href="http://www.j-archive.com/media/2004-07-19_DJ_11.jpg" target="_blank">Introduced</a> in 1990, it's the <a href="http://www.j-archive.com/media/2004-07-19_DJ_11a.jpg" target="_blank">program</a> I'm using to create an <a href="http://www.j-archive.com/media/2004-07-19_DJ_11b.jpg" target="_blank">unusual image</a></td>
    <td id="clue_DJ_3_4_r" class="clue_text" style="display:none;">(Alex: Not so unusual.)<br /><br /><em class="correct_response">Photoshop</em><br /><br /><table width="100%"><tr><td class="right">Ken</td></tr></table></td>
  </tr>
</table>
    </td></tr></table>`
	html_reader := strings.NewReader(html_string)
	dom, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("??? failed to parse test input\n%s\n\n%s\n",
			err, html_string)
	}
	clue_td := jarchive.NextDescendantWithClass(dom, "td", "clue")

	cat_name := "CAT E GORY"
	expected := jarchive.NewChallenge(cat_name, 146851)
	expected.Value = 1600
	expected.Clue = "(Cheryl of the Clue Crew drags a mustache back onto Alex. [0] ) Introduced [1] in 1990, it's the program [2] I'm using to create an unusual image [3]"
	expected.Media = []schema.MediaRef{
		{MimeType: schema.MediaVideoMP4, MediaURL: "/2004-07-19_DJ_11.mp4"},
		{MimeType: schema.MediaImageJPG, MediaURL: "/2004-07-19_DJ_11.jpg"},
		{MimeType: schema.MediaImageJPG, MediaURL: "/2004-07-19_DJ_11a.jpg"},
		{MimeType: schema.MediaImageJPG, MediaURL: "/2004-07-19_DJ_11b.jpg"},
	}

	parsed, err := jarchive.ParseChallenge(clue_td, cat_name)
	if err != nil {
		t.Error(err)
	}

	if err := equalChallenge(parsed, expected); err != nil {
		t.Error(err)
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
			return fmt.Errorf("different correct/accepted answers #%d: %s != %s",
				i, have.Correct[i], expect.Correct[i])
		}
	}

	if len(have.Media) != len(expect.Media) {
		return fmt.Errorf("Different media references %v\n!=\n%v", have.Media, expect.Media)
	}
	for i := range len(have.Media) {
		if have.Media[i].MimeType != expect.Media[i].MimeType {
			return fmt.Errorf("different mime types for media #%d; %s != %s",
				i, have.Media[i].MimeType, expect.Media[i].MimeType)
		}
		if have.Media[i].MediaURL != expect.Media[i].MediaURL {
			return fmt.Errorf("different media URLs for media #%d; %s != %s",
				i, have.Media[i].MediaURL, expect.Media[i].MediaURL)
		}
	}

	return nil
}
