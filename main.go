package main

import (
	"fmt"
	"strconv"
	"strings"
	"v2ex-spider/article"

	prompt "github.com/c-bata/go-prompt"
)

var (
	BaseUrl    = "https://v2ex.com/?tab="
	ArticleUrl = "https://v2ex.com"
	M          = map[string]string{}
)

func v2ex(tab, index string) {
	var title, url string
	url = BaseUrl + tab
	currentArticles := article.GetArticles(url)
	for k, v := range currentArticles {
		if len(index) == 0 {
			fmt.Printf("序号: %d   评论: %s \tT: %s\n", k, v.Comment, v.Title)
		} else {
			if index == strconv.Itoa(k) {
				title = v.Title
			}
		}

		M[strconv.Itoa(k)] = v.Href
	}

	if len(index) > 0 {
		article := article.GetArticle(ArticleUrl + M[index])
		fmt.Printf("标题:%s\n描述:%s\n", title, article.Content)
		for _, v := range article.CommentContent {
			fmt.Println(v.Floor, "楼:", v.Content)
		}
	}
}

func inSlice(v string, sl []string) bool {
	if len(sl) == 0 {
		return false
	}
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func executor(in string) {
	var num string
	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	}
	arr := strings.Split(in, " ")
	if len(arr) == 2 {
		in = strings.TrimSpace(arr[0])
		num = strings.TrimSpace(arr[1])
	}

	if !inSlice(in, []string{"index", "tech", "jobs", "creative", "play", "apple", "deal", "city", "qna", "hot", "all", "r2"}) {
		fmt.Println("Your input not allowed")
		return
	}
	v2ex(in, num)
	LivePrefixState.LivePrefix = in + "> "
	LivePrefixState.IsEnable = true
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "index", Description: "首页"},
		{Text: "tech", Description: "技术"},
		{Text: "creative", Description: "创意"},
		{Text: "play", Description: "好玩"},
		{Text: "apple", Description: "苹果"},
		{Text: "jobs", Description: "酷工作"},
		{Text: "deal", Description: "交易"},
		{Text: "city", Description: "城市"},
		{Text: "qna", Description: "问与答"},
		{Text: "hot", Description: "最热"},
		{Text: "all", Description: "全部"},
		{Text: "r2", Description: "R2"},
		{Text: "example", Description: "举例: index 2 (查询首页列表下编号为2的文章)"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func main() {
	fmt.Println("请选择想要获取的Tab")
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("live-v2ex-spider"),
	)
	p.Run()
}
