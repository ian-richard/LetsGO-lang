package main

import ("fmt")

func main() {
	//printnums(10)
	//ascii()
	printrune()
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