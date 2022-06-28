package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for yIndex := 0; yIndex < dy; yIndex++ {
		result[yIndex] = make([]uint8, dx)

		for xIndex := 0; xIndex < dx; xIndex++ {
			result[yIndex][xIndex] = uint8((xIndex + yIndex) / 2)
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
