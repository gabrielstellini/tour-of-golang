package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	result := 1.0
	for decimalPlaces := 1; decimalPlaces <= 10; decimalPlaces++ {
		result -= (result * result - x) / (2 * result)
	}
	return result
}

func main() {
	fmt.Println(Sqrt(2))
}
