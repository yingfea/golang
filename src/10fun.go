package main

import "fmt"

var mapFun = map[int]func(){
	1: func() {
		fmt.Println("Func1")
	},
	2: func() {
		fmt.Println("Func2")
	},
}

func add(num1, num2 int) int {
	return num1 + num2
}

func adds(sumList ...int) int {
	sum := 0
	for _, i := range sumList {
		sum += i
	}
	return sum
}

func fun1() (val string, ok bool) {
	return val, ok
}

func fun2() func() {
	return func() {
		fmt.Println("func里的func")
	}
}

func fun3(num *int) *int {
	*num += 1
	return num
}

func main() {
	num := add(1, 2)
	fmt.Println(num)
	nums := adds(1, 2, 3, 4, 5, 6, 7, 9)
	fmt.Println(nums)

	fun := fun2()
	fun()

	if fun, ok := mapFun[1]; ok {
		fun()
	}
	num3 := 10

	fun3(&num3)

}
