package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(AbbrevName("sam harris")) //== S.H
}

func AbbrevName(name string) string {
	
	splitname := strings.Split(name, " ")
	return strings.ToUpper(string(splitname[0][0])) + "." + strings.ToUpper(string(splitname[1][0]))
	

}

