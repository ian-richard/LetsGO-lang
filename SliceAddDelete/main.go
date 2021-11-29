package main

import "fmt"

func main(){
	sliceOfStates()
}

var x = []int{42,43,44,45,46,47,48,49,50,51}

func sliceAdd(){
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	y := []int{56,57,58,59,60}
	x = append(x, y...)
	fmt.Println(x)
}

func sliceDelete(){
	y := x[:3]
	z := x[len(x)-4:]
	fmt.Println(y)
	fmt.Println(z)
	result := append(y, z...)
	fmt.Println(result)
	}

	//create slice to store all the names of the states in the united states of america
func sliceOfStates(){
	states := make([]string, 50, 50)
	states = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
	
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	// for i, v := range states{
	// 	fmt.Println(i+1, v)
	// }
	
	for i := 0; i < len(states); i++{
		fmt.Println(i+1, states[i])
	}
}