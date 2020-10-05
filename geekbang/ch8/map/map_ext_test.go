package _map

import "testing"

func TestMapWithFunValue(t *testing.T) {
	//  value:方法(参数整型)返回值整型
	map1 := map[int]func(op int) int{}
	//一次方
	map1[1] = func(op int) int {
		return op
	}
	//平方
	map1[2] = func(op int) int {
		return op * op
	}
	//三次方
	map1[3] = func(op int) int {
		return op * op * op
	}
	t.Log(map1)
	t.Log("--------------")
	t.Log(map1[1](2), map1[2](2), map1[2](2))
	for _, v := range map1 {
		t.Log(v(2))
	}
	t.Log("--------------")
	if v, ok := map1[0]; ok {
		t.Log(v(2))
	}
	t.Log("---------------")
	fun1 := map1[1]
	t.Log(fun1(1))
}
