package main

import ( "fmt"
        s "strings")

func main(){
	result := Points([]string{"1:0","2:0","3:0","4:0","2:1","1:3","1:4","2:3","2:4","3:4"})
	fmt.Println(result == 15)

	result1 := PointsRefactor([]string{"1:0","2:0","3:0","4:0","2:1","1:3","1:4","2:3","2:4","3:4"})
	fmt.Println(result1 == 15)

	result2 := PointsWithRange([]string{"1:0","2:0","3:0","4:0","2:1","1:3","1:4","2:3","2:4","3:4"})
	fmt.Println(result2 == 15)


}


func PointsWithRange(games []string) int {
	result := 0
	for _, game := range games {
	  str := s.Split(game, ":")
	  x, y := str[0], str[1]
	  switch {
		case x > y:
		  result += 3
		case x == y:
		  result += 1
	  }
	}
	return result
  }

func Points(games []string) int {
  result := 0
  i := 0
  for i < len(games) {
    event := games[i]
    stringsplit := s.Split(event, "")
    if stringsplit[0] == stringsplit[2] {
      result++
    } else if stringsplit[0]  > stringsplit[2] {
      result += 3
    }
    
    i++
  }
  return result;
}

func PointsRefactor(games []string) int {
	points := 0
	for _, g := range games {
		if g[0]>g[2] {
		  points += 3
		} else if g[0] == g[2] {
		  points += 1
		}
	}
	return points
  }