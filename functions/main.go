package main

import "fmt"

func main(){
	fmt.Println(foo(42))
	fmt.Println(bar(42, "hello"))
	}

func foo(num int) int {
	return num
	}

func bar (num int, s string) (int, string) {
	return num, s
	}

