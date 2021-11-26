package main

import "fmt"

//use 'var' if you ever want to delcare a variable outside of the function body (scope is the global doc)
var n = 19;



func main() {
	// x := 42;
	// y := x + 42
    // fmt.Println("I'm your main")
    // fmt.Println(x, y, n)

	foo();
}

// var x int = 42	
var y string = "jb"
var z bool = true

//create new type of variable Spanner
type spanner int
//create new variable of type spanner with identifier x
var x spanner

//new var with type int
var g int



func foo(){
	// x := 42
	// y := "jjb"
	// z := true
	// s := fmt.Sprintf("%v\t%v\t%v",x, y, z )
	// fmt.Println(s)
	// x = 42
	// fmt.Printf("%T\n", x)
	// fmt.Println(x)
	g = 10000
	// fmt.Println(g)
	// g = int(x)
	// fmt.Println(g)

	//print g in decimal
	fmt.Printf("%d\n", g)
	//print g in binary 
	fmt.Printf("%b\n", g)
	//print g in hex
	fmt.Printf("%#x\n", g)

}