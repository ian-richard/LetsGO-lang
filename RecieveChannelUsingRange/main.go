package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}

func gen() <-chan int {
	c := make(chan int)
	//anon func (func literal)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
