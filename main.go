package main

import "fmt"

func main() {
	/*
	int
	float32
	bool
string
*/
var x=10
fmt.Println(x)

result := sum(10,20)

fmt.Println("final result is: ", result)
}


func sum(x int, y int ) int {
	var sum = x + y;
	return sum;
}
