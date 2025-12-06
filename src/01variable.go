package main

import "fmt"

// global variable
var (
	User   string
	Age    int
	Height float32
)

func main() {
	var num = 1
	var str1 = "hello world!"
	fmt.Println(num)
	fmt.Println(str1)
	fmt.Printf("he say: %v", str1)

	// local variable
	a := 10
	b := 20
	c := "hello! I am alex!"
	/*
		%d int
		%s string
		%v any type
		%T tepy of
		%f float
		...
	*/
	fmt.Printf("a=%v,b=%v, and alex say:%v \n", a, b, c)

	var name string = "alex"
	var age int = 20
	var height float32 = 1.61
	fmt.Printf("name=%v, age=%v, and how tall:%v \n", name, age, height)

	User = "Hello"
	fmt.Println(User)

}
