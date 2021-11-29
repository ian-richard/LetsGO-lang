package main

import ("fmt")

func main() {
	var test = demoSwitch("Bond")
	fmt.Println(test)
	var test2 = demoSwitch("notaname")
	fmt.Println(test2)
	var test3 = demoMulticaseSwitch("Dr No")
	fmt.Println(test3)
}

func demoSwitch(name string) string {
	switch name {
	case "MoneyPenny":
		return "Miss MoneyPenny"
	case "Bond":
		return "J.Bond"
	default:
		return "access denied"

	}
}

func demoMulticaseSwitch(name string) string {
	switch name {
	case "MoneyPenny", "Bond", "Dr No":
		return "hey gang"
	case "M":
		return "M"
	case "Q":
		return "Q"
	default:
		return "access denied"

	}
}


