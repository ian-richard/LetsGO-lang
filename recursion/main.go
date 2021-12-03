package main

import "fmt"

func main() {
	ex := 4 * 3 * 2 * 1
	n := factorial(4)
	fmt.Println(n == ex)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}