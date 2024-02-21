package main

import "censys/src"

func main() {
	// 这里输入你要查询的域名
	q := "198.23.246.120"
	// 提供你的 Censys API ID 和 Secret
	apiID := "c37adb8f-eba1-4650-9426-8f8c61762233"
	apiSecret := "sOkOmcofIxAvhPFMr7U0IY3K4ODUlZ9H"

	//src.SearchCensys(q, apiID, apiSecret)
	src.Ipv4Censys(q, apiID, apiSecret)
}
