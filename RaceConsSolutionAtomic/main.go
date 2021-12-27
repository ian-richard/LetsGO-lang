package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			r := atomic.LoadInt64(&incrementer)
			fmt.Println(r)
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}

// //docs: - https://pkg.go.dev/sync/atomic 

// //Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

// // These functions require great care to be used correctly. Except for special, low-level applications, synchronization is better done with channels or the facilities of the sync package. Share memory by communicating; don't communicate by sharing memory. 
