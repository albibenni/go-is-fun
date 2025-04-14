package cex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Rate struct {
	CriptoCurr string
	RealCurr   string
	Price      float64
	Timestamp  time.Time
}

type ApiCex struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   int     `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}

const apiUrl = "https://cex.io/api/ticker/%s/EUR"

func GetRate(currency string) (*Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Http failed")
	}
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var apiCex ApiCex
	err = json.Unmarshal(result, &apiCex)
	if err != nil {
		return nil, err
	}
	rate, err := apiCex.apiCexToRate()
	return rate, nil
}

func (apiCex *ApiCex) apiCexToRate() (*Rate, error) {
	currencies := strings.Split(apiCex.Pair, ":")
	crypto, realCurr := currencies[0], currencies[1]
	price, err := strconv.ParseFloat(apiCex.Last, 64)

	if err != nil {
		return nil, err
	}
	epochInt, err := strconv.ParseInt(apiCex.Timestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	utcTime := time.Unix(epochInt, 0).UTC()
	return &Rate{
		CriptoCurr: crypto,
		RealCurr:   realCurr,
		Price:      price,
		Timestamp:  utcTime,
	}, nil
}
