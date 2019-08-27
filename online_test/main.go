package main

import (
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
)

func main() {
	var devToolWsUrl string
	flag.StringVar(&devToolWsUrl, "devtools-ws-url", "ws://127.0.0.1:9222/devtools/page/66FB4CCF8B2460C8BE092DA4D92E00D7", "DevTools Websocket URL")
	flag.Parse()

	actxt, cancelActxt := chromedp.NewRemoteAllocator(context.Background(), devToolWsUrl)
	defer cancelActxt()

	ctxt, cancelCtxt := chromedp.NewContext(actxt) // create new tab
	defer cancelCtxt()                             // close tab afterwards

	var body string
	if err := chromedp.Run(ctxt,
		chromedp.Navigate("https://www.baidu.com"),
		chromedp.WaitVisible("#logo_homepage_link"),
		chromedp.OuterHTML("html", &body),
	); err != nil {
		log.Fatalf("Failed getting body of duckduckgo.com: %v", err)
	}

	if err := ioutil.WriteFile("baidu.html", []byte(body), 0644); err != nil {
		log.Fatal(err)
	}
}
