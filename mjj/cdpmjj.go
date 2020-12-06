package mjj

import (
	"context"
	"fmt"
	"log"
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

func cdpmjjex() {

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

func cdpmjjmobile() (map[string]string, error) {
	options := chromedp.DefaultExecAllocatorOptions[:]
	myoptions := []chromedp.ExecAllocatorOption{
		// chromedp.Flag("headless", false),
		chromedp.NoSandbox,
		// chromedp.UserAgent(ipua),
	}
	options = append(options, myoptions[:]...)
	ctx, cc := chromedp.NewExecAllocator(context.Background(), options...)
	defer cc()

	ctx, _ = chromedp.NewContext(
		ctx,
		// chromedp.WithLogf(log.Printf),
	)
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
	var m = make(map[string]string, 8)
	doc.Find(`li>a`).Each(func(i int, doc *goquery.Selection) {
		title := doc.Text()
		href, _ := doc.Attr("href")
		id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(href)
		if len(id) == 2 && id[1] != "" {
			m[id[1]] = title
		}
	})
	return m, nil
}

func cdpmjj() {

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

// func handless() {
// 	var nodes []*cdp.Node
// 	options := []chromedp.ExecAllocatorOption{
// 		chromedp.NoSandbox,
// 		// chromedp.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.13(0x17000d2a) NetType/WIFI Language/zh_CN"),
// 	}
// 	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)
// 	c, cc := chromedp.NewExecAllocator(context.Background(), options...)
// 	defer cc()
// 	ctx, cancel := chromedp.NewContext(c)
// 	defer cancel()

// 	err := chromedp.Run(ctx, )
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, node := range nodes {
// 		data, _ := node.MarshalJSON()
// 		fmt.Println(string(data))
// 		// fmt.Println(node.Children[0].NodeValue)

// 	}

// }

func mjjactions(datalist *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Emulate(device.IPhoneXR),
		chromedp.Navigate(mjjphoneurl),
		chromedp.WaitVisible(`#select_a`, chromedp.ByID),
		chromedp.OuterHTML(`html`, datalist),
	}
}
