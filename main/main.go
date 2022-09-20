package main

import (
	"fmt"
	"github.com/ykallan/requests"
)

func main() {
	headers := requests.Headers{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36",
	}
	url := "https://www.baidu.com/"
	resp := requests.Get(url, headers)

	// extract info success
	str1, ok := resp.Xpath("//title/text()").Get()
	if ok {
		fmt.Println(str1)
	} else {
		fmt.Println("没提取到信息")
	}

	// extract info empty
	str2, ok := resp.Xpath("//titlexx/text()").Get()
	if ok {
		fmt.Println(str2)
	} else {
		fmt.Println("没提取到信息")
	}

}