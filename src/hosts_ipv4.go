package src

import (
	"encoding/json"
	"fmt"
)

func Ipv4Censys(q, apiID, apiSecret string) {

	q = "hosts/" + q

	body := get(q, apiID, apiSecret)

	var data Ipv4Data
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("解组响应正文时出错：", err)
		return
	}

	res := data.Result.IP
	dns := data.Result.Dns.Names

	for _, record := range data.Result.Dns.Records {
		fmt.Printf("IP: %s, Domain: %s, Type: %s, Time: %s\n", res, dns, record.RecordType, record.ResolvedAt)
	}

}
