package main

import "fmt"

type Vertex struct {
	X ComplexX
	Y int
}

type ComplexX struct {
	Z string
}

func main() {
	vertex := Vertex{
		X: ComplexX{Z: "asd"},
		Y: 2,
	}

	p := &vertex
	p.Y = 22

	fmt.Println(vertex)
}
