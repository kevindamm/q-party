package main

import (
	"strings"
	"testing"

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

	challenge := new(JArchiveChallenge)
	challenge.parseChallenge(clue_td)

	if challenge.Round != ROUND_UNKNOWN {
		t.Errorf("incorrect round (%s)", challenge.Round)
	}
	if challenge.Category != "" {
		t.Errorf("parseChallenge should not set Category property")
	}
	if challenge.Commentary != "" {
		t.Errorf("parseChallenge should not set Commentary property")
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
	if challenge.Prompt != "Style that got its name from 1925's Exposition Internationale des Arts Decoratifs" {
		t.Errorf("incorrect prompt '%s'", challenge.Prompt)
	}
	if challenge.Correct != "art deco" {
		t.Errorf("incorrect response '%s'", challenge.Correct)
	}
}

func TestParseEmptyChallenge(t *testing.T) {
	html_raw := `
<table><tr><td class="clue">
    </td></tr></table>`

	challenge := new(JArchiveChallenge)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("failed to parse (html_reader)\n%s", err)
	}
	td := nextDescendantWithClass(doc, "td", "clue")
	challenge.parseChallenge(td)

	if challenge.Category != "" || challenge.Commentary != "" ||
		challenge.Round != ROUND_UNKNOWN || challenge.Value != 0 ||
		challenge.Prompt != "" || challenge.Correct != "" {
		t.Error("incorrect (non-zero) result from empty challenge")
	}
	if !challenge.IsEmpty() {
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

	challenge := new(JArchiveChallenge)
	challenge.parseChallenge(clue_td)

	expected_prompt := "A Veronica is a movement done in [this](/media/2004-07-19_DJ_11.jpg) sport, popular in Mexico"
	if challenge.Prompt != expected_prompt {
		t.Errorf("incorrect prompt, got\n%s\n%s\nexpected", challenge.Prompt, expected_prompt)
	}
	if challenge.Media[0].MediaType != MediaImageJPG {
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

	challenge := new(JArchiveChallenge)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("failed to parse (html_reader)\n%s", err)
	}
	td := nextDescendantWithClass(doc, "td", "clue")
	challenge.parseChallenge(td)

	if challenge.Round != ROUND_UNKNOWN {
		t.Errorf("incorrect round (%s)", challenge.Round)
	}
	if challenge.Category != "" {
		t.Errorf("parseChallenge should not set Category property")
	}
	if challenge.Commentary != "" {
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
	if challenge.Prompt != expected_text {
		t.Errorf("incorrect prompt '%s' !=\n'%s'", challenge.Prompt, expected_text)
	}
	if len(challenge.Media) != 1 || challenge.Media[0].MediaType != MediaAudioMP3 {
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

	challenge := new(JArchiveChallenge)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("failed to parse (html_reader)\n%s", err)
	}
	td := nextDescendantWithClass(doc, "td", "clue")
	challenge.parseChallenge(td)

	expected_prompt := "( [Cheryl of the Clue Crew drags a mustache back onto Alex.](/media/2004-07-19_DJ_11.mp4) ) [Introduced](/media/2004-07-19_DJ_11.jpg) in 1990, it's the [program](/media/2004-07-19_DJ_11a.jpg) I'm using to create an [unusual image](/media/2004-07-19_DJ_11b.jpg)"
	if challenge.Prompt != expected_prompt {
		t.Errorf("incorrect prompt, got\n%s\n%s\nexpected", challenge.Prompt, expected_prompt)
	}
	expected_media := []Media{
		{MediaVideoMP4, "/media/2004-07-19_DJ_11.mp4"},
		{MediaImageJPG, "/media/2004-07-19_DJ_11.jpg"},
		{MediaImageJPG, "/media/2004-07-19_DJ_11a.jpg"},
		{MediaImageJPG, "/media/2004-07-19_DJ_11b.jpg"},
	}
	for i, expected := range expected_media {
		if challenge.Media[i].MediaType != expected.MediaType {
			t.Errorf("incorrect media type %s\n%s", challenge.Media[i].MediaType, expected.MediaType)
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

	final_challenge := new(JArchiveFinalChallenge)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatalf("failed to parse (html_reader)\n%s", err)
	}
	final_round := nextDescendantWithClass(doc, "table", "final_round")
	final_challenge.parseChallenge(final_round)

	if final_challenge.Round != ROUND_FINAL_JEOPARDY {
		t.Errorf("incorrect round (%s)", final_challenge.Round)
	}
	if final_challenge.Category != "WORLD AFFAIRS" {
		t.Errorf("incorrect category '%s'", final_challenge.Category)
	}
	if final_challenge.Commentary != "comment this" {
		t.Errorf("incorrect comment '%s'", final_challenge.Commentary)
	}

	if final_challenge.Prompt != "In 1963 these 3 nations signed the Nuclear Test Ban Treaty" {
		t.Errorf("incorrect prompt '%s'", final_challenge.Prompt)
	}
	if final_challenge.Correct != "the U.S., U.S.S.R. & the U.K." {
		t.Errorf("incorrect response '%s'", final_challenge.Correct)
	}
}
