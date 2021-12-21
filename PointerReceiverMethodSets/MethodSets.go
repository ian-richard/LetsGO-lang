package main

/*
Hands-on exercise #2 - This exercise will reinforce our understanding of method sets:
● create a type person struct
● attach a method speak to type person using a pointer receiver
○ *person
● create a type human interface
○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
○ have it take in a human as a parameter
○ have it call the speak method
● show the following in your code
○ you CAN pass a value of type *person into saySomething
○ you CANNOT pass a value of type person into saySomething
● here is a hint if you need some help
○ https://play.golang.org/p/FAwcQbNtMG

Receivers 	|	Values
(t T) 		|	T and *T
(t *T)		|	*T

*/

import "fmt"

type person struct {
	name   string
	saying string
}

type human interface {
	speak() string
	saySomething() string
}

func (p *person) speak() string {
	return p.name
}

func (p *person) saySomething() string {
	return p.saying + " from " + p.name
}

func info(h human) {
	fmt.Println(h.speak())
}

func giveString(h human) {
	fmt.Println(h.saySomething())
}

func main() {
	p1 := person{name: "Pat", saying: "Hello"}
	//info(p1) error, because p1 is a value, not a pointer
	info(&p1)                      //ok, because p1 is a pointer
	giveString(&p1)                // ok, because p1 is a pointer
	fmt.Println(p1.saySomething()) //does work, despite not being a pointer

}
