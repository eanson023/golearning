package ch13

import (
	"math"
	"testing"
)

// 常量：在编译阶段就确定下来的值，程序运行时无法改变。
const A = 3
const PI float32 = 3.1415
const mask = 1 << 3 //常量与表达式

// 错误写法：常量赋值是一个编译期行为，右边的值不能出现在运行时才能得到结果的值。
// const HOME = os.Getenv("HOME")

// 无类型常量
// 一个常量可以有任意一个确定的基础类型，例如int或float64，但是许多常量并没有一个明确的基础类型。

// 无类型常量的作用：

// - 编译器会为这些没有明确基础类型的数字常量提供比基础类型更高精度的算术运算，可以认为至少有256bit的运算精度
// - 无类型的常量可以直接用于更多的表达式而不需要显式的类型转换

func TestConst1(t *testing.T) {
	// 示例：math.Pi无类型的浮点数常量，可以直接用于任意需要浮点数或复数的地方：
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	t.Log(x, y, z)
}

// 如果math.Pi被确定为特定类型，比如float64，
// 那么结果精度可能会不一样，同时对于需要float32或complex128类型值的地方则会强制需要一个明确的类型转换：
const Pi64 float64 = math.Pi

func TestConst2(t *testing.T) {
	var x float32 = float32(Pi64)
	var tmp = Pi64
	var y int32 = int32(tmp)
	t.Log(x, y)
	var z float64 = Pi64
	var c complex128 = complex128(Pi64)
	t.Log(z, c)
}
