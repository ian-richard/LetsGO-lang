package main

import "fmt"

func main() {
	mapDict()
	//mapPeople();
}

func mapDict(){
	//map with key of type string and value of type string
	m := map[string][]string{
		"bond_james": []string{"martini, shaken, not stirred", "swish cars", "misogyny"},
		"moneypenny_miss": []string{"JB", "Books", "Computer Science"},
	}

	m["no_dr"] = []string{"evil", "more evil", "mean cats"}


	
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	for k, val := range m {
		fmt.Println("This record is for: ", k)
			for i, values := range val {
				fmt.Println("\t", i+1, values)
			}
	}
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
	fmt.Println("Look up Dr No: ")
	if v, ok := m["Dr No"]; ok {
		fmt.Println("Found: ", v)
	} else {
		fmt.Println("Key not found")
	}

	fmt.Println("\nHow to delete (Austin Powers) a key value pair from a map")
	if _, ok := m["Austin Powers"]; ok {
		//fmt.Println("Value:", v)
		fmt.Println("Deleted:", ok)
		delete(m, "Austin Powers")
	} else {
		fmt.Println("Key not found")
	}
	fmt.Println(m) //==> Mr Powers is no more 

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