package cex

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Rate struct {
	currency string
	// rate     float64
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

	json := string(result)
	fmt.Printf(json)
	return &Rate{
		currency: json,
	}, nil
}
