package main

import ("fmt"
		"math")

func main() {
	var res = combat(100, 20)
	var dead = combat(100, 200)
	fmt.Println(res == 80)
	fmt.Println(dead == 0)
}

// Description:

// Create a combat function that takes the player's current health and the amount of damage recieved, and returns the player's new health. Health can't be less than 0.

func combat_og(health, damage float64) float64 {
	var result = health - damage
	if result > 0 { return result}
	return 0;
  }

//refactor - math.Max() returns the larger of x or y.

func combat(health, damage float64) float64 {
  return math.Max(health-damage, 0.0)
}
