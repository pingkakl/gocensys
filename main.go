package main

import "censys/src"

func main() {
	// 这里输入你要查询的域名
	q := "www.baidu.com"
	// 提供你的 Censys API ID 和 Secret
	apiID := ""
	apiSecret := ""

	src.SearchCensys(q, apiID, apiSecret)
}
