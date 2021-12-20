package main

import (
	"fmt"
	"runtime"
	"sync"
)

//A “mutex” is a mutual exclusion lock. Mutexes allow us to lock our code
// so that only one goroutine can access that locked chunk of code at a time.

//think of mutex as like checking out a book, you can only have one person with a book at the time and you can only check out one book at a time.

//more info https://pkg.go.dev/sync#Mutex

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
