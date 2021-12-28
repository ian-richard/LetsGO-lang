package main

import (
	"fmt"
)

func main() {
	xs := []string{"James", "Moneypenny", "Q", "M"}
	m := xsToMap(xs)

	if ok := m["James"]; ok {
		fmt.Println("James exists")
	}

	if ok := m["Goldfinger"]; ok {
		fmt.Println("James exists")
	}

	//returns false:
	fmt.Println("Does Goldfinger exist in the map?", m["Goldfinger"])
}

func xsToMap(xs []string) map[string]bool {
	m := map[string]bool{}
	for _, v := range xs {
		m[v] = true
	}
	return m
}
