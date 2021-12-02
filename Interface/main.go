package main

import ("fmt"
		"math")

func main () {
// fmt.Println(s.area())
// fmt.Println(c.area())

var s square = square{length: 15}
var c circle = circle{radius: 12.345}

info(s)
info(c)
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}





