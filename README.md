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


### tools 目录下包装的方法
- 解析网页
    ``` 
    AnalysisHtml(url string) (string, error)
    ```
    
### 如何并发的去获取呢？
``` 
// 这个是伪语言啊！不要直接用啊
	// 这并不似最佳实践

	tagUrl := make(chan string)  // 假设里面有内容是目标url

	ints := make(chan int, 10)  // 控制并发数量
	ints <- 1
	for {
		go func(urlList chan string) {
			for {
				select {
				case urlList := <- urlList:
					s, e := tools.AnalysisHtml(urlList)
					if e == nil {
						//将s写入解析的channel中
					}
				}
			}
		}(tagUrl)
	}
	<-ints
```