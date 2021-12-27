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

	var mutex sync.Mutex

	for i := 0; i < goRoutines; i++ {
		go func() {
			mutex.Lock()
			fmt.Println("value of incrementer", incrementer)
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", incrementer)
}
