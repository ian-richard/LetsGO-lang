package main

import "fmt"

func main(){
	//icecream()
	//mapOfTypePerson()
	//printVehicle()
	anonymousStruct()
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

type vehicle struct{
	doors int
	color string
}

type truck struct{
	vehicle
	fourWheel bool
}

type sedan struct{
	vehicle
	luxury bool
}

/*Embedding is composition, not inheritance, but Go also does something called “promotion”, whereby the fields or methods of embedded types become available on the outer/embedding type. This provides a sort of automatic delegation and gives the pseudo-impression of “inheritance”*/

func printVehicle(){
	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(truck1.doors, truck1.color, truck1.fourWheel)
	fmt.Println(sedan1.doors, sedan1.color, sedan1.luxury)
}

//anonymous struct
func anonymousStruct(){
	s := struct{
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 121,
			"Q": 191,
			"M": 001,
		},
		favDrinks: []string{
			"Martini",
			"expresso martini",
			"sparking water",
		},
	}
	fmt.Println(s.first)
	for k, v := range s.friends{
		fmt.Println(k, v)
	}
	for i, v := range s.favDrinks{
		fmt.Println(i, v)
	}
}