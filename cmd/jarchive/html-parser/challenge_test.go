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
// github:kevindamm/q-party/cmd/jarchive/html/challenge_test.go

package html

import (
	"strings"
	"testing"

	qparty "github.com/kevindamm/q-party"
	"golang.org/x/net/html"
)

func TestParseChallenge(t *testing.T) {
	html_raw := `<table><tbody><tr><td class="clue">
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

	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatal("failed to parse (html_reader)\n", err)
	}
	clue_td := nextDescendantWithClass(doc, "td", "clue")

	challenge := new(qparty.FullChallenge)
	parseChallenge(clue_td, challenge)

	if challenge.Category != "" {
		t.Errorf("parseChallenge should not set Category property")
	}
	if challenge.Comments != "" {
		t.Errorf("parseChallenge should not set Commentary property")
	}

	if challenge.ChallengeID != 164163 {
		t.Errorf("parseChallenge should get the correct challenge ID from its edit link")
	}

	if challenge.Value.Abs() != 400 {
		t.Errorf("incorrect dollar value '%d'", challenge.Value.Abs())
	}
	if challenge.Value.IsWager() != false {
		t.Error("incorrect dollar value wager (true)")
	}
	if challenge.Value.String() != "$400" {
		t.Errorf("incorrect dollar value string '%s'", challenge.Value.String())
	}
	if challenge.Clue != "Style that got its name from 1925's Exposition Internationale des Arts Decoratifs" {
		t.Errorf("incorrect prompt '%s'", challenge.Clue)
	}
	if challenge.Correct != "art deco" {
		t.Errorf("incorrect response '%s'", challenge.Correct)
	}
}

func TestParseEmptyChallenge(t *testing.T) {
	html_raw := `
<table><tr><td class="clue">
    </td></tr></table>`

	challenge := new(qparty.FullChallenge)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("failed to parse (html_reader)\n%s", err)
	}
	td := nextDescendantWithClass(doc, "td", "clue")
	parseChallenge(td, challenge)

	if challenge.Category != "" || challenge.Comments != "" ||
		challenge.Clue != "" || challenge.Correct != "" {
		t.Error("incorrect (non-zero) result from empty challenge")
	}
	if challenge.ChallengeID != 0 {
		t.Error("incorrect (non-zero) result from empty challenge")
	}

	if innerText(doc) != "" {
		t.Error("expected empty string for empty challenge, got", innerText(doc))
	}
}

func TestParseImageChallenge(t *testing.T) {
	html_raw := `  <table><tr><td class="clue">
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

	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("failed to parse (html_reader)\n%s", err)
	}
	clue_td := nextDescendantWithClass(doc, "td", "clue")

	challenge := new(qparty.FullChallenge)
	parseChallenge(clue_td, challenge)

	expected_prompt := "A Veronica is a movement done in [this](/media/2004-07-19_DJ_11.jpg) sport, popular in Mexico"
	if challenge.Clue != expected_prompt {
		t.Errorf("incorrect prompt, got\n%s\n%s\nexpected", challenge.Clue, expected_prompt)
	}
	if challenge.Media[0].MimeType != qparty.MediaImageJPG {
		t.Error("incorrect media type")
	}
	if challenge.Media[0].MediaURL != "/media/2004-07-19_DJ_11.jpg" {
		t.Errorf("incorrect media URL\n%s", challenge.Media[0].MediaURL)
	}
}

func TestParseAudioWageringChallenge(t *testing.T) {
	html_raw := `<table><tr><td class="clue">
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

	challenge := new(qparty.FullChallenge)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("failed to parse (html_reader)\n%s", err)
	}
	td := nextDescendantWithClass(doc, "td", "clue")
	parseChallenge(td, challenge)

	if challenge.Category != "" {
		t.Errorf("parseChallenge should not set Category property")
	}
	if challenge.Comments != "" {
		t.Errorf("parseChallenge should not set Commentary property")
	}

	if challenge.Value.Abs() != 1900 {
		t.Errorf("incorrect dollar value '%d'", challenge.Value.Abs())
	}
	if challenge.Value.IsWager() != true {
		t.Error("incorrect dollar value wager (true)")
	}
	if challenge.Value.String() != "$-1900" {
		t.Errorf("incorrect dollar value string '%s'", challenge.Value.String())
	}
	expected_text := "*\"I dream of Jeanie with the light brown hair / Borne, like a vapor...\"*\n\nWhat we call [this](/media/1984-10-02_DJ_09.mp3) singing _sans_ accompaniment:"
	if challenge.Clue != expected_text {
		t.Errorf("incorrect prompt '%s' !=\n'%s'", challenge.Clue, expected_text)
	}
	if len(challenge.Media) != 1 || challenge.Media[0].MimeType != qparty.MediaAudioMP3 {
		t.Errorf("did not detect audio media type in challenge")
	}
	if challenge.Correct != "a capella" {
		t.Errorf("incorrect response '%s'", challenge.Correct)
	}
}

func TestParseVideoChallenge(t *testing.T) {
	html_raw := `<table><tr><td class="clue">
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

	challenge := new(qparty.FullChallenge)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("failed to parse (html_reader)\n%s", err)
	}
	td := nextDescendantWithClass(doc, "td", "clue")
	parseChallenge(td, challenge)

	expected_prompt := "( [Cheryl of the Clue Crew drags a mustache back onto Alex.](/media/2004-07-19_DJ_11.mp4) ) [Introduced](/media/2004-07-19_DJ_11.jpg) in 1990, it's the [program](/media/2004-07-19_DJ_11a.jpg) I'm using to create an [unusual image](/media/2004-07-19_DJ_11b.jpg)"
	if challenge.Clue != expected_prompt {
		t.Errorf("incorrect prompt, got\n%s\n%s\nexpected", challenge.Clue, expected_prompt)
	}
	expected_media := []qparty.Media{
		{MimeType: qparty.MediaVideoMP4, MediaURL: "/media/2004-07-19_DJ_11.mp4"},
		{MimeType: qparty.MediaImageJPG, MediaURL: "/media/2004-07-19_DJ_11.jpg"},
		{MimeType: qparty.MediaImageJPG, MediaURL: "/media/2004-07-19_DJ_11a.jpg"},
		{MimeType: qparty.MediaImageJPG, MediaURL: "/media/2004-07-19_DJ_11b.jpg"},
	}
	for i, expected := range expected_media {
		if challenge.Media[i].MimeType != expected.MimeType {
			t.Errorf("incorrect media type %s\n%s", challenge.Media[i].MimeType, expected.MimeType)
		}
		if challenge.Media[i].MediaURL != expected.MediaURL {
			t.Errorf("incorrect media URL\n%s\n%s", challenge.Media[i].MediaURL, expected.MediaURL)
		}
	}
}

func TestParseFinalChallenge(t *testing.T) {
	html_raw := `<div id="final_jeopardy_round">
               <h2>Final Jeopardy! Round</h2>

<table class="final_round">
  <tr>
    <td class="category">
      
  <table>
    <tr><td class="category_name">WORLD AFFAIRS</td></tr>
    <tr><td class="category_comments">comment this</td></tr>
  </table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td id="clue_FJ" class="clue_text">In 1963 these 3 nations signed the Nuclear Test Ban Treaty</td>
    <td id="clue_FJ_r" class="clue_text" style="display:none;"><em class="correct_response">the U.S., U.S.S.R. & the U.K.</em></td>
  </tr>
</table>
    </td>
   </tr>
</table>
              </div>`

	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatal("failed to parse (html_reader)\n", err)
	}
	final_round := nextDescendantWithClass(doc, "table", "final_round")
	final_challenge, err := parseFinalChallenge(final_round)
	if err != nil {
		t.Fatal("failed to parse (final challenge)\n", err)
	}

	if final_challenge.Category != "WORLD AFFAIRS" {
		t.Errorf("incorrect category '%s'", final_challenge.Category)
	}
	if final_challenge.Comments != "comment this" {
		t.Errorf("incorrect comment '%s'", final_challenge.Comments)
	}

	if final_challenge.Clue != "In 1963 these 3 nations signed the Nuclear Test Ban Treaty" {
		t.Errorf("incorrect prompt '%s'", final_challenge.Clue)
	}
	if final_challenge.Correct != "the U.S., U.S.S.R. & the U.K." {
		t.Errorf("incorrect response '%s'", final_challenge.Correct)
	}
}
