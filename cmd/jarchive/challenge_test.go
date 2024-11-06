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

	challenge := new(JArchiveChallenge)
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatal("failed to parse (html_reader)\n", err)
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
	html_raw := `<td class="clue">
    </td>`

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
		challenge.Prompt != "" || challenge.Correct != "" || len(challenge.Accept) != 0 {
		t.Error("incorrect (non-zero) result from empty challenge")
	}
	if !challenge.IsEmpty() {
		t.Error("incorrect (non-zero) result from empty challenge")
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
