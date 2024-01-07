package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	ch <- 3
	i := <-ch
	fmt.Println(<-ch)
	fmt.Println(i)
}
