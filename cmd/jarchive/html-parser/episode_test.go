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
// github:kevindamm/q-party/cmd/jarchive/html/episode_test.go

package html

import (
	"strings"
	"testing"

	qparty "github.com/kevindamm/q-party"
	"golang.org/x/net/html"
)

func TestJArchiveEpisode_parseContent(t *testing.T) {
	html_raw := `<div class="content">
            <div id="game_title"><h1>Show #4857 - Tuesday, October 25, 2005</h1></div>
            <div id="game_comments">(Kelly: We're here in wonderful <a href="http://www.j-archive.com/media/2005-10-25_ComingUp.jpg" target="_blank">Copenhagen</a> celebrating 200 years of Hans Christian Andersen.  Join us, next on <i>Jeopardy!</i>)</div>
            <div id="contestants">
            <h2>Contestants</h2>
              <table id="contestants_table"><tr>
                <td align="left" valign="middle" style="width:150px">
                  <a href="showgame.php?game_id=578">[&lt;&lt; previous game]</a>
                </td>
                <td>
                  <p class="contestants"><a href="showplayer.php?player_id=1200" rel="external">Diane Mettam</a>, a pastor from Independence, California</p>
                  <p class="contestants"><a href="showplayer.php?player_id=1201" rel="external">Chris Jones</a>, a loan consultant from Chalfont, Pennsylvania</p>
                  <p class="contestants"><a href="showplayer.php?player_id=1199" rel="external">John Kelly</a>, an attorney from Atlanta, Georgia (whose 1-day cash winnings total $21,601)</p>
                </td>
                <td align="right" valign="middle" style="width:150px">
                  <a href="showgame.php?game_id=581">[next game &gt;&gt;]</a>
                </td>
              </tr></table>
            </div>
            <div id="jeopardy_round">
              <h2>Jeopardy! Round</h2>

<table class="round">
  <tr>
    <td class="category">
      
<table>
  <tr><td class="category_name">HANS CHRISTIAN ANDERSEN</td></tr>
  <tr><td class="category_comments">(Alex: We're celebrating his bicentennial this year.)</td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">ABBA-SOLUTELY FABULOUS</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">FIRST NAME'S THE SAME</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">THE BRADY BRUNCH</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">SPEECH!  SPEECH!</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">OH, BE "SILENT"!</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
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
    <td id="clue_J_1_1" class="clue_text">(<a href="http://www.j-archive.com/media/2005-10-25_J_12.jpg" target="_blank">Kelly of the Clue Crew reads from the statue of The Little Mermaid in Copenhagen, Denmark.</a>)  In the original story, the prince never knew how much the Little Mermaid loved him because the Sea Witch had taken this away</td>
    <td id="clue_J_1_1_r" class="clue_text" style="display:none;"><em class="correct_response">her voice</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_2_1', 'clue_J_2_1_r', 'clue_J_2_1_stuck')" onmouseout="toggle('clue_J_2_1_r', 'clue_J_2_1', 'clue_J_2_1_stuck')" onclick="togglestick('clue_J_2_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_2_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32576" title="Suggest a correction for this clue" rel="nofollow">16</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_2_1" class="clue_text">Made up of Frida Lyngstad, Agnetha Faltskog, Bjorn Ulvaeus &amp; Benny Andersson, ABBA hails from this country</td>
    <td id="clue_J_2_1_r" class="clue_text" style="display:none;"><em class="correct_response">Sweden</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_3_1', 'clue_J_3_1_r', 'clue_J_3_1_stuck')" onmouseout="toggle('clue_J_3_1_r', 'clue_J_3_1', 'clue_J_3_1_stuck')" onclick="togglestick('clue_J_3_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_3_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32562" title="Suggest a correction for this clue" rel="nofollow">3</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_3_1" class="clue_text">Blackmun,<br />Houdini,<br />Cohn</td>
    <td id="clue_J_3_1_r" class="clue_text" style="display:none;"><em class="correct_response">Harry</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_4_1', 'clue_J_4_1_r', 'clue_J_4_1_stuck')" onmouseout="toggle('clue_J_4_1_r', 'clue_J_4_1', 'clue_J_4_1_stuck')" onclick="togglestick('clue_J_4_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_4_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32567" title="Suggest a correction for this clue" rel="nofollow">11</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_4_1" class="clue_text">"Oh, my nose!", Marcia exclaimed, after bumping into the oven door &amp; causing this light, fluffy egg dish to fall flat</td>
    <td id="clue_J_4_1_r" class="clue_text" style="display:none;">(John: What is an omelette?)<br />(Chris: What is a quiche?)<br /><br /><em class="correct_response">a soufflé</em><br /><br /><table width="100%"><tr><td class="wrong">John</td><td class="wrong">Chris</td><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_5_1', 'clue_J_5_1_r', 'clue_J_5_1_stuck')" onmouseout="toggle('clue_J_5_1_r', 'clue_J_5_1', 'clue_J_5_1_stuck')" onclick="togglestick('clue_J_5_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_5_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32584" title="Suggest a correction for this clue" rel="nofollow">22</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_5_1" class="clue_text">In a 1991 speech this ex-president endorsed the Brady gun-control bill</td>
    <td id="clue_J_5_1_r" class="clue_text" style="display:none;"><em class="correct_response">Reagan</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_6_1', 'clue_J_6_1_r', 'clue_J_6_1_stuck')" onmouseout="toggle('clue_J_6_1_r', 'clue_J_6_1', 'clue_J_6_1_stuck')" onclick="togglestick('clue_J_6_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_6_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32569" title="Suggest a correction for this clue" rel="nofollow">1</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_6_1" class="clue_text">In German it begins, "Stille Nacht! Heilige Nacht!"</td>
    <td id="clue_J_6_1_r" class="clue_text" style="display:none;"><em class="correct_response">"Silent Night"</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_1_2', 'clue_J_1_2_r', 'clue_J_1_2_stuck')" onmouseout="toggle('clue_J_1_2_r', 'clue_J_1_2', 'clue_J_1_2_stuck')" onclick="togglestick('clue_J_1_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_1_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32560" title="Suggest a correction for this clue" rel="nofollow">13</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_1_2" class="clue_text">H.C. Andersen is often compared to this title character of his, who felt unloved until he made a big discovery</td>
    <td id="clue_J_1_2_r" class="clue_text" style="display:none;"><em class="correct_response">the Ugly Duckling</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_2_2', 'clue_J_2_2_r', 'clue_J_2_2_stuck')" onmouseout="toggle('clue_J_2_2_r', 'clue_J_2_2', 'clue_J_2_2_stuck')" onclick="togglestick('clue_J_2_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_2_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32577" title="Suggest a correction for this clue" rel="nofollow">25</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_2_2" class="clue_text">"See that girl, watch that scene, dig in" this ABBA title teen</td>
    <td id="clue_J_2_2_r" class="clue_text" style="display:none;"><em class="correct_response">the Dancing Queen</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_3_2', 'clue_J_3_2_r', 'clue_J_3_2_stuck')" onmouseout="toggle('clue_J_3_2_r', 'clue_J_3_2', 'clue_J_3_2_stuck')" onclick="togglestick('clue_J_3_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_3_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32563" title="Suggest a correction for this clue" rel="nofollow">4</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_3_2" class="clue_text">Reynolds,<br />Allen,<br />Harry</td>
    <td id="clue_J_3_2_r" class="clue_text" style="display:none;"><em class="correct_response">Debbie</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_4_2', 'clue_J_4_2_r', 'clue_J_4_2_stuck')" onmouseout="toggle('clue_J_4_2_r', 'clue_J_4_2', 'clue_J_4_2_stuck')" onclick="togglestick('clue_J_4_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_4_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32568" title="Suggest a correction for this clue" rel="nofollow">15</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_4_2" class="clue_text">A hungry Peter keeps mentioning <a href="http://www.j-archive.com/media/2005-10-25_J_15.mp3">these &amp; applesauce, sweetheart</a> (like in one famous episode)</td>
    <td id="clue_J_4_2_r" class="clue_text" style="display:none;">[Alex imitates Humphrey Bogart for "these &amp; applesauce, sweetheart."]<br /><br /><em class="correct_response">pork chops</em><br /><br /><table width="100%"><tr><td class="wrong">Triple Stumper</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_5_2', 'clue_J_5_2_r', 'clue_J_5_2_stuck')" onmouseout="toggle('clue_J_5_2_r', 'clue_J_5_2', 'clue_J_5_2_stuck')" onclick="togglestick('clue_J_5_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_5_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32585" title="Suggest a correction for this clue" rel="nofollow">23</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_5_2" class="clue_text">He gave the plea for mercy speech at Leopold &amp; Loeb's trial on August 22, 1924</td>
    <td id="clue_J_5_2_r" class="clue_text" style="display:none;"><em class="correct_response">Clarence Darrow</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_6_2', 'clue_J_6_2_r', 'clue_J_6_2_stuck')" onmouseout="toggle('clue_J_6_2_r', 'clue_J_6_2', 'clue_J_6_2_stuck')" onclick="togglestick('clue_J_6_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_6_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32570" title="Suggest a correction for this clue" rel="nofollow">2</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_6_2" class="clue_text">Sale in which bids are submitted in sealed envelopes</td>
    <td id="clue_J_6_2_r" class="clue_text" style="display:none;"><em class="correct_response">a silent auction</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_1_3', 'clue_J_1_3_r', 'clue_J_1_3_stuck')" onmouseout="toggle('clue_J_1_3_r', 'clue_J_1_3', 'clue_J_1_3_stuck')" onclick="togglestick('clue_J_1_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_1_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32561" title="Suggest a correction for this clue" rel="nofollow">14</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_1_3" class="clue_text">(<a href="http://www.j-archive.com/media/2005-10-25_J_14.jpg" target="_blank">Jon of the Clue Crew reads from Nyhavn, Denmark.</a>)  Hans Christian Andersen lived at several addresses here in Nyhavn, a nautical neighborhood whose name means "new" this</td>
    <td id="clue_J_1_3_r" class="clue_text" style="display:none;"><em class="correct_response">harbor</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_2_3', 'clue_J_2_3_r', 'clue_J_2_3_stuck')" onmouseout="toggle('clue_J_2_3_r', 'clue_J_2_3', 'clue_J_2_3_stuck')" onclick="togglestick('clue_J_2_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_2_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32578" title="Suggest a correction for this clue" rel="nofollow">26</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_2_3" class="clue_text">It was the end for Napoleon but this song was the beginning for ABBA, marking its first foray into the Top 40</td>
    <td id="clue_J_2_3_r" class="clue_text" style="display:none;"><em class="correct_response">"Waterloo"</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_3_3', 'clue_J_3_3_r', 'clue_J_3_3_stuck')" onmouseout="toggle('clue_J_3_3_r', 'clue_J_3_3', 'clue_J_3_3_stuck')" onclick="togglestick('clue_J_3_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_3_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32564" title="Suggest a correction for this clue" rel="nofollow">5</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_3_3" class="clue_text">Bergen,<br />Degas,<br />Winter</td>
    <td id="clue_J_3_3_r" class="clue_text" style="display:none;"><em class="correct_response">Edgar</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_4_3', 'clue_J_4_3_r', 'clue_J_4_3_stuck')" onmouseout="toggle('clue_J_4_3_r', 'clue_J_4_3', 'clue_J_4_3_stuck')" onclick="togglestick('clue_J_4_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_4_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32581" title="Suggest a correction for this clue" rel="nofollow">19</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_4_3" class="clue_text">Jan impressed the family by making the blueberry-filled type of these Jewish pancakes</td>
    <td id="clue_J_4_3_r" class="clue_text" style="display:none;"><em class="correct_response">blintzes</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_5_3', 'clue_J_5_3_r', 'clue_J_5_3_stuck')" onmouseout="toggle('clue_J_5_3_r', 'clue_J_5_3', 'clue_J_5_3_stuck')" onclick="togglestick('clue_J_5_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_5_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value_daily_double">DD: $2,800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32586" title="Suggest a correction for this clue" rel="nofollow">24</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_5_3" class="clue_text">On July 4, 1939 he gave a famous farewell speech in the Bronx</td>
    <td id="clue_J_5_3_r" class="clue_text" style="display:none;">(Alex: You loved the baseball category yesterday, and oddly enough, it came up for you again here!)<br /><br /><em class="correct_response">Lou Gehrig</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_6_3', 'clue_J_6_3_r', 'clue_J_6_3_stuck')" onmouseout="toggle('clue_J_6_3_r', 'clue_J_6_3', 'clue_J_6_3_stuck')" onclick="togglestick('clue_J_6_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_6_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32571" title="Suggest a correction for this clue" rel="nofollow">7</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_6_3" class="clue_text">Richard Nixon identified this large group that chooses not to express its views publicly</td>
    <td id="clue_J_6_3_r" class="clue_text" style="display:none;"><em class="correct_response">the silent majority</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_1_4', 'clue_J_1_4_r', 'clue_J_1_4_stuck')" onmouseout="toggle('clue_J_1_4_r', 'clue_J_1_4', 'clue_J_1_4_stuck')" onclick="togglestick('clue_J_1_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_1_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32574" title="Suggest a correction for this clue" rel="nofollow">17</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_1_4" class="clue_text">Andersen's early fairy tales include this one that inspired the musical "Once Upon a Mattress"</td>
    <td id="clue_J_1_4_r" class="clue_text" style="display:none;"><em class="correct_response">"The Princess and the Pea"</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_2_4', 'clue_J_2_4_r', 'clue_J_2_4_stuck')" onmouseout="toggle('clue_J_2_4_r', 'clue_J_2_4', 'clue_J_2_4_stuck')" onclick="togglestick('clue_J_2_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_2_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32579" title="Suggest a correction for this clue" rel="nofollow">27</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_2_4" class="clue_text">In the spring of 1999 this show featuring the songs of ABBA premiered in London</td>
    <td id="clue_J_2_4_r" class="clue_text" style="display:none;"><em class="correct_response"><i>Mamma Mia!</i></em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_3_4', 'clue_J_3_4_r', 'clue_J_3_4_stuck')" onmouseout="toggle('clue_J_3_4_r', 'clue_J_3_4', 'clue_J_3_4_stuck')" onclick="togglestick('clue_J_3_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_3_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32565" title="Suggest a correction for this clue" rel="nofollow">6</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_3_4" class="clue_text">Halley,<br />Muskie,<br />Burke</td>
    <td id="clue_J_3_4_r" class="clue_text" style="display:none;"><em class="correct_response">Edmund</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_4_4', 'clue_J_4_4_r', 'clue_J_4_4_stuck')" onmouseout="toggle('clue_J_4_4_r', 'clue_J_4_4', 'clue_J_4_4_stuck')" onclick="togglestick('clue_J_4_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_4_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32582" title="Suggest a correction for this clue" rel="nofollow">20</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_4_4" class="clue_text">When no one was looking, Cindy &amp; Bobby tried to sneak a sip of this orange juice &amp; champagne cocktail</td>
    <td id="clue_J_4_4_r" class="clue_text" style="display:none;"><em class="correct_response">a mimosa</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_5_4', 'clue_J_5_4_r', 'clue_J_5_4_stuck')" onmouseout="toggle('clue_J_5_4_r', 'clue_J_5_4', 'clue_J_5_4_stuck')" onclick="togglestick('clue_J_5_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_5_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32587" title="Suggest a correction for this clue" rel="nofollow">29</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_5_4" class="clue_text">In a 1922 speech he declared, "Nonviolence is the first article of my faith"</td>
    <td id="clue_J_5_4_r" class="clue_text" style="display:none;"><em class="correct_response">(Mahatma) Gandhi</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_6_4', 'clue_J_6_4_r', 'clue_J_6_4_stuck')" onmouseout="toggle('clue_J_6_4_r', 'clue_J_6_4', 'clue_J_6_4_stuck')" onclick="togglestick('clue_J_6_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_6_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32572" title="Suggest a correction for this clue" rel="nofollow">8</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_6_4" class="clue_text">Ironically, Marcel Marceau has the only line of dialogue in this film</td>
    <td id="clue_J_6_4_r" class="clue_text" style="display:none;"><em class="correct_response"><i>Silent Movie</i></em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_1_5', 'clue_J_1_5_r', 'clue_J_1_5_stuck')" onmouseout="toggle('clue_J_1_5_r', 'clue_J_1_5', 'clue_J_1_5_stuck')" onclick="togglestick('clue_J_1_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_1_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32575" title="Suggest a correction for this clue" rel="nofollow">18</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_1_5" class="clue_text">(<a href="http://www.j-archive.com/media/2005-10-25_J_18.jpg" target="_blank">Jimmy of the Clue Crew reports from Tivoli, Copenhagen, Denmark.</a>)  A <a href="http://www.j-archive.com/media/2005-10-25_J_18a.jpg" target="_blank">visit</a> to Tivoli in its very first season in 1843 inspired Andersen to write this, which he called his "Chinese fairy tale"</td>
    <td id="clue_J_1_5_r" class="clue_text" style="display:none;">(Chris: What is "The Emperor's New Clothes"?)<br /><br /><em class="correct_response">"The Nightingale"</em><br /><br /><table width="100%"><tr><td class="wrong">Chris</td><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_2_5', 'clue_J_2_5_r', 'clue_J_2_5_stuck')" onmouseout="toggle('clue_J_2_5_r', 'clue_J_2_5', 'clue_J_2_5_stuck')" onclick="togglestick('clue_J_2_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_2_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32580" title="Suggest a correction for this clue" rel="nofollow">28</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_2_5" class="clue_text">ABBA hit the Top 40 for the third time in 1975 with this 3-letter palindrome</td>
    <td id="clue_J_2_5_r" class="clue_text" style="display:none;"><em class="correct_response">"SOS"</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_3_5', 'clue_J_3_5_r', 'clue_J_3_5_stuck')" onmouseout="toggle('clue_J_3_5_r', 'clue_J_3_5', 'clue_J_3_5_stuck')" onclick="togglestick('clue_J_3_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_3_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32566" title="Suggest a correction for this clue" rel="nofollow">10</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_3_5" class="clue_text">Klesko,<br />Longwell,<br />Gosling</td>
    <td id="clue_J_3_5_r" class="clue_text" style="display:none;"><em class="correct_response">Ryan</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_4_5', 'clue_J_4_5_r', 'clue_J_4_5_stuck')" onmouseout="toggle('clue_J_4_5_r', 'clue_J_4_5', 'clue_J_4_5_stuck')" onclick="togglestick('clue_J_4_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_4_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32583" title="Suggest a correction for this clue" rel="nofollow">21</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_4_5" class="clue_text">Alice chopped up lots of vegetables for this Italian omelet that resembles a large pancake</td>
    <td id="clue_J_4_5_r" class="clue_text" style="display:none;"><em class="correct_response">a frittata</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_5_5', 'clue_J_5_5_r', 'clue_J_5_5_stuck')" onmouseout="toggle('clue_J_5_5_r', 'clue_J_5_5', 'clue_J_5_5_stuck')" onclick="togglestick('clue_J_5_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_5_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32588" title="Suggest a correction for this clue" rel="nofollow">30</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_5_5" class="clue_text">In 1848, in this New York town, Elizabeth Cady Stanton said, "The right is ours. Have it, we must. Use it, we will"</td>
    <td id="clue_J_5_5_r" class="clue_text" style="display:none;">(John: What is Poughkeepsie?)<br /><br /><em class="correct_response">Seneca Falls</em><br /><br /><table width="100%"><tr><td class="wrong">John</td></tr></table><table width="100%"><tr><td class="wrong">Triple Stumper</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_J_6_5', 'clue_J_6_5_r', 'clue_J_6_5_stuck')" onmouseout="toggle('clue_J_6_5_r', 'clue_J_6_5', 'clue_J_6_5_stuck')" onclick="togglestick('clue_J_6_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_J_6_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32573" title="Suggest a correction for this clue" rel="nofollow">9</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_J_6_5" class="clue_text">It's a small container with a hinged lid used to collect crumbs or ashes from the dinner table</td>
    <td id="clue_J_6_5_r" class="clue_text" style="display:none;"><em class="correct_response">a silent butler</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
</table>
              <h3>Scores at the first commercial break (after clue 15):</h3>

              <table>
                <tr>
                  <td class="score_player_nickname">John</td>
                  <td class="score_player_nickname">Chris</td>
                  <td class="score_player_nickname">Diane</td>
                </tr>
                <tr>
                  <td class="score_positive">$2,400</td>
                  <td class="score_positive">$600</td>
                  <td class="score_positive">$4,000</td>
                </tr>
              </table>

              <h3>Scores at the end of the Jeopardy! Round:</h3>

              <table>
                <tr>
                  <td class="score_player_nickname">John</td>
                  <td class="score_player_nickname">Chris</td>
                  <td class="score_player_nickname">Diane</td>
                </tr>
                <tr>
                  <td class="score_positive">$6,000</td>
                  <td class="score_positive">$2,200</td>
                  <td class="score_positive">$8,200</td>
                </tr>
              </table>
            </div>
            <div id="double_jeopardy_round">
              <h2>Double Jeopardy! Round</h2>

<table class="round">
  <tr>
    <td class="category">
      
<table>
  <tr><td class="category_name">WORLD CAPITALS</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">MOVIES TO THE MAX</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">ALMOST ASSASSINATED</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">AMERICAN WOMEN</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">FURRED, FEATHERED, FINNED</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
    <td class="category">
      
<table>
  <tr><td class="category_name">THE "O.C."</td></tr>
  <tr><td class="category_comments"></td></tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_1_1', 'clue_DJ_1_1_r', 'clue_DJ_1_1_stuck')" onmouseout="toggle('clue_DJ_1_1_r', 'clue_DJ_1_1', 'clue_DJ_1_1_stuck')" onclick="togglestick('clue_DJ_1_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_1_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32751" title="Suggest a correction for this clue" rel="nofollow">3</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_1_1" class="clue_text">By population, it's the largest capital in the Western Hemisphere</td>
    <td id="clue_DJ_1_1_r" class="clue_text" style="display:none;"><em class="correct_response">Mexico City</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_2_1', 'clue_DJ_2_1_r', 'clue_DJ_2_1_stuck')" onmouseout="toggle('clue_DJ_2_1_r', 'clue_DJ_2_1', 'clue_DJ_2_1_stuck')" onclick="togglestick('clue_DJ_2_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_2_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32756" title="Suggest a correction for this clue" rel="nofollow">2</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_2_1" class="clue_text">Max Detweiler is the impresario character who enters a family in music festival in this beloved 1965 film</td>
    <td id="clue_DJ_2_1_r" class="clue_text" style="display:none;"><em class="correct_response"><i>The Sound of Music</i></em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_3_1', 'clue_DJ_3_1_r', 'clue_DJ_3_1_stuck')" onmouseout="toggle('clue_DJ_3_1_r', 'clue_DJ_3_1', 'clue_DJ_3_1_stuck')" onclick="togglestick('clue_DJ_3_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_3_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32768" title="Suggest a correction for this clue" rel="nofollow">22</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_3_1" class="clue_text">February 13, 1933 after making a speech in Miami</td>
    <td id="clue_DJ_3_1_r" class="clue_text" style="display:none;"><em class="correct_response">FDR</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_4_1', 'clue_DJ_4_1_r', 'clue_DJ_4_1_stuck')" onmouseout="toggle('clue_DJ_4_1_r', 'clue_DJ_4_1', 'clue_DJ_4_1_stuck')" onclick="togglestick('clue_DJ_4_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_4_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32761" title="Suggest a correction for this clue" rel="nofollow">13</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_4_1" class="clue_text">"Brooklyn Bridge" was one of the last of her NYC paintings before she moved permanently to New Mexico</td>
    <td id="clue_DJ_4_1_r" class="clue_text" style="display:none;"><em class="correct_response">(Georgia) O'Keeffe</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_5_1', 'clue_DJ_5_1_r', 'clue_DJ_5_1_stuck')" onmouseout="toggle('clue_DJ_5_1_r', 'clue_DJ_5_1', 'clue_DJ_5_1_stuck')" onclick="togglestick('clue_DJ_5_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_5_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32762" title="Suggest a correction for this clue" rel="nofollow">17</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_5_1" class="clue_text">The brown type of this pouched bird plunges from the air to fish; the white one scoops up fish as it swims</td>
    <td id="clue_DJ_5_1_r" class="clue_text" style="display:none;"><em class="correct_response">a pelican</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_6_1', 'clue_DJ_6_1_r', 'clue_DJ_6_1_stuck')" onmouseout="toggle('clue_DJ_6_1_r', 'clue_DJ_6_1', 'clue_DJ_6_1_stuck')" onclick="togglestick('clue_DJ_6_1_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_6_1_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$400</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32763" title="Suggest a correction for this clue" rel="nofollow">7</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_6_1" class="clue_text">These crispy snacks were created for a certain mollusk soup</td>
    <td id="clue_DJ_6_1_r" class="clue_text" style="display:none;"><em class="correct_response">oyster crackers</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_1_2', 'clue_DJ_1_2_r', 'clue_DJ_1_2_stuck')" onmouseout="toggle('clue_DJ_1_2_r', 'clue_DJ_1_2', 'clue_DJ_1_2_stuck')" onclick="togglestick('clue_DJ_1_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_1_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32752" title="Suggest a correction for this clue" rel="nofollow">9</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_1_2" class="clue_text">This capital is located on Luzon Island</td>
    <td id="clue_DJ_1_2_r" class="clue_text" style="display:none;"><em class="correct_response">Manila</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_2_2', 'clue_DJ_2_2_r', 'clue_DJ_2_2_stuck')" onmouseout="toggle('clue_DJ_2_2_r', 'clue_DJ_2_2', 'clue_DJ_2_2_stuck')" onclick="togglestick('clue_DJ_2_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_2_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32757" title="Suggest a correction for this clue" rel="nofollow">1</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_2_2" class="clue_text">Laurence Olivier starred as Maxim de Winter in this film, but no one played the title role</td>
    <td id="clue_DJ_2_2_r" class="clue_text" style="display:none;"><em class="correct_response"><i>Rebecca</i></em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_3_2', 'clue_DJ_3_2_r', 'clue_DJ_3_2_stuck')" onmouseout="toggle('clue_DJ_3_2_r', 'clue_DJ_3_2', 'clue_DJ_3_2_stuck')" onclick="togglestick('clue_DJ_3_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_3_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32769" title="Suggest a correction for this clue" rel="nofollow">23</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_3_2" class="clue_text">May 12, 1982 on a visit to the famous shrine in Fatima, Portugal</td>
    <td id="clue_DJ_3_2_r" class="clue_text" style="display:none;"><em class="correct_response">Pope John Paul II</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_4_2', 'clue_DJ_4_2_r', 'clue_DJ_4_2_stuck')" onmouseout="toggle('clue_DJ_4_2_r', 'clue_DJ_4_2', 'clue_DJ_4_2_stuck')" onclick="togglestick('clue_DJ_4_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_4_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32773" title="Suggest a correction for this clue" rel="nofollow">27</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_4_2" class="clue_text">6 years after Chuck Yeager, Jacqueline Cochran became the first woman to do this</td>
    <td id="clue_DJ_4_2_r" class="clue_text" style="display:none;"><em class="correct_response">break the sound barrier</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_5_2', 'clue_DJ_5_2_r', 'clue_DJ_5_2_stuck')" onmouseout="toggle('clue_DJ_5_2_r', 'clue_DJ_5_2', 'clue_DJ_5_2_stuck')" onclick="togglestick('clue_DJ_5_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_5_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32777" title="Suggest a correction for this clue" rel="nofollow">18</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_5_2" class="clue_text">The mulgara of Queensland is among the carnivorous members of this order</td>
    <td id="clue_DJ_5_2_r" class="clue_text" style="display:none;">(Alex: It's a [*]. That's a tough one.)<br /><br /><em class="correct_response">marsupial</em><br /><br /><table width="100%"><tr><td class="wrong">Triple Stumper</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_6_2', 'clue_DJ_6_2_r', 'clue_DJ_6_2_stuck')" onmouseout="toggle('clue_DJ_6_2_r', 'clue_DJ_6_2', 'clue_DJ_6_2_stuck')" onclick="togglestick('clue_DJ_6_2_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_6_2_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32764" title="Suggest a correction for this clue" rel="nofollow">8</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_6_2" class="clue_text">This type of fabric often used for shirts is named for a city northwest of London</td>
    <td id="clue_DJ_6_2_r" class="clue_text" style="display:none;"><em class="correct_response">Oxford cloth</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_1_3', 'clue_DJ_1_3_r', 'clue_DJ_1_3_stuck')" onmouseout="toggle('clue_DJ_1_3_r', 'clue_DJ_1_3', 'clue_DJ_1_3_stuck')" onclick="togglestick('clue_DJ_1_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_1_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32753" title="Suggest a correction for this clue" rel="nofollow">14</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_1_3" class="clue_text">The famous Gold Museum in this Colombian capital houses a large collection of pre-Columbian gold objects</td>
    <td id="clue_DJ_1_3_r" class="clue_text" style="display:none;"><em class="correct_response">Bogota</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_2_3', 'clue_DJ_2_3_r', 'clue_DJ_2_3_stuck')" onmouseout="toggle('clue_DJ_2_3_r', 'clue_DJ_2_3', 'clue_DJ_2_3_stuck')" onclick="togglestick('clue_DJ_2_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_2_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32758" title="Suggest a correction for this clue" rel="nofollow">4</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_2_3" class="clue_text">This former TV Hillbilly directed the 1976 film "Ode to Billie Joe"</td>
    <td id="clue_DJ_2_3_r" class="clue_text" style="display:none;"><em class="correct_response">Max Baer (Jr.)</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_3_3', 'clue_DJ_3_3_r', 'clue_DJ_3_3_stuck')" onmouseout="toggle('clue_DJ_3_3_r', 'clue_DJ_3_3', 'clue_DJ_3_3_stuck')" onclick="togglestick('clue_DJ_3_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_3_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32770" title="Suggest a correction for this clue" rel="nofollow">24</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_3_3" class="clue_text">September 5, 1975 while greeting a crowd in Sacramento, California</td>
    <td id="clue_DJ_3_3_r" class="clue_text" style="display:none;"><em class="correct_response">Ford</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_4_3', 'clue_DJ_4_3_r', 'clue_DJ_4_3_stuck')" onmouseout="toggle('clue_DJ_4_3_r', 'clue_DJ_4_3', 'clue_DJ_4_3_stuck')" onclick="togglestick('clue_DJ_4_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_4_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32774" title="Suggest a correction for this clue" rel="nofollow">30</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_4_3" class="clue_text">(<a href="http://www.j-archive.com/media/2005-10-25_DJ_30.jpg" target="_blank">Hi, I'm Bob Woodward,</a>)  One of the most influential women of the 20th century, she was publisher of the Washington Post from 1969 to 1979 &amp; CEO until 1991</td>
    <td id="clue_DJ_4_3_r" class="clue_text" style="display:none;">(Alex: [*], Bob's boss.)<br /><br /><em class="correct_response">(Katharine) Graham</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_5_3', 'clue_DJ_5_3_r', 'clue_DJ_5_3_stuck')" onmouseout="toggle('clue_DJ_5_3_r', 'clue_DJ_5_3', 'clue_DJ_5_3_stuck')" onclick="togglestick('clue_DJ_5_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_5_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32778" title="Suggest a correction for this clue" rel="nofollow">19</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_5_3" class="clue_text">The dogfish is a small type of this fish, &amp; one dogfish is also called the "Greenland" one</td>
    <td id="clue_DJ_5_3_r" class="clue_text" style="display:none;"><em class="correct_response">a shark</em><br /><br /><table width="100%"><tr><td class="wrong">Triple Stumper</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_6_3', 'clue_DJ_6_3_r', 'clue_DJ_6_3_stuck')" onmouseout="toggle('clue_DJ_6_3_r', 'clue_DJ_6_3', 'clue_DJ_6_3_stuck')" onclick="togglestick('clue_DJ_6_3_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_6_3_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1200</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32765" title="Suggest a correction for this clue" rel="nofollow">10</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_6_3" class="clue_text">Me, him &amp; her are examples of this form of a pronoun</td>
    <td id="clue_DJ_6_3_r" class="clue_text" style="display:none;"><em class="correct_response">objective case</em><br /><br /><table width="100%"><tr><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_1_4', 'clue_DJ_1_4_r', 'clue_DJ_1_4_stuck')" onmouseout="toggle('clue_DJ_1_4_r', 'clue_DJ_1_4', 'clue_DJ_1_4_stuck')" onclick="togglestick('clue_DJ_1_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_1_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32754" title="Suggest a correction for this clue" rel="nofollow">15</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_1_4" class="clue_text">The world's largest govt. building after the Pentagon is this city's Parliament Palace, built by Ceausescu</td>
    <td id="clue_DJ_1_4_r" class="clue_text" style="display:none;"><em class="correct_response">Bucharest</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_2_4', 'clue_DJ_2_4_r', 'clue_DJ_2_4_stuck')" onmouseout="toggle('clue_DJ_2_4_r', 'clue_DJ_2_4', 'clue_DJ_2_4_stuck')" onclick="togglestick('clue_DJ_2_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_2_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32759" title="Suggest a correction for this clue" rel="nofollow">5</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_2_4" class="clue_text">He directed a 1984 documentary about Marlene Dietrich, his co-star in "Judgment at Nuremberg"</td>
    <td id="clue_DJ_2_4_r" class="clue_text" style="display:none;">(Chris: Who is Max von Sydow?)<br />...<br />(Alex: He appeared in the film and it's Max, [*].)<br /><br /><em class="correct_response">Maximilian Schell</em><br /><br /><table width="100%"><tr><td class="wrong">Chris</td></tr></table><table width="100%"><tr><td class="wrong">Triple Stumper</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_3_4', 'clue_DJ_3_4_r', 'clue_DJ_3_4_stuck')" onmouseout="toggle('clue_DJ_3_4_r', 'clue_DJ_3_4', 'clue_DJ_3_4_stuck')" onclick="togglestick('clue_DJ_3_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_3_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value_daily_double">DD: $2,800</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32771" title="Suggest a correction for this clue" rel="nofollow">25</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_3_4" class="clue_text">August 22, 1962 when his limousine was attacked near Paris</td>
    <td id="clue_DJ_3_4_r" class="clue_text" style="display:none;">(Chris: Who is Malcolm X?)<br />(Alex: Oh, no.  Who is [*]? [*].  They tried to get Big Chuck.  Select.)<br /><br /><em class="correct_response">Charles de Gaulle</em><br /><br /><table width="100%"><tr><td class="wrong">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_4_4', 'clue_DJ_4_4_r', 'clue_DJ_4_4_stuck')" onmouseout="toggle('clue_DJ_4_4_r', 'clue_DJ_4_4', 'clue_DJ_4_4_stuck')" onclick="togglestick('clue_DJ_4_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_4_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32775" title="Suggest a correction for this clue" rel="nofollow">29</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_4_4" class="clue_text">Gwendolyn Brooks, the first African American to win a Pulitzer Prize, won for this category in 1950</td>
    <td id="clue_DJ_4_4_r" class="clue_text" style="display:none;"><em class="correct_response">Poetry</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_5_4', 'clue_DJ_5_4_r', 'clue_DJ_5_4_stuck')" onmouseout="toggle('clue_DJ_5_4_r', 'clue_DJ_5_4', 'clue_DJ_5_4_stuck')" onclick="togglestick('clue_DJ_5_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_5_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value_daily_double">DD: $2,500</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32779" title="Suggest a correction for this clue" rel="nofollow">20</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_5_4" class="clue_text">It can be a 1/6-inch printer's unit, or (spelled differently) a 7- or 8-inch furred creature</td>
    <td id="clue_DJ_5_4_r" class="clue_text" style="display:none;">(John: What is an inchworm?)<br /><br /><em class="correct_response">a pica</em><br /><br /><table width="100%"><tr><td class="wrong">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_6_4', 'clue_DJ_6_4_r', 'clue_DJ_6_4_stuck')" onmouseout="toggle('clue_DJ_6_4_r', 'clue_DJ_6_4', 'clue_DJ_6_4_stuck')" onclick="togglestick('clue_DJ_6_4_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_6_4_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$1600</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32766" title="Suggest a correction for this clue" rel="nofollow">11</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_6_4" class="clue_text">He was Jackie Kennedy's official White House designer</td>
    <td id="clue_DJ_6_4_r" class="clue_text" style="display:none;"><em class="correct_response">Oleg Cassini</em><br /><br /><table width="100%"><tr><td class="right">Diane</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_1_5', 'clue_DJ_1_5_r', 'clue_DJ_1_5_stuck')" onmouseout="toggle('clue_DJ_1_5_r', 'clue_DJ_1_5', 'clue_DJ_1_5_stuck')" onclick="togglestick('clue_DJ_1_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_1_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$2000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32755" title="Suggest a correction for this clue" rel="nofollow">16</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_1_5" class="clue_text">It may ring a bell that this capital of Belize also starts with "Bel"</td>
    <td id="clue_DJ_1_5_r" class="clue_text" style="display:none;"><em class="correct_response">Belmopan</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_2_5', 'clue_DJ_2_5_r', 'clue_DJ_2_5_stuck')" onmouseout="toggle('clue_DJ_2_5_r', 'clue_DJ_2_5', 'clue_DJ_2_5_stuck')" onclick="togglestick('clue_DJ_2_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_2_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$2000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32760" title="Suggest a correction for this clue" rel="nofollow">6</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_2_5" class="clue_text">In 1935 this great German stage director brought "A Midsummer Night's Dream" to the screen</td>
    <td id="clue_DJ_2_5_r" class="clue_text" style="display:none;"><em class="correct_response">Max Reinhardt</em><br /><br /><table width="100%"><tr><td class="wrong">Triple Stumper</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_3_5', 'clue_DJ_3_5_r', 'clue_DJ_3_5_stuck')" onmouseout="toggle('clue_DJ_3_5_r', 'clue_DJ_3_5', 'clue_DJ_3_5_stuck')" onclick="togglestick('clue_DJ_3_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_3_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$2000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32772" title="Suggest a correction for this clue" rel="nofollow">26</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_3_5" class="clue_text">October 14, 1912 on his way to a campaign rally in Milwaukee, Wisconsin</td>
    <td id="clue_DJ_3_5_r" class="clue_text" style="display:none;"><em class="correct_response">Theodore Roosevelt</em><br /><br /><table width="100%"><tr><td class="right">John</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_4_5', 'clue_DJ_4_5_r', 'clue_DJ_4_5_stuck')" onmouseout="toggle('clue_DJ_4_5_r', 'clue_DJ_4_5', 'clue_DJ_4_5_stuck')" onclick="togglestick('clue_DJ_4_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_4_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$2000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32776" title="Suggest a correction for this clue" rel="nofollow">28</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_4_5" class="clue_text">Elected in 1993, she's the first woman to represent Texas in the U.S. Senate</td>
    <td id="clue_DJ_4_5_r" class="clue_text" style="display:none;">(John: Who is Kay Bailey Hutchinson?)<br />(Alex: Say it again.)<br />(John: Hutchinson?  Who is Hutchinson?)<br />...<br />(Alex: [*], no "N" in the middle.)<br /><br /><em class="correct_response">Kay Bailey Hutchison</em><br /><br /><table width="100%"><tr><td class="wrong">John</td><td class="right">Chris</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_5_5', 'clue_DJ_5_5_r', 'clue_DJ_5_5_stuck')" onmouseout="toggle('clue_DJ_5_5_r', 'clue_DJ_5_5', 'clue_DJ_5_5_stuck')" onclick="togglestick('clue_DJ_5_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_5_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$2000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32780" title="Suggest a correction for this clue" rel="nofollow">21</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_5_5" class="clue_text">The dorsal spines give this fish <a href="http://www.j-archive.com/media/2005-10-25_DJ_21.jpg" target="_blank">seen</a> <a href="http://www.j-archive.com/media/2005-10-25_DJ_21a.jpg" target="_blank">here</a> its name</td>
    <td id="clue_DJ_5_5_r" class="clue_text" style="display:none;"><em class="correct_response">a stickleback</em><br /><br /><table width="100%"><tr><td class="wrong">Triple Stumper</td></tr></table></td>
  </tr>
</table>
    </td>
    <td class="clue">
<table>
  <tr>
    <td>
      <div onmouseover="toggle('clue_DJ_6_5', 'clue_DJ_6_5_r', 'clue_DJ_6_5_stuck')" onmouseout="toggle('clue_DJ_6_5_r', 'clue_DJ_6_5', 'clue_DJ_6_5_stuck')" onclick="togglestick('clue_DJ_6_5_stuck')">
        <table class="clue_header">
          <tr>
            <td id="clue_DJ_6_5_stuck" class="clue_unstuck">&#160;&#160;&#160;</td>
            <td class="clue_value">$2000</td>
            <td class="clue_order_number"><a href="suggestcorrection.php?clue_id=32767" title="Suggest a correction for this clue" rel="nofollow">12</a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
  <tr>
    <td id="clue_DJ_6_5" class="clue_text">This Latin abbreviation is used in footnotes</td>
    <td id="clue_DJ_6_5_r" class="clue_text" style="display:none;"><em class="correct_response"><i>op. cit.</i></em><br /><br /><table width="100%"><tr><td class="wrong">Triple Stumper</td></tr></table></td>
  </tr>
</table>
    </td>
  </tr>
</table>
              <h3>Scores at the end of the Double Jeopardy! Round:</h3>
              <table>
                <tr>
                  <td class="score_player_nickname">John</td>
                  <td class="score_player_nickname">Chris</td>
                  <td class="score_player_nickname">Diane</td>
                </tr>
                <tr>
                  <td class="score_positive">$12,300</td>
                  <td class="score_positive">$7,000</td>
                  <td class="score_positive">$11,400</td>
                </tr>
                <tr>
                  <td class="score_remarks"></td>
                  <td class="score_remarks"></td>
                  <td class="score_remarks"></td>
                </tr>
              </table>
              <p><a href="wageringcalculator.php?a=12300&amp;b=11400&amp;c=7000&amp;player_a=John&amp;player_b=Diane&amp;player_c=Chris">[wagering suggestions for these scores]</a></p>
            </div>
            <div id="final_jeopardy_round">
              <h2>Final Jeopardy! Round</h2>

<table class="final_round">
  <tr>
    <td class="category">
      
<div onmouseover="toggle('clue_FJ', 'clue_FJ_r', 'clue_FJ_stuck')" onmouseout="toggle('clue_FJ_r', 'clue_FJ', 'clue_FJ_stuck')" onclick="togglestick('clue_FJ_stuck')">
  <table>
    <tr><td class="category_name">YOUTH ORGANIZATIONS</td></tr>
    <tr><td class="category_comments"></td></tr>
  </table>
</div>
    </td>
  </tr>
  <tr>
    <td class="clue">
<table>
  <tr>
    <td id="clue_FJ" class="clue_text">This organization pledges it will strive for "clearer thinking... greater loyalty... larger service, and... better living" in that order</td>
    <td id="clue_FJ_r" class="clue_text" style="display:none;">(Alex: I know you were tempted to go for Scouts or Guides. Let's see what you came up with.)<br />...<br />(Alex: Head, heart, hands, health, you're right!)<br /><a href="http://www.j-archive.com/media/2005-10-25_FJ_ZergCam.jpg" target="_blank">[ZergCam]</a><table><tr><td class="wrong">Chris</td><td rowspan="2" valign="top">What is the Boy Scouts?</td></tr><tr><td>$2,600</td></tr><tr><td class="right">Diane</td><td rowspan="2" valign="top">What is the 4H?</td></tr><tr><td>$8,000</td></tr><tr><td class="wrong">John</td><td rowspan="2" valign="top">What is the Girl Scouts</td></tr><tr><td>$10,501</td></tr></table><em class="correct_response">4H</em></td>
  </tr>
</table>
<var id="clue_FJ_stuck"></var>
    </td>
   </tr>
</table>
              <h3>Final scores:</h3>
              <table>
                <tr>
                  <td class="score_player_nickname">John</td>
                  <td class="score_player_nickname">Chris</td>
                  <td class="score_player_nickname">Diane</td>
                </tr>
                <tr>
                  <td class="score_positive">$1,799</td>
                  <td class="score_positive">$4,400</td>
                  <td class="score_positive">$19,400</td>
                </tr>
                <tr>
                  <td class="score_remarks">3rd place: $1,000</td>
                  <td class="score_remarks">2nd place: $2,000</td>
                  <td class="score_remarks">New champion: $19,400</td>
                </tr>
              </table>

              <h3>Game dynamics:</h3>
              <img class="game_dynamics" src="chartgame.php?game_id=579" width="400" height="200" alt="Game dynamics graph" />

              <h3><a href="help.php#coryatscore">Coryat scores</a>:</h3>
              <table>
                <tr>
                  <td class="score_player_nickname">John</td>
                  <td class="score_player_nickname">Chris</td>
                  <td class="score_player_nickname">Diane</td>
                </tr>
                <tr>
                  <td class="score_positive">$12,600</td>
                  <td class="score_positive">$9,800</td>
                  <td class="score_positive">$11,400</td>
                </tr>
                <tr>
                  <td class="score_remarks">18 R<br />(including 1 DD),<br />4 W<br />(including 1 DD)</td>
                  <td class="score_remarks">15 R,<br />4 W<br />(including 1 DD)</td>
                  <td class="score_remarks">17 R,<br />0 W</td>
                </tr>
              </table>

              <h3><a href="help.php#combinedcoryat">Combined Coryat</a>: $33,800</h3>
              <h4><a href="showgameresponses.php?game_id=579" rel="external">[game responses]</a> <a href="showscores.php?game_id=579" rel="external">[game scores]</a> <a href="suggestcorrection.php?game_id=579" rel="nofollow">[suggest correction]</a></h4>
              <h6 style="margin-left: 0em;">Game tape date: 2005-08-03</h6>
            </div></div>`
	html_reader := strings.NewReader(html_raw)
	doc, err := html.Parse(html_reader)
	if err != nil {
		t.Fatal("failed to parse html_reader(string)", err)
	}
	content_div := nextDescendantWithClass(doc, "div", "content")

	jeid := qparty.EpisodeID(579)
	aired := qparty.ShowDate{Year: 2005, Month: 10, Day: 25}
	taped := qparty.ShowDate{Year: 2005, Month: 9, Day: 23}

	episode := new(qparty.FullEpisode)
	episode.EpisodeID = jeid
	episode.Taped = taped
	episode.Aired = aired
	parseContent(content_div, episode)

	if episode.EpisodeID != jeid {
		t.Error("did not carry the episode ID from the metadata definition", episode.EpisodeID)
	}
	if episode.Aired != aired || episode.Taped != taped {
		t.Error("did not carry the episode show dates from metadata", episode.EpisodeID)
	}

	expected_show_number := uint(4857)
	if episode.Show.Number != expected_show_number {
		t.Errorf("incorrect show number %d (see div#game_title %d)", episode.Show.Number, expected_show_number)
	}

	expected_comments := "(Kelly: We're here in wonderful [Copenhagen](/media/2005-10-25_ComingUp.jpg) celebrating 200 years of Hans Christian Andersen. Join us, next on *Jeopardy!* )"
	if episode.Comments != expected_comments {
		t.Errorf("incorrect comments %s (has link and italics)", episode.Comments)
	}

	// TODO more content checks
}
