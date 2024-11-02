package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

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
    <td id="clue_FJ_r" class="clue_text" style="display:none;"><table><tr><td class="right">Ken</td><td rowspan="2" valign="top">What are the U.S., USSR and UK?</td></tr><tr><td>$2,500</td></tr><tr><td class="right">Jim</td><td rowspan="2" valign="top">What are the U.S., U.S.S.R and the U.K.?</td></tr><tr><td>$5,000</td></tr><tr><td class="wrong">Judy</td><td rowspan="2" valign="top">What were the USSR, USA, and China?</td></tr><tr><td>$1,401</td></tr></table><em class="correct_response">the U.S., U.S.S.R. & the U.K.</em></td>
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
