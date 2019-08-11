package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Name string
	Time string
}

func (a Article) String() string {
	return fmt.Sprintf("%s %s", a.Name, a.Time)
}

func getPagination(document *goquery.Document) (int, int) {
	start, end := 0XFFFF, -1
	document.Find("nav.pagination .page-number").Each(func(index int, selection *goquery.Selection) {
		if i, err := strconv.Atoi(selection.Text()); err == nil {
			if i < start {
				start = i
			}
			if i > end {
				end = i
			}
		}
	})
	return start, end
}

func getArticle(document *goquery.Document) []Article {
	articles := []Article{}

	document.Find("article.post").Each(func(index int, selection *goquery.Selection) {
		title := strings.TrimSpace(selection.Find("h1.post-title > a.post-title-link").Text())
		time := strings.TrimSpace(selection.Find("div.post-meta > span.post-time > time").Text())
		articles = append(articles, Article{title, time})
	})
	return articles
}

func main() {
	startTime := time.Now()
	url := "https://imsilence.github.io/"
	document, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	start, end := getPagination(document)
	articles := getArticle(document)

	var group sync.WaitGroup
	channel := make(chan Article, 1024)

	for i := start + 1; i <= end; i++ {
		group.Add(1)
		go func(i int, channel chan<- Article) {
			document, err := goquery.NewDocument(fmt.Sprintf("%s/page/%d/", url, i))
			if err != nil {
				fmt.Println(err)
			}
			particles := getArticle(document)
			for _, article := range particles {
				channel <- article
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
		id = i
		fmt.Println(i, ":", article)
	}

	for article := range channel {
		id++
		fmt.Println(id, ":", article)
	}
	fmt.Println("time:", time.Now().Sub(startTime))
}
