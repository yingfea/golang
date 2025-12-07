package main

import (
	"fmt"
)

func main() {
	MySlice()
	MyMap()
}

func MySlice() {
	//make()
	s1 := make([]string, 5)  // 创建一个长度和容量都为 5 的 string 切片
	s2 := make([]int, 0, 10) // 创建一个长度为 0，容量为 10 的 int 切片
	//Literal
	nums := []int{1, 2, 3, 4, 5}
	nums = append(nums, 1, 2, 3)
	s1 = append(s1, "我", "是", "超", "人")
	s2 = append(s2, 1, 22, 3, 4, 5, 67, 7)
	PrintInt(s2)
	PrintInt(nums)
	PrintString(s1)

}

func MyMap() {
	nums := make(map[int]int)
	str := make(map[string]int)
	nums[1] = 1
	str["哈哈"] = 1

	strValue, ok := str["哈哈"]
	if ok {
		fmt.Printf("【哈哈】的值%v 存在\n", strValue)
	} else {
		fmt.Println("不存在")
	}

}

func PrintInt(array []int) {
	for _, value := range array {
		fmt.Println(value)
	}
}

func PrintString(array []string) {
	for _, value := range array {
		fmt.Println(value)
	}
}
