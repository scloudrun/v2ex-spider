package article

import (
	"fmt"
	"strings"
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
	doc := getDocument(url)
	currentArticles := parse.ParseArticles(doc)
	return currentArticles

}

func GetArticle(url string) parse.V2exArticle {
	doc := getDocument(url)
	currentArticle := parse.ParseArticle(doc)
	return currentArticle
}
