package main

import (
	"fmt"
)

//uncomment the code below to replace c := ... to break the deadlock and see the output

func main() {
	c := make(chan int)
	//c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
