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

	fmt.Println(m) //=> whole map
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
}