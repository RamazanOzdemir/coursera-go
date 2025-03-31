package main

import (
	"fmt"
	"math"
)

func main() {
	var floatNumber float64


	fmt.Print("Please, Enter a float number: ")

	fmt.Scan(&floatNumber)

	truncatedNumber :=int(math.Trunc(floatNumber))

	fmt.Printf("Truncated number: %d\n", truncatedNumber)
}