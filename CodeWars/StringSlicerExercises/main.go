package main

import "fmt"

// It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string. You're given one parameter, 
//the original string. You don't have to worry with strings with less than two characters.

func main() {

	fmt.Println(RemoveChar("hello")) //== ell
	fmt.Println(RemoveChar("TI"))    //== TI

}

func RemoveChar(word string) string {
	if len(word) > 2 {
		s := word[1:(len(word) - 1)]
		return s
	}
	return word
}
