package mjj

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"qBot/pkg/config"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

var (
	ipua = `Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.13(0x17000d2a) NetType/WIFI Language/zh_CN"`
	ua   = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36`
)

func cdpex() {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.cnblogs.com/"),
		chromedp.WaitVisible(`#footer`, chromedp.ByID),
		chromedp.Nodes(`.//a[@class="post-item-title"]`, &nodes),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		data, _ := node.MarshalJSON()
		fmt.Println(string(data))
		fmt.Println(node.Children[0].NodeValue, ":", node.AttributeValue("href"))
	}
}

func MjjCdpMobile() ([][]string, error) {
	options := chromedp.DefaultExecAllocatorOptions[:]
	myoptions := []chromedp.ExecAllocatorOption{
		// chromedp.Flag("headless", false),
		chromedp.NoSandbox,
		// chromedp.UserAgent(ipua),
	}
	options = append(options, myoptions[:]...)

	var (
		ctx    context.Context
		cc     context.CancelFunc
		cancel context.CancelFunc
	)
	if config.GlobalConfig.GetString("devToolWsUrl") != "" {
		ctx, cc = chromedp.NewRemoteAllocator(context.Background(), config.GlobalConfig.GetString("devToolWsUrl"))
		defer cc()
	} else {
		ctx, cc = chromedp.NewExecAllocator(context.Background(), options...)
		defer cc()
	}

	ctx, cancel = chromedp.NewContext(
		ctx,
		// chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	tctx, tcancel := context.WithTimeout(
		ctx, 15*time.Second,
	)
	defer tcancel()
	defer chromedp.Stop()

	var datalist string

	err := chromedp.Run(tctx, mjjactions(&datalist))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(datalist))
	if err != nil {
		return nil, err
	}
	// var m = make(map[string]string, 8)
	var newslist = make([][]string, 8)
	doc.Find(`li>a`).Each(func(i int, doc *goquery.Selection) {
		title := doc.Contents().FilterFunction(func(i int, s *goquery.Selection) bool {
			return !s.Is("span")
		}).Text()
		title = strings.ReplaceAll(title, " ", "")
		title = strings.ReplaceAll(title, "\n", "")
		href, _ := doc.Attr("href")
		id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(href)
		if len(id) == 2 && id[1] != "" {
			newslist[i] = []string{id[1], title}
		}
	})
	return newslist, nil
}

func MjjCdp() {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(mjjurl),
		chromedp.WaitVisible(`#autopbn`, chromedp.ByID),
		chromedp.Nodes(`.//a[@class="s xst"]`, &nodes),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {

		id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(node.AttributeValue("href"))
		fmt.Println(node.Children[0].NodeValue, id[1])

	}
}

func mjjactions(datalist *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Emulate(device.IPhoneXR),
		chromedp.Navigate(mjjphoneurl),
		chromedp.WaitVisible(`#select_a`, chromedp.ByID),
		chromedp.OuterHTML(`html`, datalist),
	}
}

func Localmobile() {
	bytedata, err := ioutil.ReadFile("../../test/mobile.html")
	if err != nil {
		fmt.Println(err)
		panic("找不到文件")
	}
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(bytedata))
	doc.Find(`li>a`).Not(`span`).Each(func(i int, doc *goquery.Selection) {
		title := doc.Contents().FilterFunction(func(i int, s *goquery.Selection) bool {
			return !s.Is("span")
		}).Text()
		title = strings.ReplaceAll(title, " ", "")
		title = strings.ReplaceAll(title, "\n", "")
		href, _ := doc.Attr("href")
		id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(href)
		if len(id) == 2 && id[1] != "" {
			fmt.Println(id[1], title)
		}
	})

}
