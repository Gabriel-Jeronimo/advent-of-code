package main

import (
	"day-1/trebuchet"
	"fmt"
)

func main() {
	sum, err := trebuchet.GetSum("data/input")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", sum)
}
