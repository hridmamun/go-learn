package main

import "fmt"

/*
func add(num1 int,num2 int) int {
	sum:=num1+num2
	return sum
}
	*/
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

/*
				func main() {
					a:=10
					b:=20
					sum:=add(a,b)
					fmt.Println("Sum:",sum)
				}
	*/

	/*
	func getNumbers(num1 int,num2 int) (int,int) {
		sum:=num1+num2
		mul:=num1*num2
		return sum,mul
	}

	func main() {
		a:=10
		b:=20
		p,q:=getNumbers(a,b)
		fmt.Println("Sum:",p)
		fmt.Println("Multiplication:",q)
	}
		*/

/*
		func printWelcomeMessage() {
			fmt.Println("Welcome to the application")
		}

		func getUserName() string {
			var name string
			fmt.Println("Enter your name- ")
			fmt.Scanln(&name)
			return name
		}

		func getTwoNumbers() (int,int) {
			var num1 int
			var num2 int
			fmt.Println("Enter first number- ")
			fmt.Scanln(&num1)
			fmt.Println("Enter second number- ")
			fmt.Scanln(&num2)
			return num1,num2
		}

		func add(num1 int,num2 int) int {
			sum:=num1+num2
			return sum
		}

		func display(name string,sum int) {
			fmt.Println("Hello",name)
			fmt.Println("The sum is:",sum)
		}
		func printGoodbyeMessage() {
			fmt.Println("Thank you for using the application.Goodbye!")
		}
		func main() {
			printWelcomeMessage()
			name:=getUserName()
			num1,num2:=getTwoNumbers()
			sum:=add(num1,num2)
			display(name,sum)
			printGoodbyeMessage()
		}
*/
/*
var (
	a=20
	b=30
)
func add(x int,y int) {
	z:=x+y
	fmt.Println(z)

}
func main () {
	p:=30
	q:=40
	add(a,b)
	add(p,q)
	add(a,p)
}
*/
/*
var totalStudents int


type student struct {
	ID int
	Name string
	Age int
	Grades [] float64
}

func(s student) AverageGrade() float64 {
	sum:=0.0
	for _ , g:= range s.Grades {
		sum+=g
	}
return sum /float64(len(s.Grades))
}
func NewStudent(id int,name string,age int,grades [] float64) student {
	totalStudents++
	return student{ID: id,Name: name,Age: age,Grades: grades}
}

func PrintStudents(students [] student) {
	fmt.Println("== student List ==")
	for _,s:= range students {
		fmt.Println("ID:%d|Name:%s|Age:%d |Average:%.2f\n",s.ID,s.Name,s.Age,s.AverageGrade() )
	}
	fmt.Println("=======")
}

func FindStudent(students [] student, id int) *student {
	for _ , s:= range students {
		if s.ID==id {
			return &students[i]
		}
	}
	return nil
}
func main() {
	students:= [] student{
		Newstudent(1,"Alice",20,[] float64{90,85,88}),
		Newstudent(2,"Bob",22,[] float64{75,80,72}),
	Newstudent(3,"Charlie",19,[] float64{95,92,98}),
}
Printstudents(students)
fmt.Println("\nSearching for student ID 2..")
found:= FindStudent(students,2)
if found != nil {
	fmt.Printf("Found: %s,Average:%.2f\n", found.Name,found.AverageGrade())
} else {
	fmt.Println("student not found")
}
totalStudent:=100
fmt.Println("\nLocal totalStudents:",totalStudents)
fmt.Println("Global totalStudents:",getglobalTotal())

newStudent:= NewStudent(4,"Daina",21,[] float64{82,79,85})
students= append(students,newStudent)
fmt.Println("\nAfter adding a new student:")
PrintStudents(students)
}
func getGlobalTotal() int {
	return totalStudents
}
*/
/*
var (
	a=10
	b=20
)

func printNum(num int) {
	fmt.Println(num)
}
func add(x int,y int) {
	res:=x+y
	printNum(res)
}

func main() {
	add(a,b)
}
	*/
	/*
	var name = "Abdullah"
	var age = 25

	func main() {
		greet()
	}
	func greet() {
		fmt.Println("Hello,", name)
		fmt.Println("Age:", age)
	}
		*/

   /*
		const pi=3.1416

		func main() {
			area:= pi*5*5
			fmt.Println("Area of circle:",area)
		}
			*/

/*
			var a=10
			func main() {
				age:=30
				if age>18 {
					a:=47
					fmt.Println(a)
				}
				fmt.Println(a)
			}
				*/

			/*	
				func main() {

					func(a int,b int) {
						c:=a+b
						fmt.Println(c)
				} (5,6)
					}

					func init() {
						fmt.Println("I'll be call first")
					}
			*/
/*
			func main() {
				add:= func(a int,b int) {
					c:=a+b
					fmt.Println(c)
				}
				add(5,6)
			}

			func init() {
				fmt.Println("I'll be call first")
			}
				*/

/*
				func main() {
					multiply:= func(a int,b int) int {
						return a*b
					}
					fmt.Println("product:",multiply(4,5))
				}
				*/
/*
				func operate(a int,b int, f func(int,int)int) int{
					return f(a,b)
				}

				func main() {
					sum:= func(x int,y int) int {return x+y}
					diff:=func(x int,y int) int {return x-y}

					fmt.Println("Sum:",operate(10,5,sum))
					fmt.Println("Difference:",operate(10,5,diff))
				}
					*/

		
					/*
					func add(a int,b int) {
						c:=a+b
						fmt.Println(c)
					}
					func main() {
						add(2,5)
						sum:=call()
						sum(4,7)
					}
					func processOperation(a int,b int,op func(p int,q int)) {
						op(a,b)
					}
					func call() func(x int,y int) {
						return add
					}
						*/

/*
						var a = 10
						func add(x int,y int) {
							z:=x+y
							fmt.Println(z)
						}

						func main() {
							add(5,4)
							add(a,3)
						}

						func init() {
							fmt.Println("Hello")
						}
							*/
                     /*
							const a = 10
							var p = 100
							func call() {
								add:=func(x int,y int) {
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
								fmt.Println("Hello")
							}
							*/
/*
							const a =10
							var p = 100
							func outer(money int) func() {
								age:=25
							fmt.Println("Age=",age)
							show:= func() {
								money=money+a+p
								fmt.Println(money)
							}
                           return show
							}

							func call() {
								incr1:=outer(100)
								incr1()
								incr1()
								incr2:=outer(100)
								incr2()
								incr2()
								
							}
							func main() {
								call()

							}
							func init() {
								fmt.Println("===Bank===")
							}
								*/

/*
								type User struct {
									Name string
									Age int
								}
								func (usr User) printDetails() {
									fmt.Println("Name:",usr.Name)
									fmt.Println("Age:",usr.Age)
								}
								func main() {
									user1:= User{
										Name: "Tuhin",
										Age: 30,
									}
									user2:=User {
										Name: "Hridoy",
										Age: 25,
									}
									user1.printDetails()
									user2.printDetails()
								}
									*/
                                /*
									type Rectangle struct {
										Width float64
										Height float64
									}

									func (r Rectangle) Area() float64 {
										return r.Width * r.Height
									}
									func main() {
										rect:=Rectangle{Width:10,Height:5}
										fmt.Println("Area:",rect.Area())
									}*/
                                   /*
									type Counter struct {
										count int
									}
									func (c* Counter) Increment() {
										c.count++
									}
									func main() {
										c:= &Counter{}
											c.Increment()
											c.Increment()
											fmt.Println("Count:",c.count)
										
									}
											*/
/*
						type Address struct {
							City,Country string
						}
						type Employee struct {
							Name string
							Age int
							Address Address
						}
						func main() {
							e:= Employee{
								Name: "Abdullah",
								Age: 25,
								Address: Address {
									City: "Jashore",
									Country: "Bangladesh",
								},
			
							}
							fmt.Println(e)
							fmt.Println("Employee City:",e.Address.City)
						}
*/
/*
func main() {
	user:= struct {
		ID int
		Name string
		Role string
	}{
		ID: 1,
		Name: "Abdullah",
		Role: "Admin",
	}
	fmt.Println(user)
}
	*/
/*
	type Student struct {
		Name string
		Marks  int
	}
	func main() {
		students:= [] Student{
			{"Rafi",85},
			{"Tuhin", 90},
			{"Hridoy",78},
		}
		for _,s:=range students {
			fmt.Println("%s scored %d\n",s.Name,s.Marks)
		}
	}
		*/
/*
		type Animal struct {
			Name string
		}
		func (a Animal) Speak() {
			fmt.Println(a.Name,"makes a sound")

		}
		type Dog struct {
			Animal
			Breed string
		}
		func main() {
			d:= Dog{Animal{"Buddy"},"Labrador"}
			d.Speak()
			fmt.Println("Dog Breed:",d.Breed)
		}
			*/
/*
			func main() {
				var numbers[5] int
				numbers[0]=10
				numbers[1]=20
				numbers[2]=30
				numbers[3]=40
				numbers[4]=50
				fmt.Println("Numbers:",numbers)
			}
				*/
/*
				func main() {
					names:= [3] string{"Abdullah","Tuhin","Hridoy"}
					fmt.Println("Names:",names)
				}*/

/*
				func main() {
					arr:=[5] int{2,4,6,9,8}
					for i:=0; i<len(arr) ;i++ {
						fmt.Println("Element at index %d=%d\n",i,arr[i])
					}
				}
					*/
/*
					func main() {
						fruits:= [3] string{"Apple","Banana","Cherry"}
						for index, value:= range fruits {
							fmt.Println("Index:",index,"Fruits:",value)
						}
					}
						*/
/*
						func main() {
							temperatures:= [4] float64{32.5,33.5,29.8,30.0}
							var sum float64
							for _,value:= range temperatures {
								sum+=value
							}
							average := sum/ float64(len(temperatures))
							fmt.Println("Average Temperature:",average)
						}
							*/
/*
							func main() {
								a := [3] int{10,20,30}
								b := [3] int{10,20,30}
								c := [3] int{5,15,25}
								fmt.Println("a==b:",a==b)
								fmt.Println("a==c:",a==c)
							}
								*/
/*
								func main() {
									var nums [3] int
									nums[0]=10
									nums[1]=20
									nums[2]=30
									fmt.Println("Array:",nums)
									fmt.Println("First element:",nums[0])
									fmt.Println("Lenght:",len(nums))
								}
									*/
/*
									func main() {
										nums:= [4] int{1,2,3,4}
										fmt.Println(nums)
										letters:= [...] string {"a","b","c"}
										fmt.Println(letters)
									}
										*/
/*
										func main() {
											arr:= [3] string{"Go","Rust","Python"}
											for i:=0; i< len(arr) ;i++ {
												fmt.Println(i,arr[i])
											}	
											fmt.Println("Using range:")
											for index,value:= range arr {
												fmt.Println(index,value)
											}								
											}
											*/
											/*
											func main() {
												var matrix [2] [3] int = [2] [3] int {
													{1,2,3},
													{4,5,6},
												}
												fmt.Println("Matrix:",matrix)
												fmt.Println("Element at (1,2):",matrix [1] [2]) 
											}
												*/
/*
												type Employee struct {
													ID int
													Name string
													Department string
													Salary float64
												}
												func main() {
													employees:= [] Employee{
														{ID:1,Name:"Alice",Department:"IT",Salary:75000.00},
														{ID:2,Name:"Bob",Department:"Finance",Salary:65000.00},
														{ID:3,Name:"Charlie",Department:"HR",Salary:60000.00},
													}
													fmt.Println("All Employees:")
													for _,emp := range employees {
														fmt.Printf("%d: %s works in %s,earns $%,2f\n",emp.ID,emp.Name,emp.Department,emp.Salary)
													}
												}
													*/
/*
													type product struct {
														Code string
														Name string
														Quantity int
														Price float64
													}
													func main() {
														inventory := [] product {
															{"P1001","Laptop",5,999.99},
															{"P1002","Mouse",25,19.99},
															{"P1003","Keyboard",10,49.99},
														}
														totalValue :=0.0
														for _,item:= range inventory {
															totalValue += float64 (item.Quantity)*item.Price
														}
														fmt.Printf("Total Inventory Value:$%,2f\n",totalValue)
													}
														*/


														type Student struct {
															ID int
															Name string
															Scores [3] float64
														}
														func main() {
															Students := [] Student {
																{1,"Tuhin", [3] float64{80,90,85}},
																{2,"Hridoy",[3] float64{88,92,95}},

															}

															for _, s:= range Students {
																avg := (s.Scores[0]+s.Scores[1]+s.Scores[2]/3)
																fmt.Printf("%s's average score:% 2f\n",s.Name,avg)
															}
														}