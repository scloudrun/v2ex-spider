package parse

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_ParseArticles(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlList))

	if err != nil {
		t.Errorf("goquery fail")
	}

	articles := ParseArticles(doc)

	for k, v := range articles {
		fmt.Printf("序号: %d   评论: %s \tT: %s\n", k, v.Comment, v.Title)
	}
}

func Test_ParseArticle(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		t.Errorf("goquery fail")
	}

	article := ParseArticle(doc)

	fmt.Println(article.Content)
	for _, v := range article.CommentContent {
		fmt.Println(v.Floor, "楼:", v.Content)
	}
}

var htmlList = `<body>
<div id="Top">
<div class="content">
<div style="padding-top: 6px;">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="110" align="left"><a href="/" name="top" title="way to explore"><div id="Logo"></div></a></td>
<td width="auto" align="left">
<div id="Search"><form action="https://www.google.com" onsubmit="return dispatch()" target="_blank"><div id="qbar"><input type="text" maxlength="40" name="q" id="q" value="" onfocus="$('#qbar').addClass('qbar_focus')" onblur="$('#qbar').removeClass('qbar_focus')"></div></form></div>
</td>
<td width="570" align="right" style="padding-top: 2px;"><a href="/" class="top">首页</a>&nbsp;&nbsp;&nbsp;<a href="/signup" class="top">注册</a>&nbsp;&nbsp;&nbsp;<a href="/signin" class="top">登录</a></td>
</tr>
</tbody></table>
</div>
</div>
</div>
<div id="Wrapper">
<div class="content">
<div id="Leftbar"></div>
<div id="Rightbar">
<div class="sep20"></div>
<div class="box">
<div class="cell">
<strong>V2EX = way to explore</strong>
<div class="sep5"></div>
<span class="fade">V2EX 是一个关于分享和探索的地方</span>
</div>
<div class="inner">
<div class="sep5"></div>
<div align="center"><a href="/signup" class="super normal button">现在注册</a>
<div class="sep5"></div>
<div class="sep10"></div>
已注册用户请 &nbsp;<a href="/signin">登录</a></div>
</div>
</div>
<div class="sep20"></div>
<div class="box">
<div class="inner">
<script async="" src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>

<ins class="adsbygoogle" style="display:inline-block;width:250px;height:250px" data-ad-client="ca-pub-3465543440750523" data-ad-slot="9619519096" data-adsbygoogle-status="done"><ins id="aswift_0_expand" style="display:inline-table;border:none;height:250px;margin:0;padding:0;position:relative;visibility:visible;width:250px;background-color:transparent;"><ins id="aswift_0_anchor" style="display:block;border:none;height:250px;margin:0;padding:0;position:relative;visibility:visible;width:250px;background-color:transparent;"><iframe width="250" height="250" frameborder="0" marginwidth="0" marginheight="0" vspace="0" hspace="0" allowtransparency="true" scrolling="no" allowfullscreen="true" onload="var i=this.id,s=window.google_iframe_oncopy,H=s&amp;&amp;s.handlers,h=H&amp;&amp;H[i],w=this.contentWindow,d;try{d=w.document}catch(e){}if(h&amp;&amp;d&amp;&amp;(!d.body||!d.body.firstChild)){if(h.call){setTimeout(h,0)}else if(h.match){try{h=s.upd(h,i)}catch(e){}w.location.replace(h)}}" id="aswift_0" name="aswift_0" style="left:0;position:absolute;top:0;border:0px;width:250px;height:250px;"></iframe></ins></ins></ins>
<script>
        (adsbygoogle = window.adsbygoogle || []).push({});
        </script>
</div>
</div>
<div class="sep20"></div>
<div class="box" id="TopicsHot">
<div class="cell"><span class="fade">今日热议主题</span></div>
<div class="cell from_121655 hot_t_596977">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/zhuangzhuang1988"><img loading="lazy" src="//cdn.v2ex.com/avatar/c6af/9dc4/121655_normal.png?m=1434286696" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/596977">程序员有哪些卑微的优越感？</a>
</span>
</td>
</tr>
</tbody></table>
</div>
<div class="cell from_136952 hot_t_597067">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/huanyingch01"><img loading="lazy" src="//cdn.v2ex.com/gravatar/79383684868b6f811f60e2aeb62eab4a?s=24&amp;d=retro" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/597067">你们都是用什么编程字体的？</a>
</span>
</td>
</tr>
</tbody></table>
</div>
<div class="cell from_88902 hot_t_597017">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/devlnt"><img loading="lazy" src="//cdn.v2ex.com/avatar/8d66/1552/88902_normal.png?m=1554351937" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/597017">12 块 6T， raid50 还是 raid10 还是软 raid？</a>
</span>
</td>
</tr>
</tbody></table>
</div>
<div class="cell from_158163 hot_t_596938">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/XiLemon"><img loading="lazy" src="//cdn.v2ex.com/gravatar/84af04d7078fb2df3fcced60e64bf214?s=24&amp;d=retro" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/596938">XS Max 摔了，想问一下大家，该修还是再买</a>
</span>
</td>
</tr>
</tbody></table>
</div>
<div class="cell from_226140 hot_t_596959">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/default7"><img loading="lazy" src="//cdn.v2ex.com/gravatar/653c4003a7997bf1df7ae24f32b4aee9?s=24&amp;d=retro" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/596959">写了 11 年的代码，发现眼睛瞳孔上边有一层灰色的</a>
</span>
</td>
</tr>
</tbody></table>
</div>
<div class="cell from_267922 hot_t_597011">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/Chell"><img loading="lazy" src="//cdn.v2ex.com/avatar/d493/1929/267922_normal.png?m=1565878690" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/597011">大家认为播客在中国有前景吗</a>
</span>
</td>
</tr>
</tbody></table>
</div>
<div class="cell from_318621 hot_t_596921">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/y0bcn"><img loading="lazy" src="//cdn.v2ex.com/avatar/7d69/de62/318621_normal.png?m=1557850658" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/596921">现在 Go 环境怎么样</a>
</span>
</td>
</tr>
</tbody></table>
</div>
<div class="cell from_337112 hot_t_597064">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/Gaussen"><img loading="lazy" src="//cdn.v2ex.com/avatar/5893/73e1/337112_normal.png?m=1532920266" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/597064">请问下后端开发，笔记本用 ThinkPad x1 carbon 19 款还是 mbp 19 款？</a>
</span>
</td>
</tr>
</tbody></table>
</div>
<div class="cell from_165218 hot_t_596970">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="24" valign="middle" align="center">
<a href="/member/iamundefined"><img loading="lazy" src="//cdn.v2ex.com/gravatar/c104d3528e6dfb78023ca9c92a22e91b?s=24&amp;d=retro" class="avatar" border="0" align="default" style="max-width: 24px; max-height: 24px;"></a>
</td>
<td width="10"></td>
<td width="auto" valign="middle">
<span class="item_hot_topic_title">
<a href="/t/596970">360 年中绩效不发，请问如何维权</a>
</span>
</td>
</tr>
</tbody></table>
</div>
</div>
<div class="sep20"></div>
<div class="box">
<div class="cell"><div class="fr"></div><span class="fade">最热节点</span></div>
<div class="cell">
<a href="/go/qna" class="item_node">问与答</a><a href="/go/jobs" class="item_node">酷工作</a><a href="/go/programmer" class="item_node">程序员</a><a href="/go/share" class="item_node">分享发现</a><a href="/go/macos" class="item_node">macOS</a><a href="/go/create" class="item_node">分享创造</a><a href="/go/python" class="item_node">Python</a><a href="/go/apple" class="item_node">Apple</a><a href="/go/android" class="item_node">Android</a><a href="/go/career" class="item_node">职场话题</a><a href="/go/iphone" class="item_node">iPhone</a><a href="/go/bb" class="item_node">宽带症候群</a><a href="/go/gts" class="item_node">全球工单系统</a><a href="/go/cv" class="item_node">求职</a><a href="/go/mbp" class="item_node">MacBook Pro</a>
</div>
<div class="inner"><a href="/index.xml" target="_blank"><img src="/static/img/rss.png" align="absmiddle" border="0" style="margin-top:-3px;"></a>&nbsp; <a href="/index.xml" target="_blank">RSS</a></div>
</div>
<div class="sep20"></div>
<div class="box">
<div class="cell"><div class="fr"></div><span class="fade">最近新增节点</span></div>
<div class="inner">
<a href="/go/rss" class="item_node">RSS</a><a href="/go/jsonfeed" class="item_node">JSON Feed</a><a href="/go/vtuber" class="item_node">Virtual YouTubers</a><a href="/go/terraform" class="item_node">Terraform</a><a href="/go/remote" class="item_node">远程工作</a><a href="/go/weekly" class="item_node">写周报</a><a href="/go/cloudflare" class="item_node">Cloudflare</a><a href="/go/libra" class="item_node">Libra</a><a href="/go/typescript" class="item_node">TypeScript</a><a href="/go/tex" class="item_node">TeX</a><a href="/go/stadia" class="item_node">Stadia</a><a href="/go/apex" class="item_node">Apex Legends</a><a href="/go/bujo" class="item_node">子弹笔记</a><a href="/go/2019" class="item_node">2019</a><a href="/go/clojurescript" class="item_node">ClojureScript</a><a href="/go/leetcode" class="item_node">LeetCode</a><a href="/go/k8s" class="item_node">Kubernetes</a><a href="/go/bf" class="item_node">Battlefield 系列</a><a href="/go/bfv" class="item_node">Battlefield V</a><a href="/go/airbnb" class="item_node">Airbnb</a>
</div>
</div>
<div class="sep20"></div>
<div class="box">
<div class="cell"><span class="fade">社区运行状况</span></div>
<div class="cell">
<table cellpadding="5" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="60" align="right"><span class="gray">注册会员</span></td>
<td width="auto" align="left"><strong>439208</strong></td>
</tr>
<tr>
<td width="60" align="right"><span class="gray">主题</span></td>
<td width="auto" align="left"><strong>597120</strong></td>
</tr>
<tr>
<td width="60" align="right"><span class="gray">回复</span></td>
<td width="auto" align="left"><strong>7842402</strong></td>
</tr>
</tbody></table>
</div>
<div class="inner">
<span class="chevron">›</span> <a href="/top/rich">财富排行榜</a>
<div class="sep5"></div>
<span class="chevron">›</span> <a href="/top/player">消费排行榜</a>
</div>
</div>
<div class="sep20"></div>
</div>
<div id="Main">
<div class="sep20"></div>
<div class="box">
<div class="inner" id="Tabs">
<a href="/?tab=tech" class="tab_current">技术</a><a href="/?tab=creative" class="tab">创意</a><a href="/?tab=play" class="tab">好玩</a><a href="/?tab=apple" class="tab">Apple</a><a href="/?tab=jobs" class="tab">酷工作</a><a href="/?tab=deals" class="tab">交易</a><a href="/?tab=city" class="tab">城市</a><a href="/?tab=qna" class="tab">问与答</a><a href="/?tab=hot" class="tab">最热</a><a href="/?tab=all" class="tab">全部</a><a href="/?tab=r2" class="tab">R2</a>
</div>
<div class="cell" id="SecondaryTabs"><a href="/go/programmer">程序员</a> &nbsp; &nbsp; <a href="/go/python">Python</a> &nbsp; &nbsp; <a href="/go/idev">iDev</a> &nbsp; &nbsp; <a href="/go/android">Android</a> &nbsp; &nbsp; <a href="/go/linux">Linux</a> &nbsp; &nbsp; <a href="/go/nodejs">node.js</a> &nbsp; &nbsp; <a href="/go/cloud">云计算</a> &nbsp; &nbsp; <a href="/go/bb">宽带症候群</a></div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/ziyuzile"><img loading="lazy" src="//cdn.v2ex.com/gravatar/09abe6350e492b0954810646de3ef183?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597121#reply1">天天坐办公室，年纪轻轻就时长感觉腰酸背痛，没有精神，求一种科学的锻炼办法。</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/ziyuzile">ziyuzile</a></strong> &nbsp;•&nbsp; 1 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/xutao881">xutao881</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597121#reply1" class="count_livid">1</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/y0bcn"><img loading="lazy" src="//cdn.v2ex.com/avatar/7d69/de62/318621_normal.png?m=1557850658" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596921#reply53">现在 Go 环境怎么样</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/y0bcn">y0bcn</a></strong> &nbsp;•&nbsp; 2 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/Galileo">Galileo</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596921#reply53" class="count_livid">53</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/Gaussen"><img loading="lazy" src="//cdn.v2ex.com/avatar/5893/73e1/337112_normal.png?m=1532920266" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597064#reply50">请问下后端开发，笔记本用 ThinkPad x1 carbon 19 款还是 mbp 19 款？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"><li class="fa fa-chevron-up"></li> &nbsp;1 &nbsp;&nbsp; </div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/Gaussen">Gaussen</a></strong> &nbsp;•&nbsp; 2 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/vaccer">vaccer</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597064#reply50" class="count_livid">50</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/huanyingch01"><img loading="lazy" src="//cdn.v2ex.com/gravatar/79383684868b6f811f60e2aeb62eab4a?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597067#reply79">你们都是用什么编程字体的？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/huanyingch01">huanyingch01</a></strong> &nbsp;•&nbsp; 3 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/Rekkles">Rekkles</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597067#reply79" class="count_livid">79</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/shanlan"><img loading="lazy" src="//cdn.v2ex.com/avatar/0cd8/c8a9/360272_normal.png?m=1566063226" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596749#reply112">用机械键盘的人那么少吗？整个公司里就我一个人用。。。</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/shanlan">shanlan</a></strong> &nbsp;•&nbsp; 3 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/taxiaohaohhh">taxiaohaohhh</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
 <a href="/t/596749#reply112" class="count_livid">112</a>
</td>
</tr>
</tbody></table>
</div>
<script async="" src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<ins class="adsbygoogle" style="display: block; border-bottom: 1px solid rgb(226, 226, 226); height: 75px; width: 770px;" data-ad-format="fluid" data-ad-layout-key="-hr-19-p-2z+is" data-ad-client="ca-pub-3465543440750523" data-ad-slot="1027854874" data-adsbygoogle-status="done"><ins id="aswift_1_expand" style="display: inline-table; border: none; height: 75px; margin: 0px; padding: 0px; position: relative; visibility: visible; width: 770px; background-color: transparent;"><ins id="aswift_1_anchor" style="display: block; border: none; height: 75px; margin: 0px; padding: 0px; position: relative; visibility: visible; width: 770px; background-color: transparent; overflow: hidden;"><iframe width="770" height="75" frameborder="0" marginwidth="0" marginheight="0" vspace="0" hspace="0" allowtransparency="true" scrolling="no" allowfullscreen="true" onload="var i=this.id,s=window.google_iframe_oncopy,H=s&amp;&amp;s.handlers,h=H&amp;&amp;H[i],w=this.contentWindow,d;try{d=w.document}catch(e){}if(h&amp;&amp;d&amp;&amp;(!d.body||!d.body.firstChild)){if(h.call){setTimeout(h,0)}else if(h.match){try{h=s.upd(h,i)}catch(e){}w.location.replace(h)}}" id="aswift_1" name="aswift_1" style="left: 0px; position: absolute; top: 0px; border: 0px; width: 770px; height: 75px;"></iframe></ins></ins></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/default7"><img loading="lazy" src="//cdn.v2ex.com/gravatar/653c4003a7997bf1df7ae24f32b4aee9?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596959#reply60">写了 11 年的代码，发现眼睛瞳孔上边有一层灰色的</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/default7">default7</a></strong> &nbsp;•&nbsp; 8 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/pandaaa">pandaaa</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596959#reply60" class="count_livid">60</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/n0th1n9"><img loading="lazy" src="//cdn.v2ex.com/avatar/dda5/60ff/408193_normal.png?m=1558077013" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597083#reply1">Android 如何抓取非 http 协议流量并进行修改，如直接使用 TCP 协议的 socket 流量</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/android">Android</a> &nbsp;•&nbsp; <strong><a href="/member/n0th1n9">n0th1n9</a></strong> &nbsp;•&nbsp; 9 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/arrow8899">arrow8899</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597083#reply1" class="count_livid">1</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/johnsonshu"><img loading="lazy" src="//cdn.v2ex.com/avatar/f2cc/d331/408814_normal.png?m=1557053941" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597118#reply0">VNCServer 服务 跟 本地 GDM 抢 Display ?</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/linux">Linux</a> &nbsp;•&nbsp; <strong><a href="/member/johnsonshu">johnsonshu</a></strong> &nbsp;•&nbsp; 12 分钟前</span>
</td>
<td width="70" align="right" valign="middle">
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/Kontinue"><img loading="lazy" src="//cdn.v2ex.com/gravatar/173c992a9c5f3baec3c96967b65b2fd0?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597059#reply4">发现有道云 md 笔记备份出来也是 md 文件，赞</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/Kontinue">Kontinue</a></strong> &nbsp;•&nbsp; 15 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/OxO">OxO</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597059#reply4" class="count_livid">4</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/Ehco1996"><img loading="lazy" src="//cdn.v2ex.com/avatar/9a29/1be0/236774_normal.png?m=1550202230" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597048#reply13">我发现搞架构升级可能就意味着瞎折腾加多花钱</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/Ehco1996">Ehco1996</a></strong> &nbsp;•&nbsp; 15 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/676529483">676529483</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597048#reply13" class="count_livid">13</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/luckrill"><img loading="lazy" src="//cdn.v2ex.com/avatar/05db/77ef/426086_normal.png?m=1565482941" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597114#reply0">有好用的在线代理推荐吗？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/luckrill">luckrill</a></strong> &nbsp;•&nbsp; 14 分钟前</span>
</td>
<td width="70" align="right" valign="middle">
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/zhuangzhuang1988"><img loading="lazy" src="//cdn.v2ex.com/avatar/c6af/9dc4/121655_normal.png?m=1434286696" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596977#reply96">程序员有哪些卑微的优越感？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/zhuangzhuang1988">zhuangzhuang1988</a></strong> &nbsp;•&nbsp; 17 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/liangkang1436">liangkang1436</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596977#reply96" class="count_livid">96</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/xiaotuzi"><img loading="lazy" src="//cdn.v2ex.com/avatar/8797/d9e8/230562_normal.png?m=1544545233" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597027#reply29">关于源码保密性及仿篡改方案的思考</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/xiaotuzi">xiaotuzi</a></strong> &nbsp;•&nbsp; 18 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/codingBug">codingBug</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597027#reply29" class="count_livid">29</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/GooGee"><img loading="lazy" src="//cdn.v2ex.com/avatar/28b6/9bc9/195302_normal.png?m=1562717775" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597041#reply7">Laravel 代码生成工具（图形界面），在线生成代码</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/php">PHP</a> &nbsp;•&nbsp; <strong><a href="/member/GooGee">GooGee</a></strong> &nbsp;•&nbsp; 21 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/ben1024">ben1024</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597041#reply7" class="count_livid">7</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/amiwrong123"><img loading="lazy" src="//cdn.v2ex.com/avatar/c563/e2bd/417414_normal.png?m=1564299815" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596995#reply11">Java 动态代理 Proxy.newProxyInstance 第一个参数到底该用哪个类的类加载器啊？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/java">Java</a> &nbsp;•&nbsp; <strong><a href="/member/amiwrong123">amiwrong123</a></strong> &nbsp;•&nbsp; 28 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/coolcfan">coolcfan</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596995#reply11" class="count_livid">11</a>
</td>
</tr>
</tbody></table>
</div>
<script async="" src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<ins class="adsbygoogle" style="display: block; border-bottom: 1px solid rgb(226, 226, 226); height: 75px; width: 770px;" data-ad-format="fluid" data-ad-layout-key="-hr-19-p-2z+is" data-ad-client="ca-pub-3465543440750523" data-ad-slot="1027854874" data-adsbygoogle-status="done"><ins id="aswift_2_expand" style="display: inline-table; border: none; height: 75px; margin: 0px; padding: 0px; position: relative; visibility: visible; width: 770px; background-color: transparent;"><ins id="aswift_2_anchor" style="display: block; border: none; height: 75px; margin: 0px; padding: 0px; position: relative; visibility: visible; width: 770px; background-color: transparent; overflow: hidden;"><iframe width="770" height="75" frameborder="0" marginwidth="0" marginheight="0" vspace="0" hspace="0" allowtransparency="true" scrolling="no" allowfullscreen="true" onload="var i=this.id,s=window.google_iframe_oncopy,H=s&amp;&amp;s.handlers,h=H&amp;&amp;H[i],w=this.contentWindow,d;try{d=w.document}catch(e){}if(h&amp;&amp;d&amp;&amp;(!d.body||!d.body.firstChild)){if(h.call){setTimeout(h,0)}else if(h.match){try{h=s.upd(h,i)}catch(e){}w.location.replace(h)}}" id="aswift_2" name="aswift_2" style="left: 0px; position: absolute; top: 0px; border: 0px; width: 770px; height: 75px;"></iframe></ins></ins></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>

<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/songdg"><img loading="lazy" src="//cdn.v2ex.com/gravatar/73587338add07b89dd761ff0c4292cc5?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597055#reply14">新版的微信太耗电了</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/songdg">songdg</a></strong> &nbsp;•&nbsp; 29 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/Lin0936">Lin0936</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597055#reply14" class="count_livid">14</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/paparika"><img loading="lazy" src="//cdn.v2ex.com/gravatar/c6fcb55ffd0d80bbde46c667d45f9237?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597058#reply7">求助一个无头绪排查的段错误问题</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/c">C/C++/Obj-C</a> &nbsp;•&nbsp; <strong><a href="/member/paparika">paparika</a></strong> &nbsp;•&nbsp; 45 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/augustheart">augustheart</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597058#reply7" class="count_livid">7</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/killersdz"><img loading="lazy" src="//cdn.v2ex.com/gravatar/68c97096c1832549088895969f74727b?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597065#reply2">请问 xadmin 本地测试的时候导出 excel 功能正常，放到服务器生产环境下报 500 错误是什么原因，，是 xadmin 的 bug，还是没配置好？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/python">Python</a> &nbsp;•&nbsp; <strong><a href="/member/killersdz">killersdz</a></strong> &nbsp;•&nbsp; 46 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/xpresslink">xpresslink</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597065#reply2" class="count_livid">2</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/douleL"><img loading="lazy" src="//cdn.v2ex.com/gravatar/790897111f59134e8f00042be37be5db?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597061#reply2">Java 如何实现 无卡状态读取 sim 卡（手机卡）的短信 突发奇想 有老铁解答吗 ，提供解决方案</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/java">Java</a> &nbsp;•&nbsp; <strong><a href="/member/douleL">douleL</a></strong> &nbsp;•&nbsp; 1 小时 25 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/douleL">douleL</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597061#reply2" class="count_livid">2</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/PainfulJoe"><img loading="lazy" src="//cdn.v2ex.com/gravatar/3522c291fd96a36c5729ebe1fd66395f?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597049#reply1">谁在生产环境用过 hyperf，感觉如何？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/php">PHP</a> &nbsp;•&nbsp; <strong><a href="/member/PainfulJoe">PainfulJoe</a></strong> &nbsp;•&nbsp; 2 小时 3 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/xnode">xnode</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597049#reply1" class="count_livid">1</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/maruize321"><img loading="lazy" src="//cdn.v2ex.com/gravatar/de4610dd7e1615728c6b6fa1253e6c3b?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597123#reply0">第一次在 vue 中使用 websocket+ProtoBuf 有什么需要注意的吗</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/maruize321">maruize321</a></strong> &nbsp;•&nbsp; 2 分钟前</span>
</td>
<td width="70" align="right" valign="middle">
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/BuilderQiu"><img loading="lazy" src="//cdn.v2ex.com/avatar/f6c3/17a5/118629_normal.png?m=1432518741" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596867#reply40">GooglePlay 美区又可以领 2$优惠券</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/android">Android</a> &nbsp;•&nbsp; <strong><a href="/member/BuilderQiu">BuilderQiu</a></strong> &nbsp;•&nbsp; 9 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/sleepm">sleepm</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596867#reply40" class="count_livid">40</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/johnsonshu"><img loading="lazy" src="//cdn.v2ex.com/avatar/f2cc/d331/408814_normal.png?m=1557053941" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/597013#reply6">怎么查找 resource ID 对应的文件？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/android">Android</a> &nbsp;•&nbsp; <strong><a href="/member/johnsonshu">johnsonshu</a></strong> &nbsp;•&nbsp; 11 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/johnsonshu">johnsonshu</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/597013#reply6" class="count_livid">6</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/lsls931011"><img loading="lazy" src="//cdn.v2ex.com/gravatar/e1dd21e86c5328d61dc8792e0c8cda3c?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596944#reply26">YourNovel-基于 Golang 的开源小说搜索引擎&amp;免费小说阅读网站 http://www.yournovel.cn</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/lsls931011">lsls931011</a></strong> &nbsp;•&nbsp; 1 小时 16 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/usslss">usslss</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596944#reply26" class="count_livid">26</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/Colorful"><img loading="lazy" src="//cdn.v2ex.com/gravatar/fbd0e1ac180739637d72ae35dc8895c6?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596984#reply9">请教下数据库取出乱码，怎么解决？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/mysql">MySQL</a> &nbsp;•&nbsp; <strong><a href="/member/Colorful">Colorful</a></strong> &nbsp;•&nbsp; 1 小时 30 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/xmai">xmai</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596984#reply9" class="count_livid">9</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/luckrill"><img loading="lazy" src="//cdn.v2ex.com/avatar/05db/77ef/426086_normal.png?m=1565482941" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596964#reply6">you-get 真是太棒了</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"><li class="fa fa-chevron-up"></li> &nbsp;1 &nbsp;&nbsp; </div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/luckrill">luckrill</a></strong> &nbsp;•&nbsp; 17 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/luckrill">luckrill</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596964#reply6" class="count_livid">6</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/Aidea"><img loading="lazy" src="//cdn.v2ex.com/avatar/66cc/2a4a/133581_normal.png?m=1482463420" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596931#reply18">有没有不刷新页面加载复杂 div 的方法？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/Aidea">Aidea</a></strong> &nbsp;•&nbsp; 8 小时 16 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/dartabe">dartabe</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596931#reply18" class="count_livid">18</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/0o0O0o0O0o"><img loading="lazy" src="//cdn.v2ex.com/gravatar/3c8d28727aa2095657894289acfc9083?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596879#reply18">Go 模块代理超大型库初始化速度实测： goproxy.cn vs goproxy.io</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"><li class="fa fa-chevron-up"></li> &nbsp;1 &nbsp;&nbsp; </div><a class="node" href="/go/go">Go</a> &nbsp;•&nbsp; <strong><a href="/member/0o0O0o0O0o">0o0O0o0O0o</a></strong> &nbsp;•&nbsp; 12 小时 12 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/aofei">aofei</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596879#reply18" class="count_livid">18</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/zhoumouren"><img loading="lazy" src="//cdn.v2ex.com/avatar/6597/2c11/360466_normal.png?m=1566955038" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596708#reply43">该不该辞职呢</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/zhoumouren">zhoumouren</a></strong> &nbsp;•&nbsp; 18 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/DamonLin">DamonLin</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596708#reply43" class="count_livid">43</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/DamienS"><img loading="lazy" src="//cdn.v2ex.com/avatar/4ef0/bf80/108661_normal.png?m=1427911692" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596958#reply4">django 有没啥好的基于属性的权限验证（Attribute Based Access Control）的开源项目</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/python">Python</a> &nbsp;•&nbsp; <strong><a href="/member/DamienS">DamienS</a></strong> &nbsp;•&nbsp; 32 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/xpresslink">xpresslink</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596958#reply4" class="count_livid">4</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/guanzhangzhang"><img loading="lazy" src="//cdn.v2ex.com/avatar/6e8b/5df0/327727_normal.png?m=1567327193" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596953#reply10">[json 的验证]后端的 json 验证</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"><li class="fa fa-chevron-up"></li> &nbsp;1 &nbsp;&nbsp; </div><a class="node" href="/go/go">Go</a> &nbsp;•&nbsp; <strong><a href="/member/guanzhangzhang">guanzhangzhang</a></strong> &nbsp;•&nbsp; 14 小时 1 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/xzc19970719">xzc19970719</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596953#reply10" class="count_livid">10</a>
</td>
</tr>
</tbody></table>
</div>
<script async="" src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<ins class="adsbygoogle" style="display: block; border-bottom: 1px solid rgb(226, 226, 226); height: 94px;" data-ad-format="fluid" data-ad-layout-key="-hr-19-p-2z+is" data-ad-client="ca-pub-3465543440750523" data-ad-slot="1027854874" data-adsbygoogle-status="done"><ins id="aswift_3_expand" style="display:inline-table;border:none;height:94px;margin:0;padding:0;position:relative;visibility:visible;width:770px;background-color:transparent;"><ins id="aswift_3_anchor" style="display:block;border:none;height:94px;margin:0;padding:0;position:relative;visibility:visible;width:770px;background-color:transparent;"><iframe width="770" height="94" frameborder="0" marginwidth="0" marginheight="0" vspace="0" hspace="0" allowtransparency="true" scrolling="no" allowfullscreen="true" onload="var i=this.id,s=window.google_iframe_oncopy,H=s&amp;&amp;s.handlers,h=H&amp;&amp;H[i],w=this.contentWindow,d;try{d=w.document}catch(e){}if(h&amp;&amp;d&amp;&amp;(!d.body||!d.body.firstChild)){if(h.call){setTimeout(h,0)}else if(h.match){try{h=s.upd(h,i)}catch(e){}w.location.replace(h)}}" id="aswift_3" name="aswift_3" style="left:0;position:absolute;top:0;border:0px;width:770px;height:94px;"></iframe></ins></ins></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/DongDongXie"><img loading="lazy" src="//cdn.v2ex.com/gravatar/cf9066a688ba15afbe91f71caafcd8e3?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596949#reply3">多条指令后台执行怎么同时写一个数组</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/DongDongXie">DongDongXie</a></strong> &nbsp;•&nbsp; 14 小时 9 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/jinliming2">jinliming2</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596949#reply3" class="count_livid">3</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/kraits"><img loading="lazy" src="//cdn.v2ex.com/gravatar/953de4234df55c1c973abb1c1588dc2e?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596811#reply26">为 typecho 写了一个 android 客户端</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"><li class="fa fa-chevron-up"></li> &nbsp;5 &nbsp;&nbsp; </div><a class="node" href="/go/android">Android</a> &nbsp;•&nbsp; <strong><a href="/member/kraits">kraits</a></strong> &nbsp;•&nbsp; 1 小时 43 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/hitaoguo">hitaoguo</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596811#reply26" class="count_livid">26</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/raftPaxos"><img loading="lazy" src="//cdn.v2ex.com/gravatar/a6dc8b1a1a407f5e3189517b78361765?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596930#reply5">如何集中 Python 脚本产生的日志信息</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/raftPaxos">raftPaxos</a></strong> &nbsp;•&nbsp; 1 小时 3 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/arrow8899">arrow8899</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596930#reply5" class="count_livid">5</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/imdong"><img loading="lazy" src="//cdn.v2ex.com/avatar/2515/bbe5/247628_normal.png?m=1509541357" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596794#reply28">前两天有 V 友问一个加密的 JS 怎么解密，于是今天脱壳？解密的脚本出来了。</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/imdong">imdong</a></strong> &nbsp;•&nbsp; 1 小时 50 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/hakono">hakono</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596794#reply28" class="count_livid">28</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/spikedingo"><img loading="lazy" src="//cdn.v2ex.com/gravatar/3875f254bd55d885976e51c8f0643573?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596861#reply17">如何让来电的固话显示在 PC web 系统上，甚至进行录音功能</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/spikedingo">spikedingo</a></strong> &nbsp;•&nbsp; 13 小时 12 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/Ymob">Ymob</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596861#reply17" class="count_livid">17</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/able"><img loading="lazy" src="//cdn.v2ex.com/gravatar/e0e2302ecfe5f3b077c180388bb0d9e0?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596847#reply7">阿里云活动页价格标错了，¥230/年 起，应该是每月。</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/cloud">云计算</a> &nbsp;•&nbsp; <strong><a href="/member/able">able</a></strong> &nbsp;•&nbsp; 6 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/codingBug">codingBug</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596847#reply7" class="count_livid">7</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/EulerChen"><img loading="lazy" src="//cdn.v2ex.com/avatar/8dca/0344/367423_normal.png?m=1557395629" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596955#reply2">Golang 如何写出同时包含字母和数字的正则?</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/go">Go</a> &nbsp;•&nbsp; <strong><a href="/member/EulerChen">EulerChen</a></strong> &nbsp;•&nbsp; 14 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/AddictX">AddictX</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596955#reply2" class="count_livid">2</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/Alfons"><img loading="lazy" src="//cdn.v2ex.com/gravatar/ae0e7c1ab9e9bfc48f5d84c4904c4f20?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596797#reply9">国庆+十日+海南+环岛+骑行</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/Alfons">Alfons</a></strong> &nbsp;•&nbsp; 1 小时 34 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/maichael">maichael</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596797#reply9" class="count_livid">9</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/luckrill"><img loading="lazy" src="//cdn.v2ex.com/avatar/05db/77ef/426086_normal.png?m=1565482941" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596566#reply74">点评几个搜索引擎</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/luckrill">luckrill</a></strong> &nbsp;•&nbsp; 4 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/luckrill">luckrill</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596566#reply74" class="count_livid">74</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/shanlan"><img loading="lazy" src="//cdn.v2ex.com/avatar/0cd8/c8a9/360272_normal.png?m=1566063226" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596562#reply172">你旧笔记本用几年了？还会一直用下去吗？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"><li class="fa fa-chevron-up"></li> &nbsp;1 &nbsp;&nbsp; </div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/shanlan">shanlan</a></strong> &nbsp;•&nbsp; 3 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/Rlyi">Rlyi</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596562#reply172" class="count_livid">172</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/iVampireSP"><img loading="lazy" src="//cdn.v2ex.com/gravatar/820794b2ee799e6bd2396fcd712a5ea8?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596721#reply20">求助一下。这乱码。。我有点迷。</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/linux">Linux</a> &nbsp;•&nbsp; <strong><a href="/member/iVampireSP">iVampireSP</a></strong> &nbsp;•&nbsp; 18 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/LFUNWF">LFUNWF</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596721#reply20" class="count_livid">20</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/neteroster"><img loading="lazy" src="//cdn.v2ex.com/avatar/7f29/11c5/191331_normal.png?m=1565838181" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596816#reply10">在 Windows 上使用 DNS over HTTPS</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/dns">DNS</a> &nbsp;•&nbsp; <strong><a href="/member/neteroster">neteroster</a></strong> &nbsp;•&nbsp; 2 小时 19 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/testcaoy7">testcaoy7</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596816#reply10" class="count_livid">10</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/badtypea"><img loading="lazy" src="//cdn.v2ex.com/avatar/9657/7508/419941_normal.png?m=1567310347" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596899#reply1">bootstrap-flask 为什么要这样设计</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/python">Python</a> &nbsp;•&nbsp; <strong><a href="/member/badtypea">badtypea</a></strong> &nbsp;•&nbsp; 13 小时 9 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/sazima">sazima</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596899#reply1" class="count_livid">1</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/johnsonshu"><img loading="lazy" src="//cdn.v2ex.com/avatar/f2cc/d331/408814_normal.png?m=1557053941" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596740#reply27">国内用 android studio?</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/android">Android</a> &nbsp;•&nbsp; <strong><a href="/member/johnsonshu">johnsonshu</a></strong> &nbsp;•&nbsp; 2 小时 11 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/youxiachai">youxiachai</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596740#reply27" class="count_livid">27</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/dafsic"><img loading="lazy" src="//cdn.v2ex.com/gravatar/4384f0ec6f56d615f276328c9a6b019c?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596606#reply15">golang 多协程操作同一个 map 的方案哪个好 ??</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/go">Go</a> &nbsp;•&nbsp; <strong><a href="/member/dafsic">dafsic</a></strong> &nbsp;•&nbsp; 5 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/tourist2018">tourist2018</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596606#reply15" class="count_livid">15</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
 <tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/hkitdog"><img loading="lazy" src="//cdn.v2ex.com/gravatar/1f783673c48e1dca5f4483fb69b305a9?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596878#reply4">Vue 如何实现画中画功能？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/vue">Vue.js</a> &nbsp;•&nbsp; <strong><a href="/member/hkitdog">hkitdog</a></strong> &nbsp;•&nbsp; 23 小时 2 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/sunjourney">sunjourney</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596878#reply4" class="count_livid">4</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/shenshenleilei"><img loading="lazy" src="//cdn.v2ex.com/gravatar/2f88ef9a2592e386da3cc1cb82161558?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596863#reply1">visual studio 装了 visual assist 插件运行速度非常慢？怎么办？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/shenshenleilei">shenshenleilei</a></strong> &nbsp;•&nbsp; 1 天前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/Athrob">Athrob</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596863#reply1" class="count_livid">1</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/swsh007"><img loading="lazy" src="//cdn.v2ex.com/gravatar/2d9c376b16df70793339f774363a88b9?s=48&amp;d=retro" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596834#reply6">原来有个 Mulan PSL v1 License</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"></div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/swsh007">swsh007</a></strong> &nbsp;•&nbsp; 20 小时 55 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/realpg">realpg</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596834#reply6" class="count_livid">6</a>
</td>
</tr>
</tbody></table>
</div>
<div class="cell item" style="">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><a href="/member/Allianzcortex"><img loading="lazy" src="//cdn.v2ex.com/avatar/6625/804b/132111_normal.png?m=1517838921" class="avatar" border="0" align="default"></a></td>
<td width="10"></td>
<td width="auto" valign="middle"><span class="item_title"><a href="/t/596850#reply5">[不可能事件求助帖] 之前可以通过单元测试的 commit 现在无法通过？</a></span>
<div class="sep5"></div>
<span class="topic_info"><div class="votes"><li class="fa fa-chevron-up"></li> &nbsp;1 &nbsp;&nbsp; </div><a class="node" href="/go/programmer">程序员</a> &nbsp;•&nbsp; <strong><a href="/member/Allianzcortex">Allianzcortex</a></strong> &nbsp;•&nbsp; 18 小时 47 分钟前 &nbsp;•&nbsp; 最后回复来自 <strong><a href="/member/poplar50">poplar50</a></strong></span>
</td>
<td width="70" align="right" valign="middle">
<a href="/t/596850#reply5" class="count_livid">5</a>
</td>
</tr>
</tbody></table>
</div>
<script async="" src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<ins class="adsbygoogle" style="display: block; border-bottom: 1px solid rgb(226, 226, 226); height: 94px;" data-ad-format="fluid" data-ad-layout-key="-hr-19-p-2z+is" data-ad-client="ca-pub-3465543440750523" data-ad-slot="1027854874" data-adsbygoogle-status="done"><ins id="aswift_4_expand" style="display:inline-table;border:none;height:94px;margin:0;padding:0;position:relative;visibility:visible;width:770px;background-color:transparent;"><ins id="aswift_4_anchor" style="display:block;border:none;height:94px;margin:0;padding:0;position:relative;visibility:visible;width:770px;background-color:transparent;"><iframe width="770" height="94" frameborder="0" marginwidth="0" marginheight="0" vspace="0" hspace="0" allowtransparency="true" scrolling="no" allowfullscreen="true" onload="var i=this.id,s=window.google_iframe_oncopy,H=s&amp;&amp;s.handlers,h=H&amp;&amp;H[i],w=this.contentWindow,d;try{d=w.document}catch(e){}if(h&amp;&amp;d&amp;&amp;(!d.body||!d.body.firstChild)){if(h.call){setTimeout(h,0)}else if(h.match){try{h=s.upd(h,i)}catch(e){}w.location.replace(h)}}" id="aswift_4" name="aswift_4" style="left:0;position:absolute;top:0;border:0px;width:770px;height:94px;"></iframe></ins></ins></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>
<div class="inner"><div class="fr"><a href="/feed/tab/tech.xml" target="_blank"><img src="/static/img/rss.png" align="absmiddle" style="margin-top:-3px;" border="0"></a>&nbsp; <a href="/feed/tab/tech.xml" target="_blank">通过 Atom Feed 订阅</a></div>
<span class="chevron">→</span> <a href="/recent">更多新主题</a>
</div>
</div>
<div class="sep20"></div>
<div class="box">
<div class="cell"><div class="fr"><a href="/planes">浏览全部节点</a></div><span class="fade"><strong>V2EX</strong> / 节点导航</span></div>
<div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">分享与探索</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/qna" style="font-size: 14px;">问与答</a>&nbsp; &nbsp; <a href="/go/share" style="font-size: 14px;">分享发现</a>&nbsp; &nbsp; <a href="/go/create" style="font-size: 14px;">分享创造</a>&nbsp; &nbsp; <a href="/go/in" style="font-size: 14px;">分享邀请码</a>&nbsp; &nbsp; <a href="/go/ideas" style="font-size: 14px;">奇思妙想</a>&nbsp; &nbsp; <a href="/go/autistic" style="font-size: 14px;">自言自语</a>&nbsp; &nbsp; <a href="/go/random" style="font-size: 14px;">随想</a>&nbsp; &nbsp; <a href="/go/design" style="font-size: 14px;">设计</a>&nbsp; &nbsp; <a href="/go/blog" style="font-size: 14px;">Blog</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">V2EX</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/v2ex" style="font-size: 14px;">V2EX</a>&nbsp; &nbsp; <a href="/go/dns" style="font-size: 14px;">DNS</a>&nbsp; &nbsp; <a href="/go/feedback" style="font-size: 14px;">反馈</a>&nbsp; &nbsp; <a href="/go/babel" style="font-size: 14px;">Project Babel</a>&nbsp; &nbsp; <a href="/go/guide" style="font-size: 14px;">使用指南</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">iOS</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/idev" style="font-size: 14px;">iDev</a>&nbsp; &nbsp; <a href="/go/icode" style="font-size: 14px;">iCode</a>&nbsp; &nbsp; <a href="/go/imarketing" style="font-size: 14px;">iMarketing</a>&nbsp; &nbsp; <a href="/go/iad" style="font-size: 14px;">iAd</a>&nbsp; &nbsp; <a href="/go/itransfer" style="font-size: 14px;">iTransfer</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">Geek</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/programmer" style="font-size: 14px;">程序员</a>&nbsp; &nbsp; <a href="/go/python" style="font-size: 14px;">Python</a>&nbsp; &nbsp; <a href="/go/android" style="font-size: 14px;">Android</a>&nbsp; &nbsp; <a href="/go/bb" style="font-size: 14px;">宽带症候群</a>&nbsp; &nbsp; <a href="/go/linux" style="font-size: 14px;">Linux</a>&nbsp; &nbsp; <a href="/go/php" style="font-size: 14px;">PHP</a>&nbsp; &nbsp; <a href="/go/cloud" style="font-size: 14px;">云计算</a>&nbsp; &nbsp; <a href="/go/outsourcing" style="font-size: 14px;">外包</a>&nbsp; &nbsp; <a href="/go/hardware" style="font-size: 14px;">硬件</a>&nbsp; &nbsp; <a href="/go/java" style="font-size: 14px;">Java</a>&nbsp; &nbsp; <a href="/go/nodejs" style="font-size: 14px;">Node.js</a>&nbsp; &nbsp; <a href="/go/server" style="font-size: 14px;">服务器</a>&nbsp; &nbsp; <a href="/go/bitcoin" style="font-size: 14px;">Bitcoin</a>&nbsp; &nbsp; <a href="/go/mysql" style="font-size: 14px;">MySQL</a>&nbsp; &nbsp; <a href="/go/programming" style="font-size: 14px;">编程</a>&nbsp; &nbsp; <a href="/go/car" style="font-size: 14px;">汽车</a>&nbsp; &nbsp; <a href="/go/docker" style="font-size: 14px;">Docker</a>&nbsp; &nbsp; <a href="/go/linode" style="font-size: 14px;">Linode</a>&nbsp; &nbsp; <a href="/go/designer" style="font-size: 14px;">设计师</a>&nbsp; &nbsp; <a href="/go/kindle" style="font-size: 14px;">Kindle</a>&nbsp; &nbsp; <a href="/go/markdown" style="font-size: 14px;">Markdown</a>&nbsp; &nbsp; <a href="/go/mongodb" style="font-size: 14px;">MongoDB</a>&nbsp; &nbsp; <a href="/go/redis" style="font-size: 14px;">Redis</a>&nbsp; &nbsp; <a href="/go/minecraft" style="font-size: 14px;">Minecraft</a>&nbsp; &nbsp; <a href="/go/tornado" style="font-size: 14px;">Tornado</a>&nbsp; &nbsp; <a href="/go/typography" style="font-size: 14px;">字体排印</a>&nbsp; &nbsp; <a href="/go/ror" style="font-size: 14px;">Ruby on Rails</a>&nbsp; &nbsp; <a href="/go/business" style="font-size: 14px;">商业模式</a>&nbsp; &nbsp; <a href="/go/ruby" style="font-size: 14px;">Ruby</a>&nbsp; &nbsp; <a href="/go/math" style="font-size: 14px;">数学</a>&nbsp; &nbsp; <a href="/go/photoshop" style="font-size: 14px;">Photoshop</a>&nbsp; &nbsp; <a href="/go/sony" style="font-size: 14px;">SONY</a>&nbsp; &nbsp; <a href="/go/csharp" style="font-size: 14px;">C#</a>&nbsp; &nbsp; <a href="/go/amazon" style="font-size: 14px;">Amazon</a>&nbsp; &nbsp; <a href="/go/nlp" style="font-size: 14px;">自然语言处理</a>&nbsp; &nbsp; <a href="/go/lego" style="font-size: 14px;">LEGO</a>&nbsp; &nbsp; <a href="/go/leetcode" style="font-size: 14px;">LeetCode</a>&nbsp; &nbsp; <a href="/go/ev" style="font-size: 14px;">电动汽车</a>&nbsp; &nbsp; <a href="/go/serverless" style="font-size: 14px;">Serverless</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">游戏</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/games" style="font-size: 14px;">游戏</a>&nbsp; &nbsp; <a href="/go/steam" style="font-size: 14px;">Steam</a>&nbsp; &nbsp; <a href="/go/ps4" style="font-size: 14px;">PlayStation 4</a>&nbsp; &nbsp; <a href="/go/lol" style="font-size: 14px;">英雄联盟</a>&nbsp; &nbsp; <a href="/go/igame" style="font-size: 14px;">iGame</a>&nbsp; &nbsp; <a href="/go/switch" style="font-size: 14px;">Nintendo Switch</a>&nbsp; &nbsp; <a href="/go/sc2" style="font-size: 14px;">StarCraft 2</a>&nbsp; &nbsp; <a href="/go/bf3" style="font-size: 14px;">Battlefield 3</a>&nbsp; &nbsp; <a href="/go/wow" style="font-size: 14px;">World of Warcraft</a>&nbsp; &nbsp; <a href="/go/5v5" style="font-size: 14px;">王者荣耀</a>&nbsp; &nbsp; <a href="/go/eve" style="font-size: 14px;">EVE</a>&nbsp; &nbsp; <a href="/go/bf4" style="font-size: 14px;">Battlefield 4</a>&nbsp; &nbsp; <a href="/go/gt" style="font-size: 14px;">Gran Turismo</a>&nbsp; &nbsp; <a href="/go/wiiu" style="font-size: 14px;">Wii U</a>&nbsp; &nbsp; <a href="/go/bfv" style="font-size: 14px;">Battlefield V</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">Apple</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/macos" style="font-size: 14px;">macOS</a>&nbsp; &nbsp; <a href="/go/iphone" style="font-size: 14px;">iPhone</a>&nbsp; &nbsp; <a href="/go/mbp" style="font-size: 14px;">MacBook Pro</a>&nbsp; &nbsp; <a href="/go/ipad" style="font-size: 14px;">iPad</a>&nbsp; &nbsp; <a href="/go/macbook" style="font-size: 14px;">MacBook</a>&nbsp; &nbsp; <a href="/go/accessory" style="font-size: 14px;">配件</a>&nbsp; &nbsp; <a href="/go/mba" style="font-size: 14px;">MacBook Air</a>&nbsp; &nbsp; <a href="/go/imac" style="font-size: 14px;">iMac</a>&nbsp; &nbsp; <a href="/go/macmini" style="font-size: 14px;">Mac mini</a>&nbsp; &nbsp; <a href="/go/macpro" style="font-size: 14px;">Mac Pro</a>&nbsp; &nbsp; <a href="/go/ipod" style="font-size: 14px;">iPod</a>&nbsp; &nbsp; <a href="/go/mobileme" style="font-size: 14px;">MobileMe</a>&nbsp; &nbsp; <a href="/go/iwork" style="font-size: 14px;">iWork</a>&nbsp; &nbsp; <a href="/go/ilife" style="font-size: 14px;">iLife</a>&nbsp; &nbsp; <a href="/go/garageband" style="font-size: 14px;">GarageBand</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">生活</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/all4all" style="font-size: 14px;">二手交易</a>&nbsp; &nbsp; <a href="/go/jobs" style="font-size: 14px;">酷工作</a>&nbsp; &nbsp; <a href="/go/afterdark" style="font-size: 14px;">天黑以后</a>&nbsp; &nbsp; <a href="/go/free" style="font-size: 14px;">免费赠送</a>&nbsp; &nbsp; <a href="/go/music" style="font-size: 14px;">音乐</a>&nbsp; &nbsp; <a href="/go/movie" style="font-size: 14px;">电影</a>&nbsp; &nbsp; <a href="/go/exchange" style="font-size: 14px;">物物交换</a>&nbsp; &nbsp; <a href="/go/tuan" style="font-size: 14px;">团购</a>&nbsp; &nbsp; <a href="/go/tv" style="font-size: 14px;">剧集</a>&nbsp; &nbsp; <a href="/go/invest" style="font-size: 14px;">投资</a>&nbsp; &nbsp; <a href="/go/creditcard" style="font-size: 14px;">信用卡</a>&nbsp; &nbsp; <a href="/go/travel" style="font-size: 14px;">旅行</a>&nbsp; &nbsp; <a href="/go/taste" style="font-size: 14px;">美酒与美食</a>&nbsp; &nbsp; <a href="/go/reading" style="font-size: 14px;">阅读</a>&nbsp; &nbsp; <a href="/go/photograph" style="font-size: 14px;">摄影</a>&nbsp; &nbsp; <a href="/go/pet" style="font-size: 14px;">宠物</a>&nbsp; &nbsp; <a href="/go/baby" style="font-size: 14px;">Baby</a>&nbsp; &nbsp; <a href="/go/soccer" style="font-size: 14px;">绿茵场</a>&nbsp; &nbsp; <a href="/go/coffee" style="font-size: 14px;">咖啡</a>&nbsp; &nbsp; <a href="/go/diary" style="font-size: 14px;">日记</a>&nbsp; &nbsp; <a href="/go/love" style="font-size: 14px;">非诚勿扰</a>&nbsp; &nbsp; <a href="/go/lohas" style="font-size: 14px;">乐活</a>&nbsp; &nbsp; <a href="/go/bike" style="font-size: 14px;">骑行</a>&nbsp; &nbsp; <a href="/go/plant" style="font-size: 14px;">植物</a>&nbsp; &nbsp; <a href="/go/mushroom" style="font-size: 14px;">蘑菇</a>&nbsp; &nbsp; <a href="/go/mileage" style="font-size: 14px;">行程控</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">Internet</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/google" style="font-size: 14px;">Google</a>&nbsp; &nbsp; <a href="/go/twitter" style="font-size: 14px;">Twitter</a>&nbsp; &nbsp; <a href="/go/coding" style="font-size: 14px;">Coding</a>&nbsp; &nbsp; <a href="/go/facebook" style="font-size: 14px;">Facebook</a>&nbsp; &nbsp; <a href="/go/wikipedia" style="font-size: 14px;">Wikipedia</a>&nbsp; &nbsp; <a href="/go/reddit" style="font-size: 14px;">reddit</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="cell"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">城市</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/beijing" style="font-size: 14px;">北京</a>&nbsp; &nbsp; <a href="/go/shanghai" style="font-size: 14px;">上海</a>&nbsp; &nbsp; <a href="/go/shenzhen" style="font-size: 14px;">深圳</a>&nbsp; &nbsp; <a href="/go/hangzhou" style="font-size: 14px;">杭州</a>&nbsp; &nbsp; <a href="/go/chengdu" style="font-size: 14px;">成都</a>&nbsp; &nbsp; <a href="/go/guangzhou" style="font-size: 14px;">广州</a>&nbsp; &nbsp; <a href="/go/wuhan" style="font-size: 14px;">武汉</a>&nbsp; &nbsp; <a href="/go/kunming" style="font-size: 14px;">昆明</a>&nbsp; &nbsp; <a href="/go/tianjin" style="font-size: 14px;">天津</a>&nbsp; &nbsp; <a href="/go/qingdao" style="font-size: 14px;">青岛</a>&nbsp; &nbsp; <a href="/go/nyc" style="font-size: 14px;">New York</a>&nbsp; &nbsp; <a href="/go/sanfrancisco" style="font-size: 14px;">San Francisco</a>&nbsp; &nbsp; <a href="/go/la" style="font-size: 14px;">Los Angeles</a>&nbsp; &nbsp; <a href="/go/boston" style="font-size: 14px;">Boston</a>&nbsp; &nbsp; </td></tr></tbody></table></div><div class="inner"><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td align="right" width="60"><span class="fade">品牌</span></td><td style="line-height: 200%; padding-left: 10px; word-break: keep-all;"><a href="/go/uniqlo" style="font-size: 14px;">UNIQLO</a>&nbsp; &nbsp; <a href="/go/lamy" style="font-size: 14px;">Lamy</a>&nbsp; &nbsp; <a href="/go/ikea" style="font-size: 14px;">宜家</a>&nbsp; &nbsp; <a href="/go/muji" style="font-size: 14px;">无印良品</a>&nbsp; &nbsp; <a href="/go/gap" style="font-size: 14px;">Gap</a>&nbsp; &nbsp; <a href="/go/nike" style="font-size: 14px;">Nike</a>&nbsp; &nbsp; <a href="/go/moleskine" style="font-size: 14px;">Moleskine</a>&nbsp; &nbsp; <a href="/go/adidas" style="font-size: 14px;">Adidas</a>&nbsp; &nbsp; <a href="/go/gstar" style="font-size: 14px;">G-Star</a>&nbsp; &nbsp; </td></tr></tbody></table></div>
</div>
</div>
</div>
<div class="c"></div>
<div class="sep20"></div>
</div>
<div id="Bottom">
<div class="content">
<div class="inner">
<div class="sep10"></div>
<div class="fr">
<a href="https://www.digitalocean.com/?refcode=1b51f1a7651d" target="_blank"><div id="DigitalOcean"></div></a>
</div>
<strong><a href="/about" class="dark" target="_self">关于</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/faq" class="dark" target="_self">FAQ</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/p/7v9TEc53" class="dark" target="_self">API</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/mission" class="dark" target="_self">我们的愿景</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/advertise" class="dark" target="_self">广告投放</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/advertise/2017.html" class="dark" target="_self">感谢</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/tools" class="dark" target="_self">实用小工具</a> &nbsp; <span class="snow">·</span> &nbsp; 3938 人在线</strong> &nbsp; <span class="fade">最高记录 5043</span> &nbsp; <span class="snow">·</span> &nbsp; <a href="/select/language" class="f11"><img src="/static/img/language.png?v=6a5cfa731dc71a3769f6daace6784739" width="16" align="absmiddle" id="ico-select-language"> &nbsp; Select Language</a>
<div class="sep20"></div>
创意工作者们的社区
<div class="sep5"></div>
World is powered by solitude
<div class="sep20"></div>
<span class="small fade">VERSION: 3.9.8.3 · 7ms · UTC 03:22 · PVG 11:22 · LAX 20:22 · JFK 23:22<br>♥ Do have faith in what you're doing.</span>
<div class="sep10"></div>
</div>
</div>
</div>
			</body>
			`

var html = `<body>
<div id="Top">
<div class="content">
<div style="padding-top: 6px;">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="110" align="left"><a href="/" name="top" title="way to explore"><div id="Logo"></div></a></td>
<td width="auto" align="left">
<div id="Search"><form action="https://www.google.com" onsubmit="return dispatch()" target="_blank"><div id="qbar"><input type="text" maxlength="40" name="q" id="q" value="" onfocus="$('#qbar').addClass('qbar_focus')" onblur="$('#qbar').removeClass('qbar_focus')"></div></form></div>
</td>
<td width="570" align="right" style="padding-top: 2px;"><a href="/" class="top">首页</a>&nbsp;&nbsp;&nbsp;<a href="/signup" class="top">注册</a>&nbsp;&nbsp;&nbsp;<a href="/signin" class="top">登录</a></td>
</tr>
</tbody></table>
</div>
</div>
</div>
<div id="Wrapper">
<div class="content">
<div id="Leftbar"></div>
<div id="Rightbar">
<div class="sep20"></div>
<div class="box">
<div class="cell">
<strong>V2EX = way to explore</strong>
<div class="sep5"></div>
<span class="fade">V2EX 是一个关于分享和探索的地方</span>
</div>
<div class="inner">
<div class="sep5"></div>
<div align="center"><a href="/signup" class="super normal button">现在注册</a>
<div class="sep5"></div>
<div class="sep10"></div>
已注册用户请 &nbsp;<a href="/signin">登录</a></div>
</div>
</div>
<div class="sep20"></div>
<div class="box">
<div class="inner">
<div class="sidebar_units">
<a href="https://www.openinstall.io/?from=v2ex" target="_blank" rel="nofollow"><img src="//cdn.v2ex.com/assets/sidebar/openinstall_20190724.gif" border="0" width="100%" alt="openinstall.io"></a>
<div class="sep10"></div>
<a href="http://www.huodongxing.com/event/5507844955900?qd=v2exweb" target="_blank"><img src="//cdn.v2ex.com/assets/sidebar/geekbang_20190828.jpg" border="0" width="100%" alt="Geekbang"></a>
<div class="sep10"></div>
<a href="http://www.yundun.com?hmsr=V2EX&amp;hmpl=&amp;hmcu=&amp;hmkw=&amp;hmci=" target="_blank" rel="nofollow"><img src="//cdn.v2ex.com/assets/sidebar/yundun_20190620_2.gif" border="0" width="100%" alt="云盾"></a>
<div class="sep10"></div>
<a href="https://datapacket.com/china-dedicated-server-chinese?utm_source=v2ex&amp;utm_medium=cpc&amp;utm_campaign=programmer&amp;utm_content=static1" target="_blank" rel="nofollow"><img src="//cdn.v2ex.com/assets/sidebar/datapacket_20190221.png" border="0" width="100%" alt="DataPacket"></a>
<div class="sep10"></div>
<a href="https://www.cdn77.com/cn?utm_source=v2ex&amp;utm_medium=cpc&amp;utm_campaign=programmer&amp;utm_content=static2" target="_blank" rel="nofollow"><img src="//cdn.v2ex.com/assets/sidebar/cdn77_20190117.png" border="0" width="100%" alt="CDN77"></a>
<div class="sep10"></div>
<a href="https://www.yunaq.com/activity/818_web/?from=v2ex90806" target="_blank" rel="nofollow"><img src="//cdn.v2ex.com/assets/sidebar/yunaq_20190806.gif" border="0" width="100%" alt="知道创宇"></a>
<div class="sep10"></div>
<a href="https://www.qiniu.com/products/kodo/goglobal?promotion=v2ex" target="_blank" rel="nofollow"><img src="//cdn.v2ex.com/assets/sidebar/qiniu_20190802.png" border="0" width="100%" alt="七牛"></a>
<div class="sep10"></div>
<a href="http://www.oray.com/jump/45803" target="_blank" rel="nofollow"><img src="//cdn.v2ex.com/assets/sidebar/sunlogin_20190801.gif" border="0" width="100%" alt="贝锐 - 向日葵"></a>
<div class="sep10"></div>
<a href="https://www.oray.com/jump/45801" target="_blank" rel="nofollow"><img src="//cdn.v2ex.com/assets/sidebar/pgy_20190729.png" border="0" width="100%" alt="蒲公英"></a>
<div class="sep10"></div>

</div>

<style type="text/css">
#Wrapper {
background-color: #e2e2e2;
background-image: url("/static/img/shadow_light.png"), url("//cdn.v2ex.com/assets/bgs/circuit.png");
background-repeat: repeat-x, repeat-x;
}
</style>
</div>
<div class="sidebar_compliance"><a href="/advertise" target="_blank">广告</a></div>
</div>
<div class="sep20"></div>
<div class="box">
<div class="inner">
<script async="" src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>

<ins class="adsbygoogle" style="display:inline-block;width:250px;height:250px" data-ad-client="ca-pub-3465543440750523" data-ad-slot="9619519096" data-adsbygoogle-status="done"><ins id="aswift_0_expand" style="display:inline-table;border:none;height:250px;margin:0;padding:0;position:relative;visibility:visible;width:250px;background-color:transparent;"><ins id="aswift_0_anchor" style="display:block;border:none;height:250px;margin:0;padding:0;position:relative;visibility:visible;width:250px;background-color:transparent;"><iframe width="250" height="250" frameborder="0" marginwidth="0" marginheight="0" vspace="0" hspace="0" allowtransparency="true" scrolling="no" allowfullscreen="true" onload="var i=this.id,s=window.google_iframe_oncopy,H=s&amp;&amp;s.handlers,h=H&amp;&amp;H[i],w=this.contentWindow,d;try{d=w.document}catch(e){}if(h&amp;&amp;d&amp;&amp;(!d.body||!d.body.firstChild)){if(h.call){setTimeout(h,0)}else if(h.match){try{h=s.upd(h,i)}catch(e){}w.location.replace(h)}}" id="aswift_0" name="aswift_0" style="left:0;position:absolute;top:0;border:0px;width:250px;height:250px;"></iframe></ins></ins></ins>
<script>
        (adsbygoogle = window.adsbygoogle || []).push({});
        </script>
</div>
</div>
<div class="sep20"></div>
</div>
<div id="Main">
<div class="sep20"></div>
<div class="box" style="border-bottom: 0px;">
<div class="header"><div class="fr"><a href="/member/y0bcn"><img loading="lazy" src="//cdn.v2ex.com/avatar/7d69/de62/318621_large.png?m=1557850658" class="avatar" border="0" align="default"></a></div>
<a href="/">V2EX</a> <span class="chevron">&nbsp;›&nbsp;</span> <a href="/go/programmer">程序员</a>
<div class="sep10"></div>
<h1>现在 Go 环境怎么样</h1>
<div id="topic_596921_votes" class="votes">
<a href="javascript:" onclick="upVoteTopic(596921);" class="vote"><li class="fa fa-chevron-up"></li></a> &nbsp;<a href="javascript:" onclick="downVoteTopic(596921);" class="vote"><li class="fa fa-chevron-down"></li></a></div> &nbsp; <small class="gray"><a href="/member/y0bcn">y0bcn</a> · 23 小时 14 分钟前 · 5462 次点击</small>
</div>
<div class="cell">
<div class="topic_content"><div class="markdown_body"><p>目前大二在读，有没有必要从 Java 转到 go，坐标二线城市，有意去一线城市发展，但无意扎根，问一下 V 友，如果现在转去学 Go 有没有好的前景呢？</p>
</div></div>
</div>
</div>
<div class="sep20"></div>
<div class="box">
<div class="cell"><div class="fr" style="margin: -3px -5px 0px 0px;"><a href="/tag/Java" class="tag"><li class="fa fa-tag"></li> Java</a><a href="/tag/坐标" class="tag"><li class="fa fa-tag"></li> 坐标</a><a href="/tag/扎根" class="tag"><li class="fa fa-tag"></li> 扎根</a><a href="/tag/城市" class="tag"><li class="fa fa-tag"></li> 城市</a></div><span class="gray">58 回复 &nbsp;<strong class="snow">|</strong> &nbsp;直到 2019-09-02 12:37:52 +08:00</span>
</div>
<div id="r_7839573" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/f5d9/6046/335574_normal.png?m=1567083069" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">1</span></div>
<div class="sep3"></div>
<strong><a href="/member/tomato1111" class="dark">tomato1111</a></strong>&nbsp; &nbsp;<span class="ago">23 小时 2 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">楼主用 java 写出过产品了么</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839584" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/067f/ec6c/369285_normal.png?m=1544414848" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">2</span></div>
<div class="sep3"></div>
<strong><a href="/member/dabaibai" class="dark">dabaibai</a></strong>&nbsp; &nbsp;<span class="ago">22 小时 59 分钟前</span> &nbsp; <span class="small fade">♥ 3</span>
<div class="sep5"></div>
<div class="reply_content">大肠都在用。小肠都死了。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839600" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/f2e7a70f6921bf3648b4906cb2900462?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">3</span></div>
<div class="sep3"></div>
<strong><a href="/member/sunriz" class="dark">sunriz</a></strong>&nbsp; &nbsp;<span class="ago">22 小时 53 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">学语言很快的，我刚入职的时候 go 就学了三天，用起来基本上也没什么问题，当然更深入的东西也要平时积累。先学好计算机的基础课程吧，这个才是重点，公司招人并不是很看重你用过什么语言，看重基础知识和解决问题的思路。实践方面，多做些有意义的东西，目的在于用语言的特性做出高效的东西而不是语言本身。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839608" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/a08e/32d2/4901_normal.png?m=1335109848" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">4</span></div>
<div class="sep3"></div>
<strong><a href="/member/reus" class="dark">reus</a></strong>&nbsp; &nbsp;<span class="ago">22 小时 51 分钟前</span> &nbsp; <span class="small fade">♥ 4</span>
<div class="sep5"></div>
<div class="reply_content">rust 才是世界上最好的编程语言</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839618" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/19128ba5a8eb55028bdddc0d51c27b58?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">5</span></div>
<div class="sep3"></div>
<strong><a href="/member/sls" class="dark">sls</a></strong>&nbsp; &nbsp;<span class="ago">22 小时 47 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">慎重，自从用了这个简单的语言后，其他复杂的语言都学不动了，但是选语言还是看使用场景和生态。所以先学其他的吧，从繁入简容易，由简入繁难啊</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839638" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/76c8a21fc206c8d98f13d0ce4c6f943a?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">6</span></div>
<div class="sep3"></div>
<strong><a href="/member/iPhoneXI" class="dark">iPhoneXI</a></strong>&nbsp; &nbsp;<span class="ago">22 小时 35 分钟前 via Android</span>
<div class="sep5"></div>
<div class="reply_content">大学时间那么多，当然是都学</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839645" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/4459/ddc6/430616_normal.png?m=1563592112" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">7</span></div>
<div class="sep3"></div>
<strong><a href="/member/linsijia1002" class="dark">linsijia1002</a></strong>&nbsp; &nbsp;<span class="ago">22 小时 31 分钟前 via iPhone</span>
<div class="sep5"></div>
<div class="reply_content">Go 是比较年轻的语言，直接机器码运行，效率比 java 快很多倍。编译配置方便简单，导入包很方便，相比于 java 的配置麻烦 xml 语法，极简主义者表示有点密集恐惧。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839706" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/f1e954294feec2f01c3f5cc4e1d2c14c?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">8</span></div>
<div class="sep3"></div>
<strong><a href="/member/troywinter" class="dark">troywinter</a></strong>&nbsp; &nbsp;<span class="ago">22 小时 3 分钟前</span> &nbsp; <span class="small fade">♥ 20</span>
<div class="sep5"></div>
<div class="reply_content">无意引战，目前在 web 开发领域 go 还是没法和 java 相比，如果你是那种完全不写测试的人，认为自己写的代码都可以手工测试或者不需要测试，那你用什么语言其实都没有区别，java 语言提供的动态性以及 jvm 的 jit 和一系列的运行时优化，都使得一是做 di 非常简单快速并有庞大的库帮助你在十几分钟内构建完善的单元测试流程，二是 jit 的优化，在一些 critical path 上的性能可以达到非常高的水平。在开发关键大型系统时，良好的测试覆盖可以有效减少 bug。<br><br>一些人总是提出直接 go 是直接机器码运行，比 java 快，其实不全对，没有 jit 加持的 java 确实性能很差，go 的方式是通过确定性的机器码的运行方式，使得性能是可预见的，不会出现 java 那种 jit 优化之后性能极速上升的例子，但性能可以维持在比较平稳的状态。<br><br>至于配置麻烦的问题，你如果不用 di，那也不需要什么配置，反正我司都是一套 bazel 构建到底，构建也都不用操心。<br><br>换语言时，要考虑到方方面面的问题，以及很多问题可选的解决思路，不应该迷失在别人的观点中。<br><br>如果你决定转 go 了，那么祝你找到心怡的工作，可以多交流，我也在写 go。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839852" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/6a75/acd4/235987_normal.png?m=1515376020" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">9</span></div>
<div class="sep3"></div>
<strong><a href="/member/jss" class="dark">jss</a></strong>&nbsp; &nbsp;<span class="ago">20 小时 54 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">有人说：我们现在不是考虑要不要学 Go，而是考虑怎样挤出更多的时间了解学习 Go。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839861" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/205a/8caf/35551_normal.png?m=1546830088" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">10</span></div>
<div class="sep3"></div>
<strong><a href="/member/jamesliu96" class="dark">jamesliu96</a></strong>&nbsp; &nbsp;<span class="ago">20 小时 49 分钟前 via Android</span>
<div class="sep5"></div>
<div class="reply_content">去看看著名的屑 openbilibili/go-common，只要写得比它好就行（</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839877" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/a9a5/69e4/139381_normal.png?m=1460618190" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">11</span></div>
<div class="sep3"></div>
<strong><a href="/member/qinxi" class="dark">qinxi</a></strong>&nbsp; &nbsp;<span class="ago">20 小时 44 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/linsijia1002">linsijia1002</a> #7 9102 年了</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839895" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/0178/b453/362314_normal.png?m=1566552099" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">12</span></div>
<div class="sep3"></div>
<strong><a href="/member/MMMMMMMMMMMMMMMM" class="dark">MMMMMMMMMMMMMMMM</a></strong>&nbsp; &nbsp;<span class="ago">20 小时 39 分钟前</span> &nbsp; <span class="small fade">♥ 1</span>
<div class="sep5"></div>
<div class="reply_content">我永远喜欢 node(划掉) Python(划掉) Go(划掉) Rust(黑体 加粗).jpg</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839919" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/4271/6be2/154020_normal.png?m=1544423582" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">13</span></div>
<div class="sep3"></div>
<strong><a href="/member/misaka19000" class="dark">misaka19000</a></strong>&nbsp; &nbsp;<span class="ago">20 小时 26 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">都学，这样工作你都能胜任</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839926" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/dfa2/2bd5/274059_normal.png?m=1517483107" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">14</span></div>
<div class="sep3"></div>
<strong><a href="/member/salamanderMH" class="dark">salamanderMH</a></strong>&nbsp; &nbsp;<span class="ago">20 小时 23 分钟前 via Android</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/troywinter">troywinter</a> 这个是正经回答</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7839964" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/bc9c/0f39/338151_normal.png?m=1536383557" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">15</span></div>
<div class="sep3"></div>
<strong><a href="/member/xmge" class="dark">xmge</a></strong>&nbsp; &nbsp;<span class="ago">20 小时 7 分钟前</span> &nbsp; <span class="small fade">♥ 1</span>
<div class="sep5"></div>
<div class="reply_content">哈哈，这个我说一下，从 java 转 go 两年，go 的应用场景呢，高并发，为什么高并发，创建的协程消耗比线程少。<br><br>一般都是大公司在用，因为大公司目前才有高并发的场景。<br><br>目前的岗位不是很多，面试的机会其实也不是很多，自我感觉，前几天去找来着。<br><br>但是，未来，如果物联网发展越来越迅猛的话，感觉 go 会起来</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840043" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/7d69/de62/318621_normal.png?m=1557850658" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">16</span></div>
<div class="sep3"></div>
<strong><a href="/member/y0bcn" class="dark">y0bcn</a></strong>&nbsp; &nbsp;<span class="ago">19 小时 29 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/tomato1111">tomato1111</a> 写过几个比赛的项目</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840058" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/d678/e742/346043_normal.png?m=1538628177" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">17</span></div>
<div class="sep3"></div>
<strong><a href="/member/tourist2018" class="dark">tourist2018</a></strong>&nbsp; &nbsp;<span class="ago">19 小时 23 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">头条 滴滴 微博 阿里 腾讯 小米 这些是我去面试过的公司 基本也都是核心业务部门在用了<br><br>我觉得如果你的目标是一线二线互联网公司的话 完全不用担心现在 golang 没有工作机会 但是我感觉和语言没啥关系 只要别学特别偏的。。。比如 C# 这个是真的不好找互联网公司的工作</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840062" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/d678/e742/346043_normal.png?m=1538628177" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">18</span></div>
<div class="sep3"></div>
<strong><a href="/member/tourist2018" class="dark">tourist2018</a></strong>&nbsp; &nbsp;<span class="ago">19 小时 22 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">补充下 面过的 小米视频业务 微博是广告业务 阿里是阿里云 腾讯腾讯视频企鹅号</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840076" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/1ed3/f6c5/22404_normal.png?m=1340101995" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">19</span></div>
<div class="sep3"></div>
<strong><a href="/member/swulling" class="dark">swulling</a></strong>&nbsp; &nbsp;<span class="ago">19 小时 16 分钟前 via iPhone</span>
<div class="sep5"></div>
<div class="reply_content">go 的第三方库质量和数量都不太好。如果你是业务为王，还是用 java 吧</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840091" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/e0f8/7536/20445_normal.png?m=1497274833" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">20</span></div>
<div class="sep3"></div>
<strong><a href="/member/hst001" class="dark">hst001</a></strong>&nbsp; &nbsp;<span class="ago">19 小时 13 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/MMMMMMMMMMMMMMMM">MMMMMMMMMMMMMMMM</a> #12 戏很足</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840153" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/c3c23853be0a1dd4b30d41516e90e70d?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">21</span></div>
<div class="sep3"></div>
<strong><a href="/member/pursuer" class="dark">pursuer</a></strong>&nbsp; &nbsp;<span class="ago">18 小时 52 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">现在各个领域有最适合的语言，可以看看自己感兴趣的领域</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840159" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/5892/b3e1/401764_normal.png?m=1562648062" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">22</span></div>
<div class="sep3"></div>
<strong><a href="/member/ScepterZ" class="dark">ScepterZ</a></strong>&nbsp; &nbsp;<span class="ago">18 小时 50 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">都学吧</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840295" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/185c/08dd/216958_normal.png?m=1549026431" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">23</span></div>
<div class="sep3"></div>
<strong><a href="/member/mamahaha" class="dark">mamahaha</a></strong>&nbsp; &nbsp;<span class="ago">18 小时 1 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">大公司用 go 可以节约服务器成本，如果进不了大公司就别瞎起哄了。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840309" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/2b1be4616bf957f8963ae0ba016d5799?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">24</span></div>
<div class="sep3"></div>
<strong><a href="/member/AX5N" class="dark">AX5N</a></strong>&nbsp; &nbsp;<span class="ago">17 小时 54 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/linsijia1002">linsijia1002</a> 胡扯，java 也就比 c/c++慢一点，你 go 能快很多倍那岂不是比 c 还要快？</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840315" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/4e77/419f/60957_normal.png?m=1567337842" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">25</span></div>
<div class="sep3"></div>
<strong><a href="/member/ZoolYe" class="dark">ZoolYe</a></strong>&nbsp; &nbsp;<span class="ago">17 小时 53 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/linsijia1002">linsijia1002</a> 都什么年代了，Java 还用 XML 配置吗？你当 SpringBoot、SpringCloud 不存在吗？ JVM 这么多年的优化，你敢说 Go 还比 Java 快很多倍？</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840639" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/c004f07e4aeea5ada1cc18fa390f6e09?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">26</span></div>
<div class="sep3"></div>
<strong><a href="/member/bakabie" class="dark">bakabie</a></strong>&nbsp; &nbsp;<span class="ago">15 小时 43 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">不是非必要，但是有时间可以学</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840671" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/8d9d/5e78/149675_normal.png?m=1458072565" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">27</span></div>
<div class="sep3"></div>
<strong><a href="/member/Srar" class="dark">Srar</a></strong>&nbsp; &nbsp;<span class="ago">15 小时 32 分钟前</span> &nbsp; <span class="small fade">♥ 8</span>
<div class="sep5"></div>
<div class="reply_content">慢的主要原因<br><a target="_blank" href="https://i.imgur.com/FYuw4yL.jpg" rel="nofollow"><img src="https://i.imgur.com/FYuw4yL.jpg" class="embedded_image"></a></div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840693" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/612b/2735/130165_normal.png?m=1461222228" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">28</span></div>
<div class="sep3"></div>
<strong><a href="/member/akring" class="dark">akring</a></strong>&nbsp; &nbsp;<span class="ago">15 小时 22 分钟前</span> &nbsp; <span class="small fade">♥ 2</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/troywinter">troywinter</a> 像您这样有价值，态度又平和诚恳的回帖真的太少了，必须点个赞</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840720" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/eed5/484b/152333_normal.png?m=1563376004" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">29</span></div>
<div class="sep3"></div>
<strong><a href="/member/janus77" class="dark">janus77</a></strong>&nbsp; &nbsp;<span class="ago">15 小时 10 分钟前 via iPhone</span>
<div class="sep5"></div>
<div class="reply_content">大二在读谈什么转？我寻思你靠他混饭吃呢。当然是都学啊</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7840881" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/f55bd771888fddbf65864a8dab858385?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">30</span></div>
<div class="sep3"></div>
<strong><a href="/member/cheesea" class="dark">cheesea</a></strong>&nbsp; &nbsp;<span class="ago">13 小时 55 分钟前</span> &nbsp; <span class="small fade">♥ 3</span>
<div class="sep5"></div>
<div class="reply_content">我觉得最忌讳的就是给自己划定圈子，然后就只关注这个圈子的东西，这样不好。<br>比如我是 java/go/python 工程师，我是后端 /前端 /运维，甚至我是程序员，然后其他领域的就什么都不懂，也不去了解，这会限制你的发展和眼界。<br>你才大二，应该多去接触和了解各个领域的东西。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841025" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/5719/62a6/20396_normal.png?m=1450629628" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">31</span></div>
<div class="sep3"></div>
<strong><a href="/member/ihciah" class="dark">ihciah</a></strong>&nbsp; &nbsp;<span class="ago">12 小时 1 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">作为一个 Golang 用户，暑假去了阿里，然后学了两个月 Java。。。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841050" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/b86d/b831/241141_normal.png?m=1514864939" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">32</span></div>
<div class="sep3"></div>
<strong><a href="/member/blless" class="dark">blless</a></strong>&nbsp; &nbsp;<span class="ago">11 小时 31 分钟前 via Android</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/troywinter">troywinter</a> 测试我还是觉得 go 写得爽，测试文件就在对应代码一起，我觉得很舒服，因为测试用例不仅仅只是测试，其实也算使用方式教程之类的。接口测试天然 mock，自带 benchmark。<br>jit 虽然快，但是也要预热外加消耗大量内存，用起来还是不如 go 爽</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841072" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/f91c/6b29/55150_normal.png?m=1559317585" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">33</span></div>
<div class="sep3"></div>
<strong><a href="/member/ericgui" class="dark">ericgui</a></strong>&nbsp; &nbsp;<span class="ago">10 小时 29 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">你才大二，着急干什么<br><br>先把 Java 学好，如果你 Java 学好了，找工作不愁，而且你再学其他任何语言，都小菜一碟</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841206" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/f86f713f0b901fe1b11fdcb5363ed2d9?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">34</span></div>
<div class="sep3"></div>
<strong><a href="/member/qbmiller" class="dark">qbmiller</a></strong>&nbsp; &nbsp;<span class="ago">5 小时 36 分钟前 via iPhone</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/jamesliu96">jamesliu96</a> #10 哈哈</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841223" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/6a75/acd4/235987_normal.png?m=1515376020" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">35</span></div>
<div class="sep3"></div>
<strong><a href="/member/jss" class="dark">jss</a></strong>&nbsp; &nbsp;<span class="ago">5 小时 26 分钟前 via iPhone</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/AX5N">AX5N</a> 保守估计在相同的环境和执行目标的情况下，Go 程序比 Java 或 Scala 应用程序要快上 2 倍，并比这两门语言使用少占用 70% 的内存。如果说 Go 比 C++ 要慢 20%，那么 Go 就要比任何非静态和编译型语言快 2 到 10 倍，并且能够更加高效地使用内存。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841257" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/6dfe440ab1f0524d4c9b190666a314cd?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">36</span></div>
<div class="sep3"></div>
<strong><a href="/member/jonsun30" class="dark">jonsun30</a></strong>&nbsp; &nbsp;<span class="ago">5 小时 12 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/jss">jss</a> 你的 Java 程序 JIT 优化过了吗？还有更高效的使用内存这句话真的太主观武断了吧？<br>'Will your toy benchmark program be faster if you write it in a different programming language? It depends on how you write it!'</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841263" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/e18b/34dd/326514_normal.png?m=1559698761" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">37</span></div>
<div class="sep3"></div>
<strong><a href="/member/tt67wq" class="dark">tt67wq</a></strong>&nbsp; &nbsp;<span class="ago">5 小时 9 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">搞 k8s 很赚，其他的一般般</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841322" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/f72f7454ce9d710baa506394f68f4132?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">38</span></div>
<div class="sep3"></div>
<strong><a href="/member/fuxiaohei" class="dark">fuxiaohei</a></strong>&nbsp; &nbsp;<span class="ago">4 小时 56 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">如果是二线城市，还是用的人多的语言机会更多</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841343" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/a8107cfefeeb689b9039dc6658d7427f?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">39</span></div>
<div class="sep3"></div>
<strong><a href="/member/tairan2006" class="dark">tairan2006</a></strong>&nbsp; &nbsp;<span class="ago">4 小时 53 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">都学啊<br><br>不过写的话还是 go 比较爽，除了没泛型</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841406" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/5c4a/f2ac/35375_normal.png?m=1438310082" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">40</span></div>
<div class="sep3"></div>
<strong><a href="/member/jianson2006" class="dark">jianson2006</a></strong>&nbsp; &nbsp;<span class="ago">4 小时 42 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/sls">sls</a> 十分赞同你的观点，我是零基自学 PY，但我觉得总缺对代码的理解。现在学习 C 语言。我认真对待学习还是从 C 开始，效果确实有很大的不同。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841476" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/ce914821c74290edf21e43381979ef94?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">41</span></div>
<div class="sep3"></div>
<strong><a href="/member/encro" class="dark">encro</a></strong>&nbsp; &nbsp;<span class="ago">4 小时 28 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">如果自认为 NB，学习 GO，写一个 NB 的开源项目，比如 NPS，frp，bleve, beego, etcd 这样的，基本找工作不会有问题了。<br><br>如果自认为不太 NB，学习 JAVA，慢慢来，毕竟 JAVA 生态环境摆在那里，还有很多螺丝需要拧。<br><br>打算回二线，具体是哪个城市，去搜索下看看，目前用 GO 的机会很少。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841510" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/f8249bbbafd7a0106dd0c2c718aa8e5c?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">42</span></div>
<div class="sep3"></div>
<strong><a href="/member/kilen3a" class="dark">kilen3a</a></strong>&nbsp; &nbsp;<span class="ago">4 小时 25 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">都学啊，反正才大二，时间多的是</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841532" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/b319f212688bdde90402d3492f3bb14c?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">43</span></div>
<div class="sep3"></div>
<strong><a href="/member/phantomzz" class="dark">phantomzz</a></strong>&nbsp; &nbsp;<span class="ago">4 小时 20 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">目前 90%以上的程序员写的代码还触及不到语言的性能瓶颈。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841669" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/9bf6d12008756afd2d5b9e03cc67be42?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">44</span></div>
<div class="sep3"></div>
<strong><a href="/member/notreami" class="dark">notreami</a></strong>&nbsp; &nbsp;<span class="ago">4 小时 0 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">为啥总有人盯着语言呢。java 用来做业务、大数据处理。golang 用来做运维工具。<br>那么你想搞业务、大数据，还是服务运维？？</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841821" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/fe4b/4d25/250548_normal.png?m=1541913171" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">45</span></div>
<div class="sep3"></div>
<strong><a href="/member/no1xsyzy" class="dark">no1xsyzy</a></strong>&nbsp; &nbsp;<span class="ago">3 小时 43 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">除了提供新的思维方法外，新语言无意义<br>一般人们所称转换语言是转换生态圈，还有就是少打几个字符（懒）<br>不同语言间生态割裂其实是个问题<br>唯一的纽带是 C</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841837" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/fe4b/4d25/250548_normal.png?m=1541913171" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">46</span></div>
<div class="sep3"></div>
<strong><a href="/member/no1xsyzy" class="dark">no1xsyzy</a></strong>&nbsp; &nbsp;<span class="ago">3 小时 41 分钟前</span> &nbsp; <span class="small fade">♥ 1</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/jianson2006">jianson2006</a> （但是 Python 对于代码抽象的处理程度不是比 C 高吗？</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841940" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/6429/6b5f/148127_normal.png?m=1512136951" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">47</span></div>
<div class="sep3"></div>
<strong><a href="/member/daryl" class="dark">daryl</a></strong>&nbsp; &nbsp;<span class="ago">3 小时 30 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">巧了，我想从 go 转 java。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841941" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/552c/b046/401115_normal.png?m=1563586025" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">48</span></div>
<div class="sep3"></div>
<strong><a href="/member/luckRay" class="dark">luckRay</a></strong>&nbsp; &nbsp;<span class="ago">3 小时 30 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">推荐还是 JAVA 为主吧，毕竟体量在这里。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7841993" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/6a75/acd4/235987_normal.png?m=1515376020" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">49</span></div>
<div class="sep3"></div>
<strong><a href="/member/jss" class="dark">jss</a></strong>&nbsp; &nbsp;<span class="ago">3 小时 24 分钟前 via iPhone</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/jonsun30">jonsun30</a> [转述大佬的话] Java 和 Scala 使用的 JVM，C# 和 <a target="_blank" href="http://VB.NET" rel="nofollow">VB.NET</a> 使用的 .NET CLR。尽管虚拟机的性能已经有了很大的提升，但任何使用 JIT 编译器和脚本语言解释器的编程语言（ Ruby、Python、Perl 和 JavaScript ）在 C 和 C++ 的绝对优势下甚至都无法在性能上望其项背。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842240" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/5c4a/f2ac/35375_normal.png?m=1438310082" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">50</span></div>
<div class="sep3"></div>
<strong><a href="/member/jianson2006" class="dark">jianson2006</a></strong>&nbsp; &nbsp;<span class="ago">2 小时 55 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/no1xsyzy">no1xsyzy</a> PY 学习下来我只其用法，但不知道为什么要这样用？有很多原理我是不懂的，所以，我想从 C 中获取。就是编码到底是如何与硬件沟通的？我只是想把在提高的过程中可以把一件原理也可以搞明白。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842434" class="cell">
 <table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/30d173558f95295004320b7a3fa4939f?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">51</span></div>
<div class="sep3"></div>
<strong><a href="/member/hyl24" class="dark">hyl24</a></strong>&nbsp; &nbsp;<span class="ago">2 小时 32 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">学 java 吧 以后 java 往中间件方向走 就学 GO</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842472" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/737aacd5309c078d3d10c9e7989a8396?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">52</span></div>
<div class="sep3"></div>
<strong><a href="/member/DovaKeen" class="dark">DovaKeen</a></strong>&nbsp; &nbsp;<span class="ago">2 小时 28 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">你才大二，可以多了解一些的</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842474" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/53d5/7871/30158_normal.png?m=1359105228" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">53</span></div>
<div class="sep3"></div>
<strong><a href="/member/Galileo" class="dark">Galileo</a></strong>&nbsp; &nbsp;<span class="ago">2 小时 28 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">请谨慎考虑</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842597" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/6dfe440ab1f0524d4c9b190666a314cd?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">54</span></div>
<div class="sep3"></div>
<strong><a href="/member/jonsun30" class="dark">jonsun30</a></strong>&nbsp; &nbsp;<span class="ago">2 小时 14 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/jss">jss</a> <a target="_blank" href="https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go.html" rel="nofollow">https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go.html</a><br>Btw it's just a benchmarking game XD</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842628" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/50d0943fb104816ef48cbc7842b975dc?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">55</span></div>
<div class="sep3"></div>
<strong><a href="/member/Raymon111111" class="dark">Raymon111111</a></strong>&nbsp; &nbsp;<span class="ago">2 小时 9 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">语言完全不是重点<br><br>把基础知识学好吧, 比如数据结构, 操作系统</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842647" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/dbaf/7932/254077_normal.png?m=1540180352" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">56</span></div>
<div class="sep3"></div>
<strong><a href="/member/gz911122" class="dark">gz911122</a></strong>&nbsp; &nbsp;<span class="ago">2 小时 6 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/jonsun30">jonsun30</a> 这个测试报告来看似乎 java 更快一些?</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842703" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/gravatar/6dfe440ab1f0524d4c9b190666a314cd?s=48&amp;d=retro" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">57</span></div>
<div class="sep3"></div>
<strong><a href="/member/jonsun30" class="dark">jonsun30</a></strong>&nbsp; &nbsp;<span class="ago">2 小时 0 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">@<a href="/member/gz911122">gz911122</a> 其实我想说的是这种 assumption 是没有意义的，每个语言都有其诞生的理由。绝大部分程序员（包括我自己）所写的程序还真的触及不到语言的瓶颈，有时候是你摸不到性能的天花板，还有时候是我们的代码由于各种各样的原因真的很屎。</div>
</td>
</tr>
</tbody></table>
</div>
<div id="r_7842977" class="cell">
<table cellpadding="0" cellspacing="0" border="0" width="100%">
<tbody><tr>
<td width="48" valign="top" align="center"><img loading="lazy" src="//cdn.v2ex.com/avatar/2f96/043a/171860_normal.png?m=1471502225" class="avatar" border="0" align="default"></td>
<td width="10" valign="top"></td>
<td width="auto" valign="top" align="left"><div class="fr"> &nbsp; &nbsp; <span class="no">58</span></div>
<div class="sep3"></div>
<strong><a href="/member/richzhu" class="dark">richzhu</a></strong>&nbsp; &nbsp;<span class="ago">1 小时 10 分钟前</span>
<div class="sep5"></div>
<div class="reply_content">一提语言就吵架，我真的服了，性能高能怎么地？ 能高多少？ 性能高的代码你能写出来？写出来了别人能看懂？你的业务多少 qps 就性能性能的....好好学学计算机基础，哪个场景就用哪个语言，个人喜欢 go 就学 go，总之就是学起来别犹豫，学会了一个再学另一个也很简单，不能因为犹豫学哪个而耽误时间。</div>
</td>
</tr>
</tbody></table>
</div>
</div>
<div class="sep20"></div>
</div>
</div>
<div class="c"></div>
<div class="sep20"></div>
</div>
<div id="Bottom">
<div class="content">
<div class="inner">
<div class="sep10"></div>
<div class="fr">
<a href="https://www.digitalocean.com/?refcode=1b51f1a7651d" target="_blank"><div id="DigitalOcean"></div></a>
</div>
<strong><a href="/about" class="dark" target="_self">关于</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/faq" class="dark" target="_self">FAQ</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/p/7v9TEc53" class="dark" target="_self">API</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/mission" class="dark" target="_self">我们的愿景</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/advertise" class="dark" target="_self">广告投放</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/advertise/2017.html" class="dark" target="_self">感谢</a> &nbsp; <span class="snow">·</span> &nbsp; <a href="/tools" class="dark" target="_self">实用小工具</a> &nbsp; <span class="snow">·</span> &nbsp; 3836 人在线</strong> &nbsp; <span class="fade">最高记录 5043</span> &nbsp; <span class="snow">·</span> &nbsp; <a href="/select/language" class="f11"><img src="/static/img/language.png?v=6a5cfa731dc71a3769f6daace6784739" width="16" align="absmiddle" id="ico-select-language"> &nbsp; Select Language</a>
<div class="sep20"></div>
创意工作者们的社区
<div class="sep5"></div>
World is powered by solitude
<div class="sep20"></div>
<span class="small fade">VERSION: 3.9.8.3 · 25ms · UTC 05:48 · PVG 13:48 · LAX 22:48 · JFK 01:48<br>♥ Do have faith in what you're doing.</span>
<div class="sep10"></div>
</div>
</div>
</div>
<script src="/b/i/FZi8FiVN9mab9PM72oEZpg8T2H196SEUGFED0paIctdTTkajsLWfLqoH-cYxu_gqVchlJFnNmH3sEQhpx0sb5DaiwSAe549oJYitZUfrWJaV1yk7BqnZ5p-XFjAGfDaoEt9ynTWJHW6Y82GCgrKsNcE3BZHnfqczpfxON57yJ0g="></script>
<script>
	  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
	  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
	  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
	  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

	  ga('create', 'UA-11940834-2', 'v2ex.com');
	  ga('send', 'pageview');
      

ga('send', 'event', 'Node', 'topic', 'programmer');



	</script>

</body>`
