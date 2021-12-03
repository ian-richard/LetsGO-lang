package main

import "fmt"

func main(){
	// fmt.Println(sumSlice([]int{1,2,3,4}...))
	// fmt.Println(sumSlice(1,2,3))

	// fmt.Println(sumSlice__notVariadic([]int{1,2,3,4}))
	//fmt.Println(p.speak())
	fmt.Println(x(2))
	}

func sumSlice(numbs ...int) int {
	result := 0
	for _, num := range numbs {
		result += num
	}
	return result
}

func sumSlice__notVariadic(numbs []int) int {
	result := 0
	for _, num := range numbs {
		result += num
	}
	return result
}

type Person struct {
	firstname string
	lastname string
	age int
}

var p = Person{
	firstname: "James",
	lastname: "Bond",
	age: 45,
}

func (p Person) speak() string {
	return fmt.Sprintf("I'm %s %s and I'm %d years old", p.firstname, p.lastname, p.age)
}

func (p Person) getNameAndAge() string {
	return fmt.Sprintf("%s %s is %d years old", p.firstname, p.lastname, p.age)
}


var x = func (num int) int {
	return num + num
	}

func bar (num int, s string) (int, string) {
	return num, s
	}


