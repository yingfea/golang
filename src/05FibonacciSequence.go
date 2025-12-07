package main

import "fmt"

func fibonacci(n int, c chan int) {
	// f(n) = f(n-1)+ f(n+1)
	x, y := 0, 1 //初始值 从0，1开始

	for i := 0; i < n; i++ {

		c <- x // 把值附送给 Channel C

		x, y = y, y+x
	}

	// 从0，1开始
	// c = 0,
	// x=1,
	// y=0+1;

	// c=1,
	// x=1,
	// y=1+1,

	// c=1
	// x = 2,
	// y=2+1

	// c=2
	// x=3
	// y=3+2

	close(c) //关闭通道
}

func main() {
	n := 10 //生成容量为10的channel

	c := make(chan int, n) //创建一个用来传递值的管道 管道要传输的数据类型是int 容量为10

	//启动 Goroutine：在后台并发地生成斐波那契数列
	go fibonacci(n, c)

	fmt.Println("开始生成斐波那契数列：")
	for v := range c {
		fmt.Printf("%d ", v)
	}
	fmt.Println("\n所有数据接收完毕。")
}
