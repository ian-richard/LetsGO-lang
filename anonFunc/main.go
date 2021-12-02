package main

import "fmt"

func main () {
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("done")
foo()()
}
//regular function return a anonymous function
func foo() func() {
	return func() {
		for i := 10; i <= 20; i++ {
			fmt.Println(i)
		}
	}
}

