package main

import (
	"fmt"
	//"sync"
)

func main() {
	c := make(chan int)

	//write a program which launches 10 go routines
	//each go routine adds 10 numbers to a channel
	//pull the values off the channel and print them out

	for j := 0; j < 10; j++ {

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()

	}

	for k := 1; k <= 100; k++ {
		fmt.Println(k, <-c)
	}
	fmt.Println("about to exit")

}
