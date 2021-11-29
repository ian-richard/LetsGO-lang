package main

import "fmt"

func main() {
	mapPeople();
}

func mapPeople() {
	m := map[string]int{
		"James Bond": 40,
		"Miss Moneypenney": 27,
	}
	//add a key value pair
	m["Austin Powers"] = 48

	fmt.Println(m) //=> whole map

	//loop over map
	for k, v := range m {
		fmt.Println(k, v)
	}
	
	fmt.Println(m["James Bond"]) //=> value of key

	//check if value exists idiom
	if v, ok := m["James Bond"]; ok {
		fmt.Println("Value:", v)
	} else {
		fmt.Println("Key not found")
	}
	//not found
	if v, ok := m["Dr No"]; ok {
		fmt.Println("Value:", v)
	} else {
		fmt.Println("Key not found")
	}

	//loop over a composite literal
	fmt.Println("\nLoop over a composite literal")
	slice := []string{"James Bond", "Miss Moneypenney", "Austin Powers"}
	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Println("\nLoop over a composite literal__without index printed")
	slice_2 := []string{"James Bond", "Miss Moneypenney", "Austin Powers"}
	for _, v := range slice_2 {
		fmt.Println(v)
	}
}

//see "comma, ok" idiom here https://go.dev/doc/effective_go.html