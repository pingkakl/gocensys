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
