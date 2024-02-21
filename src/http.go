package src

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const url = "https://search.censys.io/api/v2/"

func post(s, apiID, apiSecret string, reqBody []byte) []byte {
	req, err := http.NewRequest("POST", url+s, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("创建请求时出错：", err)
		return nil
	}
	req.SetBasicAuth(apiID, apiSecret)
	req.Header.Set("content-type", "application/json")
	fmt.Println(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("发出请求时出错：", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应正文时出错：", err)
		return nil
	}
	return body
}

func get(q, apiID, apiSecret string) []byte {

	req, err := http.NewRequest("GET", url+q, nil)
	if err != nil {
		fmt.Println("url:", req)
		return nil
	}
	req.SetBasicAuth(apiID, apiSecret)
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("发出请求时出错：", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应正文时出错：", err)
		return nil
	}
	return body
}
