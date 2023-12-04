package scrape

import (
	"net/http"
	"testing"
)

func httpServerSetup() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(cURLsample))
	})
	go http.ListenAndServe("localhost:8085", nil)
}

func TestScraper(t *testing.T) {
	httpServerSetup()
	Scrape("http://localhost:8085", "")
}

// the sample crawled webpage
var cURLsample string = `
<!DOCTYPE html>
<html lang="en-us">
<head>
<meta charset="utf-8"/>
<title>Private Leaderboard - Advent of Code 2023</title>
<link rel="stylesheet" type="text/css" href="/static/style.css?31"/>
<link rel="stylesheet alternate" type="text/css" href="/static/highcontrast.css?1" title="High Contrast"/>
<link rel="shortcut icon" href="/favicon.png"/>
<script>window.addEventListener('click', function(e,s,r){if(e.target.nodeName==='CODE'&&e.detail===3){s=window.getSelection();s.removeAllRanges();r=document.createRange();r.selectNodeContents(e.target);s.addRange(r);}});</script>
</head><!--




Oh, hello!  Funny seeing you here.

I appreciate your enthusiasm, but you aren't going to find much down here.
There certainly aren't clues to any of the puzzles.  The best surprises don't
even appear in the source until you unlock them for real.

Please be careful with automated requests; I'm not a massive company, and I can
only take so much traffic.  Please be considerate so that everyone gets to play.

If you're curious about how Advent of Code works, it's running on some custom
Perl code. Other than a few integrations (auth, analytics, social media), I
built the whole thing myself, including the design, animations, prose, and all
of the puzzles.

The puzzles are most of the work; preparing a new calendar and a new set of
puzzles each year takes all of my free time for 4-5 months. A lot of effort
went into building this thing - I hope you're enjoying playing it as much as I
enjoyed making it for you!

If you'd like to hang out, I'm @ericwastl@hachyderm.io on Mastodon and
@ericwastl on Twitter.

- Eric Wastl


















































-->
<body>
<header><div><h1 class="title-global"><a href="/">Advent of Code</a></h1><nav><ul><li><a href="/2023/about">[About]</a></li><li><a href="/2023/events">[Events]</a></li><li><a href="https://teespring.com/stores/advent-of-code" target="_blank">[Shop]</a></li><li><a href="/2023/settings">[Settings]</a></li><li><a href="/2023/auth/logout">[Log Out]</a></li></ul></nav><div class="user">PeeK1e <span class="star-count">4*</span></div></div><div><h1 class="title-event">&nbsp;&nbsp;&nbsp;<span class="title-event-wrap">sub y{</span><a href="/2023">2023</a><span class="title-event-wrap">}</span></h1><nav><ul><li><a href="/2023">[Calendar]</a></li><li><a href="/2023/support">[AoC++]</a></li><li><a href="/2023/sponsors">[Sponsors]</a></li><li><a href="/2023/leaderboard">[Leaderboard]</a></li><li><a href="/2023/stats">[Stats]</a></li></ul></nav></div></header>

<div id="sidebar">
<div id="sponsor"><div class="quiet">Our <a href="/2023/sponsors">sponsors</a> help make Advent of Code possible:</div><div class="sponsor"><a href="https://it-jobs.de/?sta_cmp=advent_2023" target="_blank" onclick="if(ga)ga('send','event','sponsor','sidebar',this.href);" rel="noopener">it-jobs.de</a> - designed with love for IT ..&quot;&quot;&quot;.....&quot;&quot;&quot; .&quot;....&quot;.&quot;....&quot; .&quot;.....&quot;.....&quot; ..&quot;.........&quot; ...&quot;.......&quot; .....&quot;...&quot; .......&quot;</div></div>
</div><!--/sidebar-->

<main>
<article>
<p>This is the private leaderboard of dj95 for <em>Advent of Code 2023</em>. You can use a different <a href="javascript:void(0)" onclick="ordering_show()">[Ordering]</a>, manage your <a href="/2023/leaderboard/private">[Private Leaderboards]</a>, use an <a href="javascript:void(0)" onclick="api_show()">[API]</a>, or switch to another <a href="/2023/events">[Event]</a>.</p>
<script>function api_show() { document.getElementById("api_info").style.display = "block"; }</script>
<p id="api_info" style="display:none;">You can also get the <a href="/2023/leaderboard/private/view/3398843.json" target="_blank">[JSON]</a> for this private leaderboard. <em>Please</em> don't make frequent automated requests to this service - avoid sending requests more often than once every <em>15 minutes</em> (<code>900</code> seconds). If you do this from a script, you'll have to provide your session cookie in the request; a fresh session cookie lasts for about a month. Timestamps use <a href="https://en.wikipedia.org/wiki/Unix_time">Unix time</a>.</p>
<script>function ordering_show() { document.getElementById("ordering_info").style.display = "block"; }</script>
<div id="ordering_info" style="display:none;">
<p>There are several different ordering methods available:</p>
<ul>
<li><a href="?order=local_score">[Local Score]</a>, which awards users on this leaderboard points much like the global leaderboard. If you add or remove users, the points will be recalculated, and the order can change. For <code>N</code> users, the first user to get each star gets <code>N</code> points, the second gets <code>N-1</code>, and the last gets <code>1</code>. This is the default.</li>
<li><a href="?order=global_score">[Global Score]</a>, which uses scores from the global leaderboard. Ties are broken by the user's local score.</li>
<li><a href="?order=stars">[Stars]</a>, which uses the number of stars the user has. Ties are broken by the time the most recent star was acquired. This used to be the default.</li>
</ul>
</div>
<p><span class="privboard-star-both">Gold</span> indicates the user got both stars for that day, <span class="privboard-star-firstonly">silver</span> means just the first star, and <span class="privboard-star-unlocked">gray</span> means none.</p>
<div class="privboard-row">      <span class="privboard-days"><a href="/2023/day/1">1</a><a href="/2023/day/2">2</a><a href="/2023/day/3">3</a><span>4</span><span>5</span><span>6</span><span>7</span><span>8</span><span>9</span><span>1<br/>0</span><span>1<br/>1</span><span>1<br/>2</span><span>1<br/>3</span><span>1<br/>4</span><span>1<br/>5</span><span>1<br/>6</span><span>1<br/>7</span><span>1<br/>8</span><span>1<br/>9</span><span>2<br/>0</span><span>2<br/>1</span><span>2<br/>2</span><span>2<br/>3</span><span>2<br/>4</span><span>2<br/>5</span></span></div>
<div class="privboard-row"><span class="privboard-position">1)</span> 21 <span class="privboard-star-both">*</span><span class="privboard-star-both">*</span><span class="privboard-star-both">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span>  <span class="privboard-name">dj95</span></div>
<div class="privboard-row"><span class="privboard-position">2)</span> 13 <span class="privboard-star-both">*</span><span class="privboard-star-both">*</span><span class="privboard-star-unlocked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span>  <span class="privboard-name"><a href="https://github.com/PeeK1e" target="_blank">PeeK1e</a></span></div>
<div class="privboard-row"><span class="privboard-position">3)</span>  6 <span class="privboard-star-both">*</span><span class="privboard-star-unlocked">*</span><span class="privboard-star-unlocked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span>  <span class="privboard-name">thled</span></div>
<div class="privboard-row"><span class="privboard-position">4)</span>  0 <span class="privboard-star-unlocked">*</span><span class="privboard-star-unlocked">*</span><span class="privboard-star-unlocked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span><span class="privboard-star-locked">*</span>  <span class="privboard-name"><a href="https://github.com/Mawarii" target="_blank">Mawarii</a></span></div>
</article>
</main>

<!-- ga -->
<script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', 'UA-69522494-1', 'auto');
ga('set', 'anonymizeIp', true);
ga('send', 'pageview');
</script>
<!-- /ga -->
</body>
`
