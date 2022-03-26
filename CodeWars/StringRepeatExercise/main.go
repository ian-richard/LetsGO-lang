package main

import "fmt"


func main(){

	fmt.Println(repeatStr(6, "I")) // "IIIIII"
	fmt.Println(repeatStr(5, "Hello")) // "HelloHelloHelloHelloHello"

}

// func RepeatStr(repititions int, value string) string {
//   return strings.Repeat(value, repititions)
// }

func repeatStr(reps int, value string) string {
	result := ""
	i := 1
	for i <= reps {
	  result += value
	  i++
	}
	
	return result
  }
