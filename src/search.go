package src

import (
	"encoding/json"
	"fmt"
	"strings"
)

func SearchCensys(q, apiID, apiSecret string) {
	query := "10"
	vhosts := "EXCLUDE"
	sort := "RELEVANCE"
	s := "hosts/search"

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

	body := post(s, apiID, apiSecret, reqBody)

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
