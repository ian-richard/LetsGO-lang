package main

import (
	"fmt"
	"strings"
)

type MyString string
var s MyString = "Hello" 

func main() {
	var test = IsUpperCase(MyString("Hello")) //=> false
	var test2 = IsUpperCase(MyString("HELLO")) //=> true
	fmt.Println(test)
	fmt.Println(test2)
}

/*
Task: Is the string uppercase?

Create a method IsUpperCase to see whether the string is ALL CAPS. For example:

type MyString string
MyString("c").IsUpperCase() == false
MyString("C").IsUpperCase() == true
MyString("hello I AM DONALD").IsUpperCase() == false
MyString("HELLO I AM DONALD").IsUpperCase() == true
MyString("ACSKLDFJSgSKLDFJSKLDFJ").IsUpperCase() == false
MyString("ACSKLDFJSGSKLDFJSKLDFJ").IsUpperCase() == true

*/

func IsUpperCase(s MyString) bool {
  res := string(s)
  if res == strings.ToUpper(res){
    return true 
    } else {
    return false
  }
}
