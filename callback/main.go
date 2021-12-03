package main

import "fmt"

func main () {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi) - 1]
	}

	callbackInAction := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(callbackInAction)
	
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n--
	return n
}

//a real life example might be a salary calculation func, which is then passed another function to add or reduce pay based on number of holiday days taken/or not
//ie.
//the function that is passed to the salary calculation func is a callback function
//the callback function is then called by the salary calculation func



