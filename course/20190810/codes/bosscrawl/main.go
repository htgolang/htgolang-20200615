package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.zhipin.com/job_detail/?query=go&city=100010000&industry=&position="

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("authority", "www.zhipin.com")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36 Avast/75.1.1528.101")
	request.Header.Set("path", "/job_detail/?query=go&city=100010000&industry=&position=")
	// 替换cookie
	request.Header.Set("cookie", "lastCity=101110100; __c=1565432457; __g=-; __l=l=%2Fwww.zhipin.com%2F&r=https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3Da7lQQ1jN7IG3MiNsA_yal4yLDMw5B8fnB0HakuA9Jx5dBByjTLCwvg9t_Sfm5YEn%26wd%3D%26eqid%3De92875d80006a11c000000025d4e9a82; _uab_collina=156543245684559605417148; __zp_stoken__=44fc5Z1oLOq%2BBfH4164sSs5m6erA1gAG%2Fn4wPZGNfslS%2FRiMR749rHW%2BjtILv3nil51nGlyKCnILFsajM4u1mUjn%2BA%3D%3D; __a=83666724.1565432457..1565432457.13.1.13.13")
	client := &http.Client{}

	response, _ := client.Do(request)
	document, _ := goquery.NewDocumentFromResponse(response)
	// cxt, _ := ioutil.ReadFile("job.html")
	// reader := bytes.NewReader(cxt)

	// document, _ := goquery.NewDocumentFromReader(reader)
	// fmt.Println(document.Html())
	document.Find("div.job-primary").Each(func(index int, selection *goquery.Selection) {
		fmt.Println("--------------------------------------------------")
		//fmt.Println(selection.Html())
		fmt.Println(selection.Find("div.info-company > div.company-text > h3 > a").Text())

		tagA := selection.Find("div.info-primary > h3.name > a")
		fmt.Println(tagA.Find("div.job-title").Text())
		fmt.Println(tagA.Find("span").Text())
	})
}
