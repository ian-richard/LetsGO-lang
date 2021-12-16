package main

// https://pkg.go.dev/encoding/json#Marshal
// https://pkg.go.dev/encoding/json#Unmarshal

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all the data", people)

	for i, v := range people {
		fmt.Println("\nPerson number:", i)
		fmt.Println("First Name:", v.First)
		fmt.Println("Last Name:", v.Last)
		fmt.Println("Age:", v.Age)
	}

}
