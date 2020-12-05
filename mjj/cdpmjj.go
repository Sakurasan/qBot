package mjj

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	ipua = `Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.13(0x17000d2a) NetType/WIFI Language/zh_CN"`
	ua   = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36`
)

func cdpmjj0() {

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

	fmt.Println("get nodes:", len(nodes))
	// print titles
	for _, node := range nodes {
		fmt.Println(node.Children[0].NodeValue, ":", node.AttributeValue("href"))
	}
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

	fmt.Println("get nodes:", len(nodes))
	// print titles
	for _, node := range nodes {
		// data, _ := node.MarshalJSON()
		// fmt.Println(string(data))
		id := regexp.MustCompile(`&tid=(.*?)&`).FindStringSubmatch(node.AttributeValue("href"))
		fmt.Println(node.Children[0].NodeValue, id[1])

	}
}

func cdpmjjm() {

	options := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent(ipua),
	}
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)
	ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)

	ctx, cancel := chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	var mjjlist string
	// var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(mjjurl),
		chromedp.WaitVisible(`#dumppage`, chromedp.ByID),
		chromedp.OuterHTML(`document.querySelector("body > div.threadlist")`, &mjjlist, chromedp.ByJSPath),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mjjlist)
	// for _, node := range nodes {
	// 	data, _ := node.MarshalJSON()
	// 	fmt.Println(string(data))
	// 	// fmt.Println(node.Children[0].NodeValue)

	// }
}

func handless() {
	var nodes []*cdp.Node
	options := []chromedp.ExecAllocatorOption{
		chromedp.NoSandbox,
		chromedp.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.13(0x17000d2a) NetType/WIFI Language/zh_CN"),
	}
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)
	c, cc := chromedp.NewExecAllocator(context.Background(), options...)
	defer cc()
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	err := chromedp.Run(ctx, mjjactions(nodes))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, node := range nodes {
		data, _ := node.MarshalJSON()
		fmt.Println(string(data))
		// fmt.Println(node.Children[0].NodeValue)

	}

}

func mjjactions(nodes []*cdp.Node) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(mjjurl),
		chromedp.WaitVisible(`#dumppage`),
		chromedp.Nodes(`.//li`, &nodes),
	}
}
