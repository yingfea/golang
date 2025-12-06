package main

import "fmt"

func main() {
	const (
		NAME = "alex"
		AGE  = 11
	)
	fmt.Println(NAME, AGE)
	myfunc()
}

func myfunc() {
	const (
		NUM1 = iota
		NUM2
		NUM3
		NUM4
	)
	fmt.Println(NUM1, NUM2, NUM3)
}
