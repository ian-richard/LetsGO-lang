package main

import ("fmt")

func main() {
	f := enclosedNum()
	fmt.Println(f())
	f()
	f()
	f()
	fmt.Println(f()) //=> 5
	fmt.Println("now to the next example")
	g := enclosedNum()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}

func enclosedNum() func() int {
	var num int
	return func() int{
		num++
		return num
	}
}
