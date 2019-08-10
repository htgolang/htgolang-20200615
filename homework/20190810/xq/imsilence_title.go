package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func getTitle(urls string) {

	fmt.Println("============================", urls)


	client := &http.Client{}
	request, err := http.NewRequest("GET", urls, nil)

	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
	res, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}



	defer res.Body.Close()

	document, _ := goquery.NewDocumentFromReader(res.Body)


	document.Find("header.post-header").Each(func(i int, selection *goquery.Selection) {

		title := selection.Find("h1.post-title > a").Text()
		title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\n", "", -1)
		fmt.Println(title)
	})

}

func main()  {
	//urls := "https://imsilence.github.io/page/2/"
	urls := "https://imsilence.github.io/"

	for i := 1; i<=8; i++ {
		if i<= 1 {

			getTitle(urls)

		}else  {

			newurls := urls + "page/" + strconv.Itoa(i) +"/"
			getTitle(newurls)

		}
	}
}