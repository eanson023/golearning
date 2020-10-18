package _func

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

/*
	多返回值
*/
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFunc1(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	//忽略一个值
	a, _ = returnMultiValues()
	t.Log(a)
}

/*
输入是函数类型 返回也是函数类型
*/
func timeSpend(inner func(op int) int) func(op int) int {
	//类似面向对象的装饰者模式
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Milliseconds())
		return ret
	}
}

func slowFunc(second int) int {
	time.Sleep(time.Second * time.Duration(second))
	return second
}
func TestFn2(t *testing.T) {
	//返回一个函数
	tsSF := timeSpend(slowFunc)
	t.Log(tsSF(5))
}

/*
可变长参数
*/
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5, 6))
	t.Log(Sum(1, 2, 3, 4, 5, 6, 8))
}

func Clear() {
	fmt.Println("clear resources1.")
}

func TestDefer(t *testing.T) {
	//方法执行完再执行defer 后执行
	defer func() {
		t.Log("clear resources2.")
	}()
	//先执行
	defer Clear()
	t.Log("started.")
	//抛出不可修复的错误仍然可以会运行defer函数
	panic("error")
	t.Log("end.")
}

// 使用函数作为参数
func Map(f func(rune) rune, s string) string {
	rs := make([]rune, 0)
	for _, r := range s {
		rs = append(rs, f(r))
	}
	return string(rs)
}

func add1(r rune) rune {
	return r + 1
}

//
func TestMap(t *testing.T) {
	s := "hello world"
	res := Map(add1, s)
	t.Log(res)
	// 匿名函数
	res = strings.Map(func(r rune) rune { return r - 1 }, res)
	t.Log(res)
}
