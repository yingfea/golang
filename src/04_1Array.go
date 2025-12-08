package main

import (
	"fmt"
	"sort"
)

func main() {
	nameList := make([]string, 0)
	nameList = append(nameList, "Alex")
	nameList = append(nameList, "Fou")
	nameList = append(nameList, "Du")
	fmt.Println(nameList[0])

	var numsList = []int{1, 3, 2}
	sort.Ints(numsList)
	fmt.Println(numsList)
	sort.Sort(sort.Reverse(sort.IntSlice(numsList)))
	fmt.Println(numsList)
}
