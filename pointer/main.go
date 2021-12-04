package main

import "fmt"

//create a person struct and func called ChangeMe as parameter
//to dereference a struct, use the (*value).field syntax
//reassign the value of the struct to the pointer. i.e "bob" to "alice"
//print out the pointer and the value of the struct

type person struct {
	name string
}

func main (){
	//ex. address in memory of the variable
	// v := 1
	// fmt.Println(&v)
	
	p1 := person {
		name: "Bob",
	}

	fmt.Println(&p1)
	fmt.Println(p1)
	changeName(&p1)
	fmt.Println(&p1)

}

//function which takes in a *person as a parameter
func changeName(p *person) {
	p.name = "Alice"
}

