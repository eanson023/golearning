package ch14

import (
	"fmt"
	"strconv"
	"testing"
)

// 切片(slice)解决了数组长度不能扩展，以及基本类型数组传递时产生副本的问题。
func TestSlice1(t *testing.T) {
	var s1 []int //和声明数组一样，只是没有长度，但是这样做没有实际意义，因为底层的数组指针为nil
	s2 := []byte{'a', 'b', 'c'}
	t.Log(s1, s2)
}

// 使用make函数创建
func TestSliceMake(t *testing.T) {
	//创建长度为5，容量为5，初始值为0的切片
	slice1 := make([]int, 5)
	//创建长度为5，容量为7，初始值为0的切片
	slice2 := make([]int, 5, 7)
	//创建长度为5，容量为5，并已经初始化的切片
	slice3 := []int{1, 2, 3, 4, 5}
	t.Log("slice1:", slice1)
	t.Log("slice2", slice2)
	t.Log("slice3", slice3)
	// 从数组创建：slice可以从一个数组再次声明。slice通过array[i:j]来获取，
	// 其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i:

	//声明一个含有10个元素类型为byte的数组
	var arr = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	//声明两个含有byte的slice
	var a, b []byte

	//a指向数组的第3个元素开始，并到第5个元素结束，现在a含有的元素:arr[2],arr[3],arr[4]
	a = arr[2:5]
	//b是数组arr的另一个slice,b的元素是:arr[3]和arr[4]
	b = arr[3:5]
	t.Log("------------------------------")
	t.Log("切片a:", a)
	t.Log("切片b:", b)
}

func TestSliceMakeFromSlice(t *testing.T) {
	oldSlice := []int{1, 2, 3}
	newSlice := oldSlice[:6]
	// runtime error: slice bounds out of range [:6] with capacity 3 [recovered]
	// 如果选择的旧切片长度超出了旧切片的cap()值（切片存储长度），则不合法
	t.Log(newSlice)
}

// 二 切片常见操作

// 切片空间与元素个数：
func TestUsualFunc(t *testing.T) {
	slice1 := make([]int, 5, 10)
	t.Log(len(slice1))
	t.Log(cap(slice1))
	t.Log(slice1)
}

// 切片操作
func TestUsualFunc2(t *testing.T) {
	// 切片增加
	slice1 := make([]int, 5, 10)
	slice1 = append(slice1, 1, 2)
	// 0 0 0 0 0 1 2
	t.Log(slice1)
	t.Log("-----------------------")
	// 切片增加一个新切片
	sliceTmp := []int{1, 2, 3}
	slice1 = append(slice1, sliceTmp...)
	t.Log(slice1)
	t.Log("-----------------------")
	//切片拷贝
	s1 := []int{2, 4, 6, 8}
	//必须给足充足的空间 否则只拷贝最大长度
	s2 := make([]int, 10)
	//拷贝拷贝成功的数量 将s1拷贝到s2
	var num int = copy(s2, s1)
	t.Log(s2, "拷贝拷贝成功的数量", num)

	//切片中删除元素
	t.Log("--------------------")
	s1 = []int{1, 3, 6, 9, 11, 24}
	for index, v := range s1 {
		t.Log("遍历:"+strconv.Itoa(index), v)
	}
	//删除该位置的元素
	index := 2
	s1 = append(s1[:index], s1[index+1:]...)
	t.Log(s1)
}

// 四 切片作为函数参数
func TestSliceAsArg(t *testing.T) {
	s := make([]int, 3)
	test(s)
	t.Logf("main----%p\n", s)
	t.Log("main----", s)
	t.Log("________________________")
	test2(&s)
	t.Logf("main----%p\n", s)
	t.Log("main----", s)
}

//通过切片的方式作为函数参数 仿佛是值传递
func test(s []int) {
	// 打印与main函数相同的地址
	fmt.Printf("test----%p\n", s)
	s = append(s, 1, 2, 3, 4, 5)
	// 一旦append的数据超过切片长度，则会打印新地址
	fmt.Printf("test----%p\n", s)
	// [0 0 0 1 2 3 4 5]
	fmt.Println("test---", s)
}

//通过传递指针可以改变 其地址指向的值
func test2(s *[]int) {
	// 打印与main函数相同的地址
	fmt.Printf("test----%p\n", *s)
	a := *s
	*s = append(a, 1, 2, 3, 4, 5)
	// 一旦append的数据超过切片长度，则会打印新地址
	fmt.Printf("test----%p\n", *s)
	// [0 0 0 1 2 3 4 5]
	fmt.Println("test---", *s)
}
