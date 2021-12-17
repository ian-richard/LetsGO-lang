package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

type jsonToGoUser struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Miss",
		Age:   27,
	}

	users := []user{u1, u2}
	fmt.Println("users in GO: ", users)

	//Marshal data into JSON
	output, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON: ", string(output))

	//Unmarshal JSON into data
	var jsonToGoUsers []jsonToGoUser
	err = json.Unmarshal(output, &jsonToGoUsers)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON to GO: ", jsonToGoUsers)

	//loop over
	fmt.Println("\nLoop over JSON to GO: ")
	for i, v := range jsonToGoUsers {
		fmt.Println("\nPerson number:", i)
		fmt.Println("First Name:", v.First)
		fmt.Println("Age:", v.Age)
	}

}
