package main

// https://pkg.go.dev/golang.org/x/crypto/bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `anExamplePassword123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `anExamplePassword1232`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("Password does not match")
	} else {
		fmt.Println("Password matches")
	}

}
