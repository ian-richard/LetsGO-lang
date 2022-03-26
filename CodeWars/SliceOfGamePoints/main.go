package main

import ( "fmt"
        s "strings")

func main(){
	result := Points([]string{"1:0","2:0","3:0","4:0","2:1","1:3","1:4","2:3","2:4","3:4"})
	fmt.Println(result == 15)
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