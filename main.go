package main

import (
	"fmt"

	. "github.com/bjartek/overflow"
)

func main() {
	_ = WithNetwork("mainnet")
	fmt.Println("Hello, overflow!")
}
