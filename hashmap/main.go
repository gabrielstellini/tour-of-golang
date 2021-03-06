package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var mapLiteral = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var mapLiteralInferredTypes = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
	// "Google2":    {37.42202, -122.08408, -33.2}, Won't work!
}

var m map[string]Vertex

func main() {
	// fmt.Println(mapLiteral)
	// fmt.Println(mapLiteralInferredTypes)
	// updateDemo()
	// setDeleteDemo()
	fmt.Println(WordCount("Hello World Hello"))

}

func updateDemo() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

}

func setDeleteDemo() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	wordMap := make(map[string]int)

	for _, value := range words {
		wordMap[value] = wordMap[value] + 1
	}

	return wordMap
}
