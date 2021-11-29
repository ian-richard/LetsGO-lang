package main

import "fmt"

func main(){
	sliceAdd()
}

var x = []int{42,43,44,45,46,47,48,49,50,51}

func sliceAdd(){
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	y := []int{56,57,58,59,60}
	x = append(x, y...)
	fmt.Println(x)
}