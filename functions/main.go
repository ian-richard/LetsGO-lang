package main

import "fmt"

func main(){
	fmt.Println(sumSlice([]int{1,2,3,4}...))
	}

func sumSlice(numbs ...int) int {
	result := 0
	for _, num := range numbs {
		result += num
	}
	return result
}

func foo(num int) int {
	return num
	}

func bar (num int, s string) (int, string) {
	return num, s
	}

