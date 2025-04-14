package main

import (
	"fmt"
	"http2/cex"
)

func main() {
	res, err := cex.GetRate("BTC")
	fmt.Printf("res: %v,   errors? %v", *res, err)
}
