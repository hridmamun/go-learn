package main

import "fmt"

func add(num1 int,num2 int) int {
	sum:=num1+num2
	return sum
}
/*
const a = 10
var p = 100

func call() {
	add:= func(x int,y int) {
		z:=x+y
		fmt.Println(z)

	}
	add(5,6)
	add(p,a)
}

func main() {
	call()
	fmt.Println(a)

}

func init() {
	fmt.Println("Hello World")
}
*/

/*
type User struct {
	Name string
	Age int
}

func newUser(name string,age int) User {
	return User {Name:name,Age:age}
}

func main() {
	user:= newUser("Abdullah,25")
	fmt.Println(user.Name,user.Age)
}
*/

/*
func rectangleStats(length,width float64) (area, perimeter float64) {
	area=length*width
	perimeter=2*(length+width)
	return
}

func main() {
	a,p:= rectangleStats(10,5)
	fmt.Println("Area:",a,"perimeter:",p)
}
*/

/*
func main() {
	x:=10

	if x>5 {
		fmt.Println("x is greater than 5")
	}
}
*/
/*
func main() {
	age:=16

	if age>=18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}
}
*/
/*
func main() {
	score:=75

	if score>=90 {
		fmt.Println("Grade A")
	} else if score>=80 {
		fmt.Println("Grade B")
	} else if score >=70 {
		fmt.Println("Grade C")
	}else {
		fmt.Println("Grade F")
	}
	}
	*/
/*
	func main(){
		if num:=8; num%2==0 {
			fmt.Println(num,"is even")
		} else {
			fmt.Println(num,"is odd")
		}
	}
	*/
/*
	func main() {
		age:=25
		citizen:=true

		if age >=18 {
			if citizen {
				fmt.Println("You are eligible to vote.")
			} else {
				fmt.Println("You are not a citizen.")
			}
			} else {
				fmt.Println("You are too young to vote.")
			
			}
		}
	
*/
/*
func main() {
	var name string="Abdullah"
	var age int=25
	height:=5.3

	fmt.Println("Name:",name)
	fmt.Println("Age:",age)
	fmt.Println("Height:",height)
}
*/
/*
func main() {
	for i :=1; i<=5; i++ {
		fmt.Println("Count:",i)
	}
}
*/
/*
func add(a int, b int) int {
	return a+b
}

func main() {
	result:=add(10,5)
	fmt.Println("Result:",result)
}
*/

/*
func main() {
	fruits:=[] string{"Apple","Banana","Cherry"}

	for index,value:= range fruits {
		fmt.Println("index:", index, "fruit:",value)
	}
}
*/
/*
func main() {
	var name string = "Abdullah"
	var age int = 25
	const country = "Bangladesh"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country", country)
}
*/
/*
func main() {
	number:=10

	if number>0 {
		fmt.Println(number,"is a positive number")
	} else if number<0 {
		fmt.Println(number,"is a negative number")
	} else {
		fmt.Println(number,"is zero")
	}
	}
	*/

/*
	func main() {
		day:=3

		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		default:
			fmt.Println("Other day")
		}
	}
*/

/*
func main() {
	person := map [string] string {
		"name": "Abdullah",
		"city": "Jashore",
	}

	fmt.Println("Name:", person["name"])
	fmt.Println("City:", person["city"])
}
*/
/*
func divide(a,b int) (int,int) {
	quotient:= a/b
	reminder:= a % b
	return quotient,reminder
}

func main() {
	q,r := divide(10,3)
	fmt.Println("Quotient:",q)
	fmt.Println("Reminder:",r)
}
*/
/*
func main() {
	var name string="Golang"
	const version = "1.22"

	fmt.Println("Language:", name)
	fmt.Println("Version:", version)
}
	*/

/*
	func greet() {
		fmt.Println("Hello,from a function!")
	}

	func main() {
		greet()
	}
		*/
/*
		func sayHello(name string) {
			fmt.Println(Hello, name)
		}
		func main() {
			sayHello("Hridoy")
		}
			*/

/*
			func operate(a,b int, operation func(int,int) int) int {
				return operation(a,b)
			}

			func main() {
				add:= func(x,y int) int {return x+y}
					mul:= func(x,y int) int {return x*y}


						fmt.Println("Addition:", operate(4,2,add))
						fmt.Println("Multiplication:",operate(4,2,mul))
					}
				*/


				func main() {
					a:=10
					b:=20
					sum:=add(a,b)
					fmt.Println("Sum:",sum)
				}
	