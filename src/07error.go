package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	//正常
	result1, err1 := ParseInt("123")
	if err1 != nil {
		fmt.Println("错误发生: ", err1)
		return
	}
	fmt.Println("转换成功: ", result1)
	//错误
	_, err2 := ParseInt("abc")
	if err2 != nil {
		fmt.Println("错误发生: ", err2)
		return
	}
}

type error interface {
	Error() string
}

func ParseInt(s string) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("输入字符不能为空")
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("转换失败：%w", err)
	}
	return val, nil
}
