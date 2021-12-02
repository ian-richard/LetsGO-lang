package main

import "fmt"

func main () {
	defer runLast()
	runFirst()
}

func runLast(){
	fmt.Println("runLast")
}

func runFirst(){
	fmt.Println("runFirst")
}