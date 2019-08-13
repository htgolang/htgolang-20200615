package crawl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Reptile interface {
	GetArticle(url string, document *goquery.Document) []Article
	GetPages(documment *goquery.Document) (int, int)
}

type Article struct {
	Title   string `json: "title"`   // 文章标题
	Url     string `json: "url"`     // 文章url
	PubTime string `json: "pubtime"` // 文章发布时间
	Ttype   string `json: "ttype"`   // 文章类型
}

func (a Article) String() string {
	return fmt.Sprintf("文章：%s, 发布时间：%s, 类型：%s, 文章链接：%s", a.Title, a.PubTime, a.Ttype, a.Url)
}

func (a Article) GetPages(document *goquery.Document) (startPage int, endPage int) {
	startPage, endPage = 1, -1
	document.Find("nav.pagination .page-number").Each(func(index int, selection *goquery.Selection) {
		if i, err := strconv.Atoi(selection.Text()); err == nil {
			if i < startPage {
				startPage = i
			}
			if i > endPage {
				endPage = i
			}
		}
	})

	return startPage, endPage
}

func (a Article) GetArticle(url string, document *goquery.Document) []Article {
	articles := []Article{}
	document.Find("article.post").Each(func(index int, selection *goquery.Selection) {
		title := strings.TrimSpace(selection.Find("h1.post-title > a.post-title-link").Text())
		u, exists := selection.Find("h1.post-title > a.post-title-link").Attr("href")
		if exists {
			u = fmt.Sprintf("%s%s", url, strings.TrimSpace(u))
		}
		pubtime := strings.TrimSpace(selection.Find("div.post-meta > span.post-time time").Text())
		ttype := strings.TrimSpace(selection.Find("div.post-meta > span.post-category > span > a > span").Text())
		articles = append(articles, Article{Title: title, Url: u, PubTime: pubtime, Ttype: ttype})
	})

	return articles
}
