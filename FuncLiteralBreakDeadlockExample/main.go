package main

import (
	"fmt"
	"sync"
)

//uncomment the code below to break the deadlock and see the output

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	// go func() {
	// 	c <- 42
	// }()

	fmt.Println(<-c)
}
