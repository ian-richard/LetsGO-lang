package main

import "fmt"

//use 'var' if you ever want to delcare a variable outside of the function body (scope is the global doc)
var n = 19;

var z int

func main() {
	x := 42;
	y := x + 42
    fmt.Println("I'm your main")
    fmt.Println(x, y, n)

	foo();
}

func foo(){
	fmt.Println("again ", n )
}