package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64
	

	goRoutines := 100
	wg.Add(goRoutines)

	var s []int

	for i := 0; i < goRoutines; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			//NB the numbers are not printed in order
			//append the value of incrementer to the slice, then check length of slice to confirm it's 100
			s = append(s, int(incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
	//print length of slice
	fmt.Println("len of slice", len(s))
}

//docs: - https://pkg.go.dev/sync/atomic 

//Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

// These functions require great care to be used correctly. Except for special, low-level applications, synchronization is better done with channels or the facilities of the sync package. Share memory by communicating; don't communicate by sharing memory. 
