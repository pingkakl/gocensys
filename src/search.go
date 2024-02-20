package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SearchCensys(q, apiID, apiSecret string) {
	query := "10"
	vhosts := "EXCLUDE"
	sort := "RELEVANCE"
	url := "https://search.censys.io/api/v2/hosts/search"

	/*
		将reqBody 的map类型的数据结构转换为json格式的字符串

		展示：
		{
			"q":"your_q",
			"per_page":"your_per_page",
			"virtual_hosts":"your_vhosts",
			"sort":"your_sort",
		}
	*/
	reqBody, err := json.Marshal(map[string]string{
		"q":             q,
		"per_page":      query,
		"virtual_hosts": vhosts,
		"sort":          sort,
	})
	if err != nil {
		fmt.Println("编组请求正文时出错：", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("创建请求时出错：", err)
		return
	}
	req.SetBasicAuth(apiID, apiSecret)
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("发出请求时出错：", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应正文时出错：", err)
		return
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("解组响应正文时出错：", err)
		return
	}

	ipPortMap := make(map[string][]string)
	for _, hit := range data.Result.Hits {
		for _, service := range hit.Services {
			ipPortMap[hit.IP] = append(ipPortMap[hit.IP], fmt.Sprintf("%d", service.Port))
		}
	}

	for ip, ports := range ipPortMap {
		fmt.Printf("IP: %s, Ports: %s\n", ip, strings.Join(ports, ","))
	}

}
