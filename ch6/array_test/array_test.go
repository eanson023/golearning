package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr[0] = 1
	t.Log(arr[0])
	t.Log(arr)

	arr2 := [4]int{1, 2, 3}
	t.Log(arr2)

	//	如果不知道数组长度 可以使用...代替
	arr3 := [...]uint{4, 5, 6}
	t.Log(arr3)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//	遍历方式1
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	t.Log("------------------")
	//遍历方式 2
	//下标 和 元素
	for idx, e := range arr {
		t.Log(idx, e)
	}
	t.Log("------------------")
	//因为go语言有严格的要求 变量定义不使用会报错
	//可以使用下划线_来忽略某一个变量
	//遍历方式 2 pro
	for _, e := range arr {
		t.Log(e)
	}
}

func TestMulWeiArr(t *testing.T) {
	c := [2][2]uint{{1, 2}, {3, 4}}
	for idx := range c {
		for _, e := range c[idx] {
			t.Log(e)
		}
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Logf("%T %d %d", arr, len(arr), cap(arr))
	arr2 := arr[:3]
	//数组截取后就变成了切片
	t.Logf("%T", arr2)
	//会将原来数组改变 共享存储结构
	arr2 = append(arr2, 99)
	//len 4  cap 9
	t.Log(arr2, len(arr2), cap(arr2))
	t.Log(arr[3:])
	t.Log(arr[1:])
	t.Log(arr[2:5])
	t.Log(arr)
}
