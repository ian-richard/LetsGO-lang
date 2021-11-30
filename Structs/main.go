package main

import "fmt"

func main(){
	//icecream()
	mapOfTypePerson()
}

func icecream(){
	type person struct{
		first string
		last string
		icecream []string
	}
	p1 := person{
		first: "James",
		last: "Bond",
		icecream: []string{"chocolate", "vanilla", "strawberry"},
	}
	
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		icecream: []string{"coffee", "double choc", "toffee"},
	}
	fmt.Println(p1.first, p1.last)
	
	for i, v := range p1.icecream{
		fmt.Println(i+1, v)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.icecream{
		fmt.Println(i+1, v)
	}
}

func mapOfTypePerson(){
	type person struct{
		first string
		last string
		icecream []string
	}
	p1 := person{
		first: "James",
		last: "Bond",
		icecream: []string{"chocolate", "vanilla", "strawberry"},
	}
	
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		icecream: []string{"coffee", "double choc", "toffee"},
	}

	m := make(map[string]person)

	// m["Bond"] = p1
	// m["Moneypenny"] = p2

	m[p1.last] = p1
	m[p2.last] = p2

	for k, v := range m{
		fmt.Println(k, v.icecream)
	}

	fmt.Println(m)
}