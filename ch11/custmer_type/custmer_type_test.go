package custmer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpend(inner IntConv) IntConv {
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
