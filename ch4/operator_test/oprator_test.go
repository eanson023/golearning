package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b, a == d)
	//长度不同比较编译错误
	//t.Log(a == c)
	//长度不同比较编译错误
	//t.Log(b == c)
	t.Log(c)
}
func TestAnweiZero(t *testing.T) {
	t.Log(1 &^ 0)
	t.Log(2 &^ 0)
	t.Log(2 &^ 1)
	//0011 1
	t.Log(3 &^ 1)
	//0011 0010
	t.Log(3 &^ 2)
	//0011 0011
	t.Log(3 &^ 3)
}
func TestConstantBit(t *testing.T) {
	a := 7 //0111
	t.Log(Readable)
	//去除可读权限
	a = a &^ Readable
	//0111 &^ 0001
	t.Log("a:", a)
	//0110 &^ 0100 去除可执行权限
	a = a &^ Executable
	t.Log("a:", a)
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
