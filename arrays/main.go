package main

import (
	"fmt"
	"strings"
)

/*
 len() - the length of a slice tells you how many elements it contains.
 cap() - The capacity is the size of the slice's underlying array
*/

func main() {
	// getAndSetValues()
	// sliceValues()
	// unlimitedSlicing()
	// structLiteral()
	// duplicateArray()
	// countArray()
	// defaultValue()
	// dynamicArray()
	// twoDimensions()
	// appendDemo()
	looping()
}

func looping() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// both index and value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}


	pow2 := make([]int, 10)
	for i := range pow2 { // only index
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 { //ignored index, only value
		fmt.Printf("%d\n", value)
	}
}

func appendDemo() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func twoDimensions() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func dynamicArray() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func defaultValue() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func countArray() {
	s := [20]int{2, 3, 5, 7, 11, 13}
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func getAndSetValues() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func sliceValues() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Print(primes)
	var s []int = primes[1:4] // only a reference to the underlying array
	fmt.Println(s)

	s[1] = 999

	fmt.Println(primes)
}

func duplicateArray() {
	src := []string{"hello", "world"}
	dst := make([]string, len(src))

	copy(dst, src)

	dst[0] = "Updated"

	fmt.Println(src)
	fmt.Println(dst)
}

func unlimitedSlicing() {
	primes := []int{2, 3, 5, 7, 11, 13}
	noLower := primes[:4]
	noUpper := primes[3:]
	noLimit := primes[:]

	printSlice(primes)
	printSlice(noLower)
	printSlice(noUpper)
	printSlice(noLimit)

	primes = primes[2:3]
	printSlice(primes)

	primes = primes[0:cap(primes)]
	printSlice(primes)

	primes = primes[:0]
	printSlice(primes)

	primes = primes[:4] // extend the length
	printSlice(primes)

	primes = primes[2:] // drop the first 2 values
	printSlice(primes)
}

func structLiteral() {
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
