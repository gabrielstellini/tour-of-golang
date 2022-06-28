package main

import (
	"fmt"
	"time"
)

func main() {
	switch {
	case time.Now().Day() < 12:
		fmt.Println("Good morning!")
	case time.Now().Day() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
