package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/huang19910425/blogcrawl/crawl"
	"os"
	"sync"
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

	var group sync.WaitGroup
	channel := make(chan crawl.Article, 1024)

	startPage, endPage := c.GetPages(document)
	onePageArticles := c.GetArticle(url, document)
	articles = append(articles, onePageArticles...)

	for i:=startPage; i<=endPage; i++ {
		group.Add(1)
		go func(i int, channel chan<- crawl.Article) {
			document, err = goquery.NewDocument(fmt.Sprintf("%s/page/%d/", url, i))
			if err != nil {
				fmt.Println(err)
			}
			pageArticles := c.GetArticle(url, document)
			for _, article := range pageArticles {
				channel<- article
			}
			group.Done()
		}(i, channel)
	}

	go func() {
		group.Wait()
		close(channel)
	}()

	id := 0
	for i, article := range articles {
		id = i + 1
		fmt.Printf("%d: %s\n", id, article)
	}

	for article := range channel {
		id ++
		fmt.Printf("%d: %s\n", id, article)
	}

	fmt.Println("time:", time.Now().Sub(startTime))
}
