package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
	fmt.Println(TwoSort(s) == "b***i***t***c***o***i***n")
}

/*
You will be given a vector of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.

The returned value must be a string, and have "***" between each of its letters.

You should not remove or add elements from/to the array.
*/

  func TwoSort(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
  }