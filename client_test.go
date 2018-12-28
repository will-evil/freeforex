package freeforex

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestPreparePairs(t *testing.T) {
	pairs := []string{"usdeur", "USDCHF", "USDARS ", "  UsDAWg"}
	expectedPairs := []string{"USDEUR", "USDCHF", "USDARS", "USDAWG"}
	pairs = preparePairs(pairs)

	for index, expectedPair := range expectedPairs {
		if expectedPair != pairs[index] {
			t.Errorf("pair is failed. Expected %s, got %s", expectedPair, pairs[index])
		}
	}
}

func TestGetRatesResponse(t *testing.T) {
	const (
		respRatesCount = 1
		respRatesKey   = "EURUSD"
		respRate       = 1.140334
		respTimestamp  = 1545922455356
	)

	jsonFile, err := os.Open("var/live_eurusd.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		err = jsonFile.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	body, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Error(err)
		return
	}
	res, err := getRatesResponse(body)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Code != successCode {
		t.Errorf("response code is failed. Expected %d, got %d", successCode, res.Code)
	}
	if len(res.Rates) != respRatesCount {
		t.Errorf("response rates count is failed. Expected %d, got %d", respRatesCount, len(res.Rates))
	}
	rate, ok := res.GetRate(respRatesKey)
	if !ok {
		t.Errorf("key '%s' not found in response rates", respRatesKey)
	}
	if rate != respRate {
		t.Errorf("response rate is failed. Expected %f, got %f", respRate, rate)
	}
	timestamp, _ := res.GetTimestamp(respRatesKey)
	if timestamp != respTimestamp {
		t.Errorf("response rate timestamp is failed. Expected %d, got %d", respTimestamp, timestamp)
	}
}

func TestGetRatesResponse2(t *testing.T) {
	const (
		respCode    = 1001
		respMessage = "'pairs' parameter is required"
	)

	jsonFile, err := os.Open("var/live_1001.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		err = jsonFile.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	body, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Error(err)
		return
	}
	var res *RatesResponse
	if err := json.Unmarshal(body, &res); err != nil {
		t.Error(err)
		return
	}

	if res.Code != respCode {
		t.Errorf("response code is failed. Expected %d, got %d", respCode, res.Code)
	}
	if res.Message != respMessage {
		t.Errorf("response message is failed. Expected %s, got %s", respMessage, res.Message)
	}
}
