package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

var bastUrl = "https://imsilence.github.io/"

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var url string
	for i := 1; i < 8; i++ {
		fmt.Println("=============第", i, "页=============")
		if i != 1 {
			url = bastUrl + "page/" + strconv.Itoa(i)
		} else {
			url = bastUrl
		}

		request, err := http.NewRequest("GET", url, nil)
		check(err)

		client := &http.Client{}
		res, err := client.Do(request)
		check(err)
		defer res.Body.Close()

		document, err := goquery.NewDocumentFromReader(res.Body)
		check(err)

		document.Find("header.post-header").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(strings.TrimSpace(selection.Find("h1.post-title > a").Text()))
		})
	}
}
