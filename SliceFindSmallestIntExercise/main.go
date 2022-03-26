package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100})) //== -345
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2}))     //== 2

	fmt.Println(SmallestIntegerFinderV2([]int{34, -345, -1, 100})) //== -345
	fmt.Println(SmallestIntegerFinderV2([]int{34, 15, 88, 2}))     //== 2

	fmt.Println(SmallestIntegerFinderSort([]int{34, -345, -1, 100})) //== -345
	fmt.Println(SmallestIntegerFinderSort([]int{34, 15, 88, 2}))     //== 2

}

func SmallestIntegerFinder(numbers []int) int {
	result := numbers[0]

	for _, i := range numbers {
		if i < result {
			result = i
		}
	}
	return result
}

//refactor

func SmallestIntegerFinderV2(numbers []int) int {
	smallest := numbers[0]

	for i := range numbers {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	return smallest
}

//sort

func SmallestIntegerFinderSort(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
