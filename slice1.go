package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(a)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, false},
		{11, false},
		{13, true},
	}

	fmt.Println(s)

}
