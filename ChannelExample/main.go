package main

import (
	"fmt"
)

func main() {
	xc := make(chan int)
	go func() {
		//put 100 values into the channel
		for i := 1; i <= 100; i++ {
			xc <- i
		}
		//close the channel
		close(xc)
	}()
	//range over the channel and print the values
	for x := range xc {
		fmt.Println(x)
	}
	fmt.Println("about to exit")
}
