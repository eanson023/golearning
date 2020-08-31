package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	//	while循环
	//	while (n<5)
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
	//	while死循环
	for {
		t.Log("xixix")
	}
}
func TestIfMultiSec(t *testing.T) {
	// 声明 再用其作为判读
	//通常用于 方法多返回值时
	if v, err := someReturnFun(); err != nil {
		t.Log(err)
	} else {
		t.Log(v)
	}
}

func someReturnFun() (interface{}, interface{}) {
	return 1, 2
}
