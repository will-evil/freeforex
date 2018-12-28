package freeforexapi

import (
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
	}

	if res.Code != successCode {
		t.Errorf("response code is failed. Expected %d, got %d", successCode, res.Code)
	}
}
