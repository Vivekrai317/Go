package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName:= "Vivek"
	lastName:="Rai"

	fullName := firstName + " "+ lastName //concatenation

	fmt.Println("Fullname: ",fullName)

	fmt.Println(strings.ToUpper(fullName))
}