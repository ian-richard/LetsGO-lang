package main

import ("fmt")

func main() {
	//printnums(10)
	//ascii()
	//printrune()
	//printyears(1901, 2001)
	//printmodulus()
	loop := recursionWithLoop(4)
	fmt.Println((loop) == 24)
}

//func which prints every number between 1 and 10
func printnums(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}
// prints every character in as num, hexidecimal, unicode + character 
func ascii(){
	for i := 33; i <= 122; i++ {
			fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
			//num, hexidecimal, unicode + character 
		}
	}

// print the rune value of the character once (or three times using j-loop)
func printrune(){
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		//for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		//}
	}
}

//print numbers from years alive to current year, take age and years as args
func printyears(birthyear int, currentyear int){
	//for i := birthyear; i <= currentyear; i++ {
	for birthyear <= currentyear {
		fmt.Println(birthyear)
		birthyear++
	}
}

//loop to print the modulus for every number between 10 and 100
func printmodulus(){
	for i := 10; i <= 100; i++ {
		fmt.Printf(" when %v is divided by 4, the remainder is% v\n", i, i%4)
	}
}

//loop with same affect as recursion
func recursionWithLoop(n int) int{
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}