package main

import (
	"fmt"
	"math"
)

func main() {
	var res1 = TwiceAsOld(36, 7)
	var res2 = TwiceAsOld(55, 30)
	var res3 = TwiceAsOld(41, 21)
	fmt.Println(res1 == 22)
	fmt.Println(res2 == 5)
	fmt.Println(res3 == 0)
	
}

/*
Description:

Your function takes two arguments:

    current father's age (years)
    current age of his son (years)

Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old).
*/

func TwiceAsOld_og(dadYearsOld, sonYearsOld int) int { 
  
	var x = dadYearsOld - (sonYearsOld * 2);
	if x < 0 {
	  return x - x - x
	} else {
	  return x
	}
  }

  //refactor - math.Abs() returns the absolute positive value of x.

  func TwiceAsOld(dadYearsOld int, sonYearsOld int) int { 
	return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))
  }

