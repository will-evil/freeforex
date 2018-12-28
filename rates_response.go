package freeforex

// Rate is structure for rates items.
type Rate struct {
	Rate      float64 `json:"rate"`
	Timestamp int     `json:"timestamp"`
}

// RatesResponse is structure for response from https://www.freeforexapi.com/api/live?pairs=${your_pair}.
type RatesResponse struct {
	BaseResponse
	Rates map[string]Rate
}

// GetRate return rate by pair.
// Returns as the first parameter pair rate.
// And as the second parameter, the indicator of existence pair in rates response.
func (rr RatesResponse) GetRate(pair string) (float64, bool) {
	pair = preparePair(pair)
	if _, ok := rr.Rates[pair]; !ok {
		return 0, false
	}

	return rr.Rates[pair].Rate, true
}

// GetRate return timestamp by pair.
// Returns as the first parameter pair timestamp.
// And as the second parameter, the indicator of existence pair in rates response.
func (rr RatesResponse) GetTimestamp(pair string) (int, bool) {
	pair = preparePair(pair)
	if _, ok := rr.Rates[pair]; !ok {
		return 0, false
	}

	return rr.Rates[pair].Timestamp, true
}
