package freeforex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/will-evil/httphelp"
	"net/url"
	"strings"
)

const apiUrl = "https://www.freeforex.com/api/live"

const successCode = 200

// Client is structure that provide access to API methods.
type Client struct {
}

// Live do live request and return link to RatesResponse and error object.
func (c Client) Live(pairs []string) (*RatesResponse, error) {
	if len(pairs) == 0 {
		return nil, errors.New("'pairs' parameter is required")
	}
	pairs = preparePairs(pairs)

	helper := httphelp.GetHelper()
	body, err := helper.GetBody(apiUrl, url.Values{"pairs": pairs})
	if err != nil {
		return nil, err
	}

	return getRatesResponse(body)
}

// Rate return rate for pair.
func (c Client) Rate(pair string) (float64, error) {
	pair = strings.TrimSpace(strings.ToUpper(pair))

	res, err := c.Live([]string{pair})
	if err != nil {
		return 0, err
	}

	if rate, ok := res.Rates[pair]; ok {
		return rate.Rate, nil
	}

	return 0, fmt.Errorf("rate for pair '%s' not found", pair)
}

func preparePairs(pairs []string) []string {
	var newPairs []string

	for _, pair := range pairs {
		newPairs = append(newPairs, preparePair(pair))
	}

	return newPairs
}

func preparePair(pair string) string {
	return strings.TrimSpace(strings.ToUpper(pair))
}

func getRatesResponse(body []byte) (*RatesResponse, error) {
	var resp *RatesResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	} else if resp.Code != successCode {
		return nil, errors.New(resp.Message)
	}

	return resp, nil
}
