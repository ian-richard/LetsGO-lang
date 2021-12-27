package main

import ("fmt"
		"runtime")

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

//go build main.go
// to run, in the same directory, enter './main'