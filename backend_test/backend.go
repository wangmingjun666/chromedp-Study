package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"io/ioutil"
	"log"
)

func main() {
	//var err error

	// 连接远程服务器
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	var  body string
	if err := chromedp.Run(ctx,
		// reset
		chromedp.Emulate(device.Reset),

		// set really large viewport
		chromedp.EmulateViewport(1920, 2000),
		chromedp.Navigate(`http://baixarquadrinhos.com/Hq-Quadrinho/baixar-tirinha-calvin-e-haroldo-e-foi-assim-que-tudo-comecou-em-pdf-cbr-ou-ler-online/`),
		chromedp.OuterHTML("html",&body),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("test.html", []byte(body), 0644); err != nil {
		log.Fatal(err)
	}
}
