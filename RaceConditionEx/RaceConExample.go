package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	goRoutines := 100
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			fmt.Println("value of incrementer", incrementer)
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", incrementer)
}
