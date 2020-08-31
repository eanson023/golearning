package ch18

import "testing"

// 闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使己经离开了自由变量的环境也不会被释放或者删除，
// 在闭包中可以继续使用这个自由变量。

// 简单的说 : 函数+引用环境=闭包

// 贴士：闭包( Closure)在某些编程语言中也被称为 Lambda表达式（如Java）

// 闭包中可以修改引用的变量：
func TestChange(t *testing.T) {
	str := "Hello"
	foo := func() { //声明一个匿名函数
		str = "World"
	}

	foo() //调用匿名函数 修改str值
	t.Log(str)
}

//闭包案例一
func fn1(a int) func(i int) int {
	return func(i int) int {
		print(&a, "----", a)
		return a
	}
}

func TestClouse(t *testing.T) {
	f := fn1(1) //输出地址
	g := fn1(1) //输出地址
	t.Log(f(1))
	t.Log(f(1))
	t.Log(g(2))
	t.Log(g(2))
}

// 闭包案例二 实现累加器
func Accumulate(value int) func() int {
	return func() int { //返回一个闭包
		value++
		return value
	}
}

func TestCounter(t *testing.T) {
	accAdd := Accumulate(1)
	//2
	t.Log(accAdd())
	//3
	t.Log(accAdd())
}
