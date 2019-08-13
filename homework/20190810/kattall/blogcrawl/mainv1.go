package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/huang19910425/blogcrawl/crawl"
	"os"
	"time"
)

const (
	url = "https://imsilence.github.io"
)

func main(){
	startTime := time.Now()
	var c crawl.Reptile = crawl.Article{}
	var articles = []crawl.Article{}

	document, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	startPage, endPage := c.GetPages(document)
	onePageArticles := c.GetArticle(url, document)
	articles = append(articles, onePageArticles...)
	for i:=startPage; i<=endPage; i++ {
		document, err = goquery.NewDocument(fmt.Sprintf("%s/page/%d/", url, i))
		if err != nil {
			fmt.Println(err)
			continue
		}

		pageArticles := c.GetArticle(url, document)
		articles = append(articles, pageArticles...)
	}
	for i, article := range articles {
		fmt.Printf("序号：%d, %s\n", i+1, article)
	}

	fmt.Println("time:", time.Now().Sub(startTime))
}
