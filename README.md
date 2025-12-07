# 突然有点迷茫，似乎学什么都没有出路 悲！`

# 笨蛋在学习Golang

## Hello World!

```go
package main  //这是特殊的包名，定义了一个独立可执行的程序。所有的可执行程序都必须有一个 `main` 包

import "fmt"  //用于导入其他包， `fmt` 包实现了格式化 I/O

//这是程序的​入口函数​。当程序运行时，它会从 `main` 包中的 `main` 函数开始执行
func main() {
	//使用打印方法 打印一行 hello world
	fmt.Println("hello world!")
}
```

## 变量（Variables）和 常量（Constants）

##### 变量（Variables）

###### 类型后置

生命变量类型放变量名后 哥们有点反直觉

```c#
// csharp
string name = "哈哈"
```

```go
//Golang
var name string ="哈哈"
```

###### 完整声明

```go
var name string = "哈哈"
```

###### 类型推断: 如果初始化了变量，Go 可以自动推断出变量的类型**

```go
var name = "哈哈" 
var pi = 3.141596
```

###### 短变量声明: 但是只能在函数中使用

```go
func main(){
	name := "哈哈"
}
```

##### 常量（Constants）

###### const  一旦定义就不能被改变

```go
const Name = "哈哈"
```

