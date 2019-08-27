# chromedp-Study
chromedp Study Chromedp 学习


#### Install ChromeDp
``` 
go get -u github.com/chromedp/chromedp
```

### Chrome无界面容器
``` 
docker pull chromedp/headless-shell
docker run -d -p 9222:9222 --rm --name headless-shell chromedp/headless-shell

测试是否正常
http://127.0.0.1:9222/json
```

### 获取加载完毕的网页
注释：这个作者更新太勤快了吧，API更新的大变样了
``` 
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
```