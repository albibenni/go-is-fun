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

// Rate represents a cryptocurrency exchange rate with its timestamp
type Rate struct {
	CriptoCurr string    // The cryptocurrency symbol (e.g., BTC)
	RealCurr   string    // The real currency symbol (e.g., EUR)
	Price      float64   // The current exchange rate
	Timestamp  time.Time // The timestamp when the rate was recorded
}

// ApiCex represents the JSON response structure from the CEX.IO API
type ApiCex struct {
	Timestamp             string  `json:"timestamp"`             // Unix timestamp of the rate
	Low                   string  `json:"low"`                   // Lowest price in the last 24h
	High                  string  `json:"high"`                  // Highest price in the last 24h
	Last                  string  `json:"last"`                  // Last trade price
	Volume                string  `json:"volume"`                // Trading volume in the last 24h
	Volume30D             string  `json:"volume30d"`             // Trading volume in the last 30 days
	Bid                   float64 `json:"bid"`                   // Current highest bid price
	Ask                   int     `json:"ask"`                   // Current lowest ask price
	PriceChange           string  `json:"priceChange"`           // Price change in the last 24h
	PriceChangePercentage string  `json:"priceChangePercentage"` // Price change percentage in the last 24h
	Pair                  string  `json:"pair"`                  // Trading pair (e.g., "BTC:EUR")
}

// apiUrl is the base URL for the CEX.IO API ticker endpoint
const apiUrl = "https://cex.io/api/ticker/%s/EUR"

// GetRate fetches the current exchange rate for a given cryptocurrency from CEX.IO
// currency: The cryptocurrency symbol (e.g., "BTC")
// Returns a Rate struct containing the exchange rate information or an error if the request fails
func GetRate(currency string) (*Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http failed")
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
	if err != nil {
		return nil, err
	}
	return rate, nil
}

// apiCexToRate converts an ApiCex struct to a Rate struct
// Returns a Rate struct containing the exchange rate information or an error if conversion fails
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
