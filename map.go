package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex, 10)
	m["Bell Labs"] = Vertex{
		40.490, -74.234,
	}
	fmt.Println(m["Bell Labs"])
}
