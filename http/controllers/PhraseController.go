package controllers

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
	xmlContent := []byte(`<div class="mw-parser-output"><p><strong class="published"><span id="publishDate" class="value-title" title="2007-07-13"></span>Friday, July 13, 2007</strong><span class="Z3988" style="display:none" title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rfr_id=rfr_id=info:sid/en.wikinews.org:User%3ADendodge&amp;rft_id=info:sid/en.wikinews.org:Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;rft.atitle=Argentine+Supreme+Court+declares+Riveros+pardon+unconstitutional&amp;rft.jtitle=Wikinews&amp;rft.date=2007-07-13&amp;rft.aucorp=Wikinews+contributors&amp;rft.genre=article&amp;rft.rights=CC-BY+2.5&amp;rft.identifier=%2F%2Fen.wikinews.org%2Fwiki%2FArgentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;rft.artnum=-">&#160;</span>
</p>
<div class="infobox noprint desktop-only" style="margin-left: 5px; float: right; width: 200px; background: #edf7ff; border: solid #666666 1px; padding: 0px; font-size: x-small; clear:right;">
<div style="text-decoration:none; background: #7ec9fd; padding: 2px; text-align: center; style; font-size: 140%; border-bottom: solid 1px #666666;"><b><a href="/wiki/Argentina" class="mw-redirect" title="Argentina">Argentina</a> <br /><div class="center"><div class="floatnone"><a href="/wiki/File:Wikinoticias_Argentina.svg" class="image"><img alt="Wikinoticias Argentina.svg" src="//upload.wikimedia.org/wikipedia/commons/thumb/5/5a/Wikinoticias_Argentina.svg/80px-Wikinoticias_Argentina.svg.png" decoding="async" width="80" height="67" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/5/5a/Wikinoticias_Argentina.svg/120px-Wikinoticias_Argentina.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/5/5a/Wikinoticias_Argentina.svg/160px-Wikinoticias_Argentina.svg.png 2x" data-file-width="612" data-file-height="515" /></a></div></div></b></div>
<div style="background: #bce1ff; padding: 2px; text-align: center; style; font-size: normal;">Other stories from Argentina</div>
<div style="padding: 2px;">
<ul>
<li>27 January 2019: <a href="/wiki/Male_Magellanic_penguins_pine_for_pairings:_Wikinews_interviews_biologist_Natasha_Gownaris" title="Male Magellanic penguins pine for pairings: Wikinews interviews biologist Natasha Gownaris">Male Magellanic penguins pine for pairings: Wikinews interviews biologist Natasha Gownaris</a></li> 
<li>12 January 2019: <a href="/wiki/State-run_bus_crashes_in_Cuba_en_route_to_Havana,_killing_seven" title="State-run bus crashes in Cuba en route to Havana, killing seven">State-run bus crashes in Cuba en route to Havana, killing seven</a></li> 
<li>3 July 2018: <a href="/wiki/FIFA_World_Cup_2018_Last_16:_France,_Uruguay_send_Argentina,_Portugal_home" title="FIFA World Cup 2018 Last 16: France, Uruguay send Argentina, Portugal home">FIFA World Cup 2018 Last 16: France, Uruguay send Argentina, Portugal home</a></li> 
<li>1 July 2018: <a href="/wiki/FIFA_World_Cup_2018_day_12,_13,_14,_15:_Iran,_Nigeria,_Germany,_Senegal_out_of_the_tournament" title="FIFA World Cup 2018 day 12, 13, 14, 15: Iran, Nigeria, Germany, Senegal out of the tournament">FIFA World Cup 2018 day 12, 13, 14, 15: Iran, Nigeria, Germany, Senegal out of the tournament</a></li> 
<li>1 July 2018: <a href="/wiki/Argentine_footballer_Mascherano_announces_international_retirement" title="Argentine footballer Mascherano announces international retirement">Argentine footballer Mascherano announces international retirement</a></li></ul>
</div>
<small><div style="text-align: right;">...<a href="/wiki/Portal:Argentina" class="mw-redirect" title="Portal:Argentina">More articles here</a></div></small>
<div style="background: #bce1ff; padding: 2px; text-align: center; style; font-size: normal;">Location of Argentina</div>
<div style="padding: 2px;">
<p><a href="/wiki/File:LocationArgentina.png" class="image" title="A map showing the location of Argentina"><img alt="A map showing the location of Argentina" src="//upload.wikimedia.org/wikipedia/commons/thumb/d/dd/LocationArgentina.png/196px-LocationArgentina.png" decoding="async" width="196" height="91" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/d/dd/LocationArgentina.png/294px-LocationArgentina.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/d/dd/LocationArgentina.png/392px-LocationArgentina.png 2x" data-file-width="1357" data-file-height="628" /></a>
</p>
</div>
<center><a href="/wiki/Portal:Argentina" class="mw-redirect" title="Portal:Argentina">To write, edit, start or view other articles on Argentina, see the Argentina Portal</a></center>
<div style="float: center; text-align: center;">
<center><a href="/wiki/Portal:Argentina" title="Portal:Argentina"><img alt="Flag of Argentina.svg" src="//upload.wikimedia.org/wikipedia/commons/thumb/1/1a/Flag_of_Argentina.svg/60px-Flag_of_Argentina.svg.png" decoding="async" width="60" height="38" class="thumbborder" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/1/1a/Flag_of_Argentina.svg/90px-Flag_of_Argentina.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/1/1a/Flag_of_Argentina.svg/120px-Flag_of_Argentina.svg.png 2x" data-file-width="800" data-file-height="500" /></a>
</center>
</div>
</div>
<p>The <a href="https://en.wikipedia.org/wiki/Supreme_Court_of_Argentina" class="extiw" title="w:Supreme Court of Argentina">Supreme Court of Argentina</a> today ruled that a decree issued by the Argentine ex-president <a href="https://en.wikipedia.org/wiki/Carlos_Sa%C3%BAl_Menem" class="extiw" title="w:Carlos Saúl Menem">Carlos Saúl Menem</a> was unconstitutional. The decree pardoned military general Santiago Omar Riveros, charged for crimes against humanity which were committed during the "Guerra Sucia", a military rule (1976-1983) in <a href="/wiki/Argentina" class="mw-redirect" title="Argentina">Argentina</a>.
</p><p>The members of the Supreme Court of Justice decided by a majority of four votes the revoke of the pardon towards Riveros. The votes were as follows:
</p>
<dl><dt>Support</dt></dl>
<ul><li>Ricardo Lorenzetti</li>
<li>Juan Carlos Maqueda</li>
<li>Eugenio Zaffaroni</li>
<li>Elena Highton</li></ul>
<dl><dt>Oppose</dt></dl>
<ul><li>Carlos Fayt</li>
<li>Carmen Argibay</li></ul>
<dl><dt>Abstain</dt></dl>
<ul><li>Enrique Petracchi</li></ul>
<p>The Supreme Court thus retracted the pardon towards Riveros. The ruling also left an open door for the cases of other pardoned generals, such as <a href="https://en.wikipedia.org/wiki/Jorge_Videla" class="extiw" title="w:Jorge Videla">Jorge Videla</a> and <a href="https://en.wikipedia.org/wiki/Eduardo_Massera" class="extiw" title="w:Eduardo Massera">Eduardo Massera</a>.
</p><p>"Crimes against humanity, because of their seriousness in nature, are not only contrary to the Constitución Nacional, but as well to the entire international community," the Supreme Court ruled. Nearly 30,000 disappeared in the Argentina military dictatiorship.
</p>
<h2><span class="mw-headline" id="Sources">Sources</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;action=edit&amp;section=1" title="Edit section: Sources">edit</a><span class="mw-editsection-bracket">]</span></span></h2>
<div class="desktop-only"><table class="plainlinks xambox xambox-type-license"><tbody><tr>
<td class="xambox-image"> <a href="/wiki/File:Wikinews-logo.svg" class="image" title="Wikinews in Spanish"><img alt="Wikinews in Spanish" src="//upload.wikimedia.org/wikipedia/commons/thumb/2/24/Wikinews-logo.svg/35px-Wikinews-logo.svg.png" decoding="async" width="35" height="19" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/2/24/Wikinews-logo.svg/53px-Wikinews-logo.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/2/24/Wikinews-logo.svg/70px-Wikinews-logo.svg.png 2x" data-file-width="759" data-file-height="415" /></a></td>
<td>
<p>This is a complete or partial translation of the article "<i><a href="https://es.wikinews.org/wiki/La_Corte_Suprema_de_Argentina_revoc%C3%B3_los_indultos_a_los_jefes_militares" class="extiw" title="es:La Corte Suprema de Argentina revocó los indultos a los jefes militares">La Corte Suprema de Argentina revocó los indultos a los jefes militares</a></i>", from the <a href="https://es.wikinews.org/wiki/" class="extiw" title="es:">Spanish language Wikinews</a>, published under the <a rel="nofollow" class="external text" href="https://creativecommons.org/licenses/by/2.5/">Creative Commons Attribution 2.5 License</a>.
</p>
</td>

</tr></tbody></table></div><div class="mobile-only"><table class="plainlinks xambox-mobile xambox-type-license"><tbody><tr>
<td> This is a complete or partial translation of the article "<i><a href="https://es.wikinews.org/wiki/La_Corte_Suprema_de_Argentina_revoc%C3%B3_los_indultos_a_los_jefes_militares" class="extiw" title="es:La Corte Suprema de Argentina revocó los indultos a los jefes militares">La Corte Suprema de Argentina revocó los indultos a los jefes militares</a></i>", from the <a href="https://es.wikinews.org/wiki/" class="extiw" title="es:">Spanish language Wikinews</a>, published under the <a rel="nofollow" class="external text" href="https://creativecommons.org/licenses/by/2.5/">Creative Commons Attribution 2.5 License</a>. </td>
</tr></tbody></table></div>
<ul><li><span class="sourceTemplate">Daniel Schweimler. "<a rel="nofollow" class="external text" href="http://news.bbc.co.uk/2/hi/americas/6898289.stm">Argentine court overturns pardon</a>" &#8212;&#160;<i><span class="interwiki-link-local"><a href="/wiki/BBC_News" class="mw-redirect" title="BBC News">BBC News</a></span></i>, July 13, 2007</span></li>
<li><span class="sourceTemplate"> "<a rel="nofollow" class="external text" href="http://www.clarin.com/diario/2007/07/13/um/m-01456528.htm">La Corte Suprema anuló los indultos de Carlos Menem a jefes militares</a>" &#8212;&#160;<i><span class="interwiki-link-foreign"><a href="https://en.wikipedia.org/wiki/Clar%C3%ADn_(newspaper)" class="extiw" title="w:Clarín (newspaper)">Clarín (newspaper)</a></span></i>, July 13, 2007 <span style="font-family: sans-serif; cursor: default; color: #0000CD; font-size: 1.2em; font-weight: bold;" title="Language: Spanish">(<span style="color: #000080;font-size: 0.7em; position: relative; bottom: 0.1em"><span class="interwiki-link-foreign"><a href="https://en.wikipedia.org/wiki/Spanish_language" class="extiw" title="w:Spanish language">Spanish</a></span></span>)</span></span></li>
<li><span class="sourceTemplate"> "<a rel="nofollow" class="external text" href="http://www.lanacion.com.ar/politica/nota.asp?nota_id=925389&amp;pid=2859908&amp;toi=5245">La Corte Suprema anuló los indultos</a>" &#8212;&#160;<i><span class="interwiki-link-foreign"><a href="https://en.wikipedia.org/wiki/La_Naci%C3%B3n" class="extiw" title="w:La Nación">La Nación</a></span></i>, July 13, 2007 <span style="font-family: sans-serif; cursor: default; color: #0000CD; font-size: 1.2em; font-weight: bold;" title="Language: Spanish">(<span style="color: #000080;font-size: 0.7em; position: relative; bottom: 0.1em"><span class="interwiki-link-foreign"><a href="https://en.wikipedia.org/wiki/Spanish_language" class="extiw" title="w:Spanish language">Spanish</a></span></span>)</span></span><br style="clear:both" /></li></ul>
<table align="center" style="width:25%; background-color:#FFFFFF; border:1px solid #a7d7f9; -moz-border-radius: 9px;-webkit-border-radius: 9px; border-radius: 9px; padding:1px;" id="social_bookmarks" class="noprint">

<tbody><tr>
<td><div class="floatleft"><a href="http://en.wikinews.org/wiki/Wikinews:Social_bookmarks"><img alt="Bookmark-new.svg" src="//upload.wikimedia.org/wikipedia/commons/thumb/f/fb/Bookmark-new.svg/25px-Bookmark-new.svg.png" decoding="async" width="25" height="25" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/f/fb/Bookmark-new.svg/38px-Bookmark-new.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/f/fb/Bookmark-new.svg/50px-Bookmark-new.svg.png 2x" data-file-width="48" data-file-height="48" /></a></div>
<div class="plainlinks" align="center">
<p><b>Share this:</b>&#160;
<span title="Share via e-mail" class="plainlinks"><a href="mailto:?subject=Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional%20–%20Wikinews&amp;body=Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional:%0Ahttps://en.wikinews.org/wiki/Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional%0A%0AFrom%20Wikinews,%20the%20free%20news%20source%20you%20can%20write!" title="E-mail this story" rel="nofollow"><img alt="E-mail this story" src="//upload.wikimedia.org/wikipedia/commons/5/5c/Email.png" decoding="async" width="16" height="16" data-file-width="16" data-file-height="16" /></a></span>
<span title="Share on Facebook"><a href="http://www.facebook.com/sharer.php?u=https://en.wikinews.org/wiki/Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;t=Argentine+Supreme+Court+declares+Riveros+pardon+unconstitutional+-+Wikinews" title="Share on Facebook" rel="nofollow"><img alt="Share on Facebook" src="//upload.wikimedia.org/wikinews/en/5/55/Facebook.png" decoding="async" width="16" height="16" data-file-width="16" data-file-height="16" /></a></span>
<span title="Share on Digg"><a href="http://digg.com/submit?url=https://en.wikinews.org/wiki/Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;title=Argentine+Supreme+Court+declares+Riveros+pardon+unconstitutional+-+Wikinews" title="Share on Digg.com" rel="nofollow"><img alt="Share on Digg.com" src="//upload.wikimedia.org/wikinews/en/9/95/Digg-icon.png" decoding="async" width="16" height="14" data-file-width="16" data-file-height="14" /></a></span>
<span title="Share on reddit"><a href="http://reddit.com/submit?url=https://en.wikinews.org/wiki/Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;title=Argentine+Supreme+Court+declares+Riveros+pardon+unconstitutional+-+Wikinews" title="Share on reddit.com" rel="nofollow"><img alt="Share on reddit.com" src="//upload.wikimedia.org/wikinews/en/1/10/Reddit.png" decoding="async" width="18" height="18" data-file-width="18" data-file-height="18" /></a></span>
<span title="Share on LinkedIn"><a href="http://www.linkedin.com/shareArticle?mini=true&amp;url=https://en.wikinews.org/wiki/Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;title=Argentine+Supreme+Court+declares+Riveros+pardon+unconstitutional+-+Wikinews" title="Share on LinkedIn.com" rel="nofollow"><img alt="Share on LinkedIn.com" src="//upload.wikimedia.org/wikipedia/commons/thumb/c/c9/Linkedin.svg/16px-Linkedin.svg.png" decoding="async" width="16" height="16" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/c/c9/Linkedin.svg/24px-Linkedin.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/c/c9/Linkedin.svg/32px-Linkedin.svg.png 2x" data-file-width="200" data-file-height="200" /></a></span>
<span title="Share on Twitter"><a href="http://twitter.com/?status=Look%20what%20I%20found%20on%20Wikinews:%20https://en.wikinews.org/wiki/Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional" title="Share on twitter.com" rel="nofollow"><img alt="Share on twitter.com" src="//upload.wikimedia.org/wikinews/en/f/f7/Twitter.png" decoding="async" width="18" height="18" data-file-width="18" data-file-height="18" /></a></span>
</p>
</div>
</td></tr></tbody></table>
<p><br style="clear:both;" />
</p>
<div class="desktop-only"><table class="plainlinks noprint xambox xambox-type-notice"><tbody><tr>
<td class="xambox-image"> <a href="/wiki/File:Padlock-silver-medium.svg" class="image"><img alt="Padlock-silver-medium.svg" src="//upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Padlock-silver-medium.svg/40px-Padlock-silver-medium.svg.png" decoding="async" width="40" height="40" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Padlock-silver-medium.svg/60px-Padlock-silver-medium.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Padlock-silver-medium.svg/80px-Padlock-silver-medium.svg.png 2x" data-file-width="128" data-file-height="128" /></a></td>
<td>
<p><b>This page is <a href="/wiki/Wikinews:Archive_conventions" title="Wikinews:Archive conventions">archived</a>, and is no longer publicly editable.</b>
</p>
<div style="padding: 5px 40px 0px 40px; font-size:90%;">Articles presented on Wikinews reflect the specific time at which they were written and published, and do not attempt to encompass events or knowledge which occur or become known after their publication.</div>
<p>Got a correction? Add the template {{<a href="/wiki/Template:Editprotected" title="Template:Editprotected">editprotected</a>}} to the <a href="/w/index.php?title=Talk:Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;action=edit&amp;redlink=1" class="new" title="Talk:Argentine Supreme Court declares Riveros pardon unconstitutional (page does not exist)">talk page</a> along with your corrections, and it will be <a href="/wiki/Wikinews:Admin_action_alerts#Edits_to_protected_pages" title="Wikinews:Admin action alerts">brought to the attention of</a> the <a href="/wiki/Wikinews:Administrators" title="Wikinews:Administrators">administrators</a>.
</p><p><small>Please note that due to our <a href="/wiki/Wikinews:ARCHIVE" class="mw-redirect" title="Wikinews:ARCHIVE">archival policy</a>, we will not alter or update the content of articles that are archived, but will only accept requests to make grammatical and formatting corrections.</small>
</p><p><small>Note that some listed sources or external links may no longer be available online due to age.</small>
</p>
</td>

</tr></tbody></table></div><div class="mobile-only"><table class="plainlinks noprint xambox-mobile xambox-type-notice"><tbody><tr>
<td class="xambox-image"> <a href="/wiki/File:Padlock-silver-medium.svg" class="image"><img alt="Padlock-silver-medium.svg" src="//upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Padlock-silver-medium.svg/40px-Padlock-silver-medium.svg.png" decoding="async" width="40" height="40" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Padlock-silver-medium.svg/60px-Padlock-silver-medium.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Padlock-silver-medium.svg/80px-Padlock-silver-medium.svg.png 2x" data-file-width="128" data-file-height="128" /></a></td>
<td> <b>This page is <a href="/wiki/Wikinews:Archive_conventions" title="Wikinews:Archive conventions">archived</a>, and is no longer publicly editable.</b>
<div style="padding: 5px 40px 0px 40px; font-size:90%;">Articles presented on Wikinews reflect the specific time at which they were written and published, and do not attempt to encompass events or knowledge which occur or become known after their publication.</div>
<p>Got a correction? Add the template {{<a href="/wiki/Template:Editprotected" title="Template:Editprotected">editprotected</a>}} to the <a href="/w/index.php?title=Talk:Argentine_Supreme_Court_declares_Riveros_pardon_unconstitutional&amp;action=edit&amp;redlink=1" class="new" title="Talk:Argentine Supreme Court declares Riveros pardon unconstitutional (page does not exist)">talk page</a> along with your corrections, and it will be <a href="/wiki/Wikinews:Admin_action_alerts#Edits_to_protected_pages" title="Wikinews:Admin action alerts">brought to the attention of</a> the <a href="/wiki/Wikinews:Administrators" title="Wikinews:Administrators">administrators</a>.
</p><p><small>Please note that due to our <a href="/wiki/Wikinews:ARCHIVE" class="mw-redirect" title="Wikinews:ARCHIVE">archival policy</a>, we will not alter or update the content of articles that are archived, but will only accept requests to make grammatical and formatting corrections.</small>
</p>
<small>Note that some listed sources or external links may no longer be available online due to age.</small> </td>

</tr></tbody></table></div>
<!-- 
NewPP limit report
Parsed by mw1222
Cached time: 20191112185907
Cache expiry: 86400
Dynamic content: true
Complications: []
CPU time usage: 0.120 seconds
Real time usage: 0.184 seconds
Preprocessor visited node count: 738/1000000
Preprocessor generated node count: 0/1500000
Post‐expand include size: 30084/2097152 bytes
Template argument size: 5322/2097152 bytes
Highest expansion depth: 18/40
Expensive parser function count: 5/500
Unstrip recursion depth: 0/20
Unstrip post‐expand size: 1453/5000000 bytes
Number of Wikibase entities loaded: 0/400
-->
<!--
Transclusion expansion time report (%,ms,calls,template)
100.00%  109.191      1 -total
 34.49%   37.662      3 Template:Source
 24.81%   27.089      1 Template:Argentina
 15.56%   16.991      1 Template:Date
 12.37%   13.509      3 Template:Source/pub
 11.71%   12.787      2 Template:Source/lang
 10.99%   12.003      2 Template:Infosection
 10.53%   11.497      5 Template:W
 10.28%   11.226      2 Template:Xambox
  8.85%    9.664      1 Template:Archived
-->

<!-- Saved in parser cache with key enwikinews:pcache:idhash:74296-0!canonical and timestamp 20191112185907 and revision id 3081652
 -->
</div>`)
	err := xml.Unmarshal(xmlContent, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}
