package main

import (
	"encoding/xml"
	"fmt"
)

type div struct {
	XMLName  xml.Name `xml:"div"`
	SpanList []p `xml:"p"`
}

type p struct {
	Text string `xml:",chardata"`
}

func main() {
	d := div{}
	xmlContent := []byte(`<div class="mw-parser-output"><p>The time is 22:00 (UTC) on August 16th, 2006, and this is Audio Wikinews News Briefs.
</p>
<div id="toc" class="toc"><input type="checkbox" role="button" id="toctogglecheckbox" class="toctogglecheckbox" style="display:none" /><div class="toctitle" lang="en" dir="ltr"><h2>Contents</h2><span class="toctogglespan"><label class="toctogglelabel" for="toctogglecheckbox"></label></span></div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#Headlines"><span class="tocnumber">1</span> <span class="toctext">Headlines</span></a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#6_Canadian_soldiers_injured_in_Afghan_attack"><span class="tocnumber">1.1</span> <span class="toctext">6 Canadian soldiers injured in Afghan attack</span></a></li>
<li class="toclevel-2 tocsection-3"><a href="#United_Airlines_Flight_923_makes_emergency_landing_at_Logan_International_Airport"><span class="tocnumber">1.2</span> <span class="toctext">United Airlines Flight 923 makes emergency landing at Logan International Airport</span></a></li>
<li class="toclevel-2 tocsection-4"><a href="#Fire_in_Amherstburg,_Ontario_dies_down"><span class="tocnumber">1.3</span> <span class="toctext">Fire in Amherstburg, Ontario dies down</span></a></li>
<li class="toclevel-2 tocsection-5"><a href="#Four_Lotto_649_winners_share_jackpot"><span class="tocnumber">1.4</span> <span class="toctext">Four Lotto 649 winners share jackpot</span></a></li>
<li class="toclevel-2 tocsection-6"><a href="#New_Zealand_&#39;Boobs_on_Bikes&#39;_parade_approved"><span class="tocnumber">1.5</span> <span class="toctext">New Zealand 'Boobs on Bikes' parade approved</span></a></li>
</ul>
</li>
<li class="toclevel-1 tocsection-7"><a href="#Closing_statements"><span class="tocnumber">2</span> <span class="toctext">Closing statements</span></a></li>
</ul>
</div>

<h2><span class="mw-headline" id="Headlines">Headlines</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=News_briefs:August_16,_2006&amp;action=edit&amp;section=1" title="Edit section: Headlines">edit</a><span class="mw-editsection-bracket">]</span></span></h2>
<h3><span class="mw-headline" id="6_Canadian_soldiers_injured_in_Afghan_attack"><a href="/wiki/6_Canadian_soldiers_injured_in_Afghan_attack" title="6 Canadian soldiers injured in Afghan attack">6 Canadian soldiers injured in Afghan attack</a></span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=News_briefs:August_16,_2006&amp;action=edit&amp;section=2" title="Edit section: 6 Canadian soldiers injured in Afghan attack">edit</a><span class="mw-editsection-bracket">]</span></span></h3>
<p>Afghanistan
</p><p>Six soldiers were injured, Tuesday, during a mortar attack on a Canadian outpost in Afghanistan. None of the injuries sustained are life threatening, according to officials.
</p><p><br />
</p>
<h3><span class="mw-headline" id="United_Airlines_Flight_923_makes_emergency_landing_at_Logan_International_Airport"><a href="/wiki/United_Airlines_Flight_923_makes_emergency_landing_at_Logan_International_Airport" title="United Airlines Flight 923 makes emergency landing at Logan International Airport">United Airlines Flight 923 makes emergency landing at Logan International Airport</a></span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=News_briefs:August_16,_2006&amp;action=edit&amp;section=3" title="Edit section: United Airlines Flight 923 makes emergency landing at Logan International Airport">edit</a><span class="mw-editsection-bracket">]</span></span></h3>
<p>United States
</p><p>United Airlines Flight 923 from London to Washington D.C. made an emergency landing at Boston's Logan International Airport following the pilot declaring an in-flight emergency.
</p>
<h3><span id="Fire_in_Amherstburg.2C_Ontario_dies_down"></span><span class="mw-headline" id="Fire_in_Amherstburg,_Ontario_dies_down"><a href="/wiki/Fire_in_Amherstburg,_Ontario_dies_down" title="Fire in Amherstburg, Ontario dies down">Fire in Amherstburg, Ontario dies down</a></span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=News_briefs:August_16,_2006&amp;action=edit&amp;section=4" title="Edit section: Fire in Amherstburg, Ontario dies down">edit</a><span class="mw-editsection-bracket">]</span></span></h3>
<p>A huge fire at a plastics recycling plant forced hundreds of residents in Amherstburg, a Southwestern Ontario town, flee their homes.
</p><p><br />
</p>
<h3><span class="mw-headline" id="Four_Lotto_649_winners_share_jackpot"><a href="/wiki/Four_Lotto_649_winners_share_jackpot" title="Four Lotto 649 winners share jackpot">Four Lotto 649 winners share jackpot</a></span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=News_briefs:August_16,_2006&amp;action=edit&amp;section=5" title="Edit section: Four Lotto 649 winners share jackpot">edit</a><span class="mw-editsection-bracket">]</span></span></h3>
<p>Canada
</p><p>Three people in the Greater Toronto Area and one person in Quebec won last week's Lotto 649. They will share the second-largest lottery jackpot in Canadian history as the lucky winners of Saturday's $43.2-million Lotto 649 draw. They will each take home $10.8 million.
</p><p><br />
</p>
<h3><span id="New_Zealand_.27Boobs_on_Bikes.27_parade_approved"></span><span class="mw-headline" id="New_Zealand_'Boobs_on_Bikes'_parade_approved"><a href="/wiki/New_Zealand_%27Boobs_on_Bikes%27_parade_approved" title="New Zealand &#39;Boobs on Bikes&#39; parade approved">New Zealand 'Boobs on Bikes' parade approved</a></span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=News_briefs:August_16,_2006&amp;action=edit&amp;section=6" title="Edit section: New Zealand &#039;Boobs on Bikes&#039; parade approved">edit</a><span class="mw-editsection-bracket">]</span></span></h3>
<p>New Zealand
</p><p>The Auckland 'Boobs on Bikes' parade for next weeks annual sex industry convention, Erotica Expo which features naked woman on motorcycles has been approved by Auckland City Council's events and promotion staff despite disapproval by Auckland Mayor, Dick Hubbard and 13 other Auckland City Council councillors.
</p><p><br />
</p>
<h2><span class="mw-headline" id="Closing_statements">Closing statements</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=News_briefs:August_16,_2006&amp;action=edit&amp;section=7" title="Edit section: Closing statements">edit</a><span class="mw-editsection-bracket">]</span></span></h2>
<p>We invite you to visit wikinews.org for up-to-date news and information. This has been Audio Wikinews Newsbriefs. Thank you for listening, and enjoy the rest of your day.
This recording has been released under the Creative Commons Attribution 2.5 License.
</p>
<!-- 
NewPP limit report
Parsed by mw1333
Cached time: 20191028182635
Cache expiry: 2592000
Dynamic content: false
Complications: []
CPU time usage: 0.004 seconds
Real time usage: 0.009 seconds
Preprocessor visited node count: 21/1000000
Preprocessor generated node count: 0/1500000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/500
Unstrip recursion depth: 0/20
Unstrip post‐expand size: 0/5000000 bytes
Number of Wikibase entities loaded: 0/400
-->
<!--
Transclusion expansion time report (%,ms,calls,template)
100.00%    0.000      1 -total
-->

<!-- Saved in parser cache with key enwikinews:pcache:idhash:47554-0!canonical and timestamp 20191028182635 and revision id 424379
 -->
</div>`)
	err := xml.Unmarshal(xmlContent, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}
