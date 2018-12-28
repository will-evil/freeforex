package freeforexapi

type Rate struct {
	Rate float64 `json:"rate"`
	Timestamp int `json:"timestamp"`
}

type RatesResponse struct {
	BaseResponse
	Rates map[string]Rate
}

