package mjj

import (
	"container/list"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	mjjurl = "https://hostloc.com/forum.php?mod=forumdisplay&fid=45&filter=author&orderby=dateline"

	headers = map[string]string{
		"Accept-Encoding":           `gzip, deflate, br`,
		"Accept-Language":           `zh-CN,zh;q=0.9`,
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Cache-Control":             "no-cache",
		"Connection":                "keep-alive",
	}
	url_hostloc = "https://www.hostloc.com/forum.php?mod=forumdisplay&fid=45&filter=author&orderby=dateline"
	MJJlist     = list.New()
)

func initlist() {
	MJJlist.Init()
	for i := 0; i < 5; i++ {
		MJJlist.PushFront(i)
	}
	fmt.Println("mjj")
	fmt.Println(MJJlist.Front().Value)
	for e := MJJlist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

func browser() {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(10) * time.Second,
		}).Dial,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: false,
		MaxIdleConns:      10,
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	req, _ := http.NewRequest(http.MethodGet, mjjurl, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	result, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	doc, _ := goquery.NewDocumentFromResponse(result)
	// id="separatorline"
	doc.Find("div[class=\"boardnav\"]").Find("div[class=\"bm_c\"]").Find("tbody").Eq(4).Each(func(i int, doc *goquery.Selection) {
		title := doc.Find("tr").Find("th[class=\"new\"]").Find("a[class=\"s xst\"]").Text()
		href, _ := doc.Find("tr").Find("th[class=\"new\"]").Find("a[class=\"s xst\"]").Attr("href")
		id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(href)
		fmt.Println(title, id[1])
	})

}
