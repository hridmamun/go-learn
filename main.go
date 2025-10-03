package main

import "fmt"

func main() {
	age	:= 20
	if age>18 {
		fmt.Println("You are eligable to married")
	} else if age<18{
		fmt.Println("You are not eligable to married")
	}
fmt.Println("final result is: ", result)
fmt.Println("final result of subtraction is : ", sub(20,10))
}


func sum(x int, y int ) int {
	var sum = x + y;
	return sum;
}


func sub(x int, y int ) int {
	var sub = x - y;
	return sub;
}
