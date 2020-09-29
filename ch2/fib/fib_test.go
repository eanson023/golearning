package fib_test

import (
	"fmt"
	"testing"
)

//快速设置连续值
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

//Bit位赋值
const (
	Open = 1 << iota
	Close
	Pending
)

func TestFibTest(t *testing.T) {
	var a int = 1
	var b int = 1
	// 或者写为:
	// var (
	// 	c int = 1
	// 	d int = 1
	// )
	// // 使用类型推断赋值
	// e := 1
	fmt.Println(a, " ")
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		// tmp := a
		// a = b
		// b = tmp + a
		a, b = b, a+b
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	t.Log(Sunday, Open, Close, Pending)
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	// 另外一种交换赋值
	a, b = b, a
	t.Log(a, b)
}
