package main

import (
	"fmt"
)

func main() {
	MyIf(70)
	MyFor()
	MyWhile(1)
	names := []string{"张三", "alex", "哈哈"}
	MyRange(names)
	MySwitch("B")
}

func MyIf(score int) {
	if rank := score; rank > 90 {
		fmt.Println("优秀！")
	} else if rank > 80 {
		fmt.Println("良好！")
	} else {
		fmt.Println("仍需努力")
	}
}

func MyFor() {
	for i := 0; i < 6; i++ {
		fmt.Println("计数: ", i)
	}
}

func MyWhile(sum int) {
	for sum < 5 {
		fmt.Println("sum =", sum)
		sum++
	}
}

func MyRange(nameArray []string) {
	for index, value := range nameArray {
		fmt.Printf("index: %v, name: %v \n", index, value)
	}
}

func MySwitch(str string) {
	switch str {
	case "A":
		fmt.Println("Lv A!")
	case "B", "C":
		fmt.Println("Lv B!")
	case "D", "F":
		fmt.Println("Lv C")
	default:
		fmt.Println("Go Out")
	}
}
