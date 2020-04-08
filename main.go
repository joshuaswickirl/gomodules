package main

import (
	"fmt"

	"rsc.io/quote"
	quotev2 "rsc.io/quote/v2"
	quotev3 "rsc.io/quote/v3"
)

func main() {
	// https://github.com/rsc/quote
	fmt.Println(quote.Hello())
	fmt.Println(quotev2.HelloV2())
	fmt.Println(quotev3.HelloV3())
}
