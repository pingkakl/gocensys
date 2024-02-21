package src

type ResponseData struct {
	Result struct {
		Hits []struct {
			IP       string `json:"ip"`
			Services []struct {
				Port int `json:"port"`
			} `json:"services"`
		} `json:"hits"`
	} `json:"result"`
}

type Ipv4Data struct {
	Result struct {
		IP  string `json:"ip"`
		Dns struct {
			Names   []string          `json:"names"`
			Records map[string]Record `json:"records"`
		} `json:"dns"`
	} `json:"result"`
}

type Record struct {
	RecordType string `json:"record_type"`
	ResolvedAt string `json:"resolved_at"`
}
