package ch20

import (
	"fmt"
	"testing"
)

//空接口与断言

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println(i)
	// }

	//type关键字
	switch v := p.(type) {
	case int:
		fmt.Println("int型 ", v)
	case string:
		fmt.Println("string型 ", v)
	case byte:
		fmt.Println("byte型 ", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func TestAssert(t *testing.T) {
	DoSomething(111)
	DoSomething('a')
	DoSomething("哈哈哈")
	var c byte = 'c'
	DoSomething(c)
}
