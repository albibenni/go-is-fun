package main

import (
	"fmt"
	"http2/cex"
)

type Curr struct {
	cur  string
	rate float64
}

func main() {
	//fmt.Printf("res: %v,   errors? %v", *res, err)
	var chBTC chan Curr = make(chan Curr)
	var chETH chan Curr = make(chan Curr)
	var chBCH chan Curr = make(chan Curr)
	go getCurrencies("BTC", chBTC)
	go getCurrencies("ETH", chETH)
	go getCurrencies("BCH", chBCH)
	btc := <-chBTC
	fmt.Printf("curr %v, rate %.2f \n", btc.cur, btc.rate)
}

func getCurrencies(currency string, ch chan Curr) (cur *string, rate *float64, err error) {

	res, err := cex.GetRate(currency)
	if err != nil {
		return nil, nil, err
	}

	currencies := fmt.Sprintf("%s | %s", res.CriptoCurr, res.RealCurr)
	ch <- Curr{cur: currencies, rate: res.Price}

	return &currencies, &res.Price, nil
}
