package article

import (
	"fmt"
	"strings"
	"v2ex-spider/cache"
	"v2ex-spider/parse"

	"github.com/PuerkitoBio/goquery"
)

func getDocument(url string) *goquery.Document {
	doc, err := goquery.NewDocument(strings.Join([]string{url}, ""))

	if err != nil {
		fmt.Println("goquery NewDocument err", err)
		return nil
	}
	return doc
}

func GetArticles(url string) []parse.V2exArticle {
	if res := cache.GetCache(url); res != nil {
		return res.([]parse.V2exArticle)
	}

	doc := getDocument(url)
	currentArticles := parse.ParseArticles(doc)

	cache.SetCache(url, currentArticles)

	return currentArticles

}

func GetArticle(url string) parse.V2exArticle {
	if res := cache.GetCache(url); res != nil {
		fmt.Println("获取cache article")
		return res.(parse.V2exArticle)
	}

	doc := getDocument(url)
	currentArticle := parse.ParseArticle(doc)

	cache.SetCache(url, currentArticle)
	fmt.Println("写入cache article")

	return currentArticle
}
