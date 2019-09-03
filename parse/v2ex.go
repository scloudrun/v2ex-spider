package parse

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type V2exArticle struct {
	Title          string
	Href           string
	Desc           string
	Year           string
	Area           string
	Tag            string
	Content        string
	Comment        string
	CommentContent []Comment
}

type Comment struct {
	Content string
	Floor   string
}

type Page struct {
	Page int
	Url  string
}

func GetPages(url string) []Page {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	return ParsePages(doc)
}

func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, Url: ""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})

	return pages
}

// 解析文章列表
func ParseArticles(doc *goquery.Document) (articles []V2exArticle) {
	doc.Find("#Main .box .cell").Each(func(i int, s *goquery.Selection) {
		var reply string
		title := s.Find(".item_title a").Eq(0).Text()

		href, _ := s.Find(".item_title a").Eq(0).Attr("href")

		if reply = s.Find(".count_livid").Eq(0).Text(); len(reply) == 0 {
			reply = "0"
		}

		if len(title) > 0 {
			article := V2exArticle{
				Title:   title,
				Href:    href,
				Comment: reply,
			}
			articles = append(articles, article)
		}

	})

	return articles
}

// 解析文章
func ParseArticle(doc *goquery.Document) (article V2exArticle) {
	doc.Find("#Main .box .cell .topic_content").Each(func(i int, s *goquery.Selection) {

		content := s.Find(".markdown_body p").Eq(0).Text()
		if len(content) == 0 {
			content = s.Text()
		}

		if len(content) > 0 {
			article = V2exArticle{
				Content: content,
			}
		}

	})

	var commentArr []Comment
	doc.Find("#Main .box .cell").Each(func(i int, s *goquery.Selection) {

		content := s.Find(".reply_content").Eq(0).Text()
		floor := s.Find(".fr .no").Eq(0).Text()

		if len(content) > 0 {
			comment := Comment{
				Content: content,
				Floor:   floor,
			}
			commentArr = append(commentArr, comment)
		}
	})

	if len(commentArr) > 0 {
		article.CommentContent = commentArr
	}

	return article
}
