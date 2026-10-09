package main

import "fmt"


func main(){
	// var name type
	var channelName string;
	channelName = "CodeWithHarry";

	var year int = 2026;

	fmt.Println("Channel: ",channelName, "Year: ", year);

	var rating float64 = 4.8;

	//there is a way to shorthand declare a var
	// := is used to declare and assign a value to a variable in one step
	subscribers := 5000
	likes, comments := 100,50
}