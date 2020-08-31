package type_test

import (
	"testing"
)

//别名
type MyInt int64

func TestImplicit(t *testing.T) {
	//不支持隐式类型转换
	var a int = 1
	var b int64
	//必须使用显式类型转换
	b = int64(a)
	var c MyInt
	//必须使用显式类型转换
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	//go语言不支持指针运算(如使用指针访问连续的存储空间) 在go里面是错误的
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	//获取变量类型
	t.Logf("%T %T", a, aPtr)
}
func TestString(t *testing.T) {
	//string在go里面是值类型 初始值是空串
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	t.Logf("%T", s)
	//	所以在字符串判空时直接这样
	if s == "" {
		t.Log("Hello World")
	}
}
