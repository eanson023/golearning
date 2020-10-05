package main

import (
	"fmt"
)

// 在C++中，开发者需要自己手动分配内存来适应不同的算法需求。比如，函数局部变量尽量使用栈（函数退出，内部变量也依次退出），全局变量、结构体使用堆。
// Go语言将这个过程整合到了编译器中，命名为“变量逃逸分析”，这个技术由编译器分析代码的特征和代码生命期，决定是堆还是栈进行内存分配。

// # -gcflags参数是编译参数，-m表示进行内存分析，-l表示避免程序内联（优化）
// go run -gcflags "-m -l" main.go
func test(num int) int {
	var t int
	t = num
	return t
}

//空函数，什么也不做
func void() {

}

func main() {

	var a int               //声明变量并打印
	void()                  //调用空函数
	fmt.Println(a, test(0)) //打印a，并调用test

}
