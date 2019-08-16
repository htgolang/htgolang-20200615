package main

import (
	"fmt"
	"net/http"
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
	return fmt.Sprintf("%s - %s", a.Name, a.Time)
}

func GetPaging(document *goquery.Document) (int, int) {
	start, end := 1, -1

	document.Find("nav.pagination .page-number").Each(func(index int, selection *goquery.Selection) {
		// 循环获取页数，且将页数转换为int类型
		if i, err := strconv.Atoi(selection.Text()); err == nil {
			if i <= start {
				start = i
			}

			if i > end {
				end = i
			}
		}
	})

	return start, end

}

func GetArticle(document *goquery.Document) []Article {
	// 定义Article类型的切片，存放读取到的数据
	articles := []Article{}

	document.Find("article.post").Each(func(index int, selection *goquery.Selection) {
		title := strings.TrimSpace(selection.Find("h1.post-title > a.post-title-link").Text())
		time := strings.TrimSpace(selection.Find("div.post-meta > span.post-time > time").Text())
		articles = append(articles, Article{title, time})
	})

	return articles
}

// 在线从浏览器中筛选数据
func main() {
	startTime := time.Now()

	url := "https://imsilence.github.io/"

	// 自定义GET的方法获取资源
	request, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	document, _ := goquery.NewDocumentFromResponse(response)

	// 分别赋值开始和结束的页数
	start, end := GetPaging(document)
	// 将第一页的数据赋值给articles
	articles := GetArticle(document)

	// 定义group信号量
	var group sync.WaitGroup
	channel := make(chan Article, 1024)

	// 循环存放第二页以后的数据
	for i := start + 1; i <= end; i++ {
		group.Add(1)
		// 仅定义写入信号量
		go func(i int, channel chan<- Article) {
			document, err := goquery.NewDocument(fmt.Sprintf("%s/page/%d/", url, i))
			if err != nil {
				fmt.Println(err)
			}
			particles := GetArticle(document)
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
	// 读取第一页的数据
	for i, article := range articles {
		id = i
		fmt.Println(i, ":", article)
	}

	// 读取第二页以后的数据
	for article := range channel {
		id++
		fmt.Println(id, ":", article)
	}

	fmt.Println("time:", time.Now().Sub(startTime))

}
