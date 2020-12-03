package mjj

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
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
		fmt.Println(node.Children[0].NodeValue)

	}
}
