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
	mjjphoneurl = `https://hostloc.com/forum.php?mod=forumdisplay&fid=45&orderby=dateline&mobile=2`
	mjjurl      = "https://hostloc.com/forum.php?mod=forumdisplay&fid=45&filter=author&orderby=dateline"

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
	treg        = regexp.MustCompile(`<tbody id="normalthread_(.*)</tbody>`)
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
	if result.StatusCode != http.StatusOK {
		fmt.Println(result.StatusCode, result.Status)
		return
	}
	// bodyReader := bufio.NewReader(result.Body)
	// bodyReader
	// ts:= treg.FindAllStringSubmatch(bodyReader.)

	doc, err := goquery.NewDocumentFromResponse(result)
	if err != nil {
		log.Println(err)
		return
	}
	// id="separatorline"
	ndoc := doc.Find("div[class=\"boardnav\"]").Find("div[class=\"bm_c\"]")
	// ndoc.Find("tbody").Each(func(i int, doc *goquery.Selection) {
	// 	title := doc.Find("tr").Find("th[class=\"new\"]").Find("a[class=\"s xst\"]").Text()
	// 	href, _ := doc.Find("tr").Find("th[class=\"new\"]").Find("a[class=\"s xst\"]").Attr("href")
	// 	id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(href)
	// 	fmt.Println(title, id)
	// })
	fmt.Println(len(ndoc.Find("tbody").Nodes))
	ndoc.Find("tbody").Slice(4, 9).Each(func(i int, doc *goquery.Selection) {
		fmt.Println(i)
		title := doc.Find("tr").Find("th[class=\"new\"]").Find("a[class=\"s xst\"]").Text()
		href, _ := doc.Find("tr").Find("th[class=\"new\"]").Find("a[class=\"s xst\"]").Attr("href")
		id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(href)
		fmt.Println(title, id)
	})

}

func browser2() {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(10) * time.Second,
		}).Dial,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: false,
		MaxIdleConns:      10,
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	req, _ := http.NewRequest(http.MethodGet, mjjphoneurl, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("User-Agent", ipua)

	result, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	if result.StatusCode != http.StatusOK {
		fmt.Println(result.StatusCode, result.Status)
		return
	}
	// bodyReader := bufio.NewReader(result.Body)
	// bodyReader
	// ts:= treg.FindAllStringSubmatch(bodyReader.)

	doc, err := goquery.NewDocumentFromResponse(result)
	if err != nil {
		log.Println(err)
		return
	}
	// /html/body/div[1]/ul/li[1]
	ndoc := doc.Find("div[class=\"threadlist\"]").Find("ul")
	// ndoc.Find("tbody").Each(func(i int, doc *goquery.Selection) {
	// 	title := doc.Find("tr").Find("th[class=\"new\"]").Find("a[class=\"s xst\"]").Text()
	// 	href, _ := doc.Find("tr").Find("th[class=\"new\"]").Find("a[class=\"s xst\"]").Attr("href")
	// 	id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(href)
	// 	fmt.Println(title, id)
	// })
	ndoc.Find("li").Each(func(i int, doc *goquery.Selection) {
		href, _ := doc.Find("a").Attr("href")
		id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(href)
		fmt.Println(doc.Find("a").Text(), id[1])
	})
	// fmt.Println(ndoc.Html())

}
