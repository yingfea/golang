package main

import "fmt"

func main(){
	var num=1; 
	var str1 = "hello world!";
	fmt.Println(num);
	fmt.Println(str1);
	fmt.Printf("he say: %v",str1);

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
	fmt.Printf("a=%v,b=%v, and alex say:%v",a,b,c)
}