package tools

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

func AnalysisHtml(url string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	var body string
	if err := chromedp.Run(ctx,
		// reset
		chromedp.Emulate(device.Reset),

		// set really large viewport
		chromedp.EmulateViewport(1920, 2000),
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &body),
	); err != nil {
		return "", err
	}
	return body, nil
}
