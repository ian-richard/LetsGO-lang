package main

import (
	"fmt"
)


func main() {

	compositeliteral()
 
// var res = multipleOfIndex([]int{22, -6, 32, 82, 9, 25})
// fmt.Println(res) //=> [-6 32 25]
}

/* 
Return a new array consisting of elements which are multiple of their own index in input array (length > 1).

Some cases:

[22, -6, 32, 82, 9, 25] => [-6, 32, 25]

[68, -1, 1, -7, 10, 10] => [-1, 10]

[-56,-85,72,-26,-14,76,-27,72,35,-21,-67,87,0,21,59,27,-92,68] => [-85, 72, 0, 68]
*/

func multipleOfIndex (ints []int) []int {
 
	var result []int
	//for i in range(len(ints)):
	for i := 1; i < len(ints); i++ {
		//if ints[i] % i == 0:
		if ints[i]%i == 0 {
			//result.append(ints[i])
			result = append(result, ints[i])
		}
	}
	return result;
}
//create array which holds 5 values of type int and assign values to it
//range over the array an print the index and value of each element
func compositeliteral(){
	x := [5]int{21,22,23,24,25}
	for i, v := range x {
		fmt.Println(i, v)
	}
	//print the type of array
	fmt.Printf("%T\n", x)
}

