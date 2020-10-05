package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	//切片申明 与数组声明的区别灭有指定长度 因为它是可变长的
	var s0 []int
	//长度 --0   capacity容量 --0
	t.Log(len(s0), cap(s0))
	t.Log("--------------------")
	//填充元素
	s0 = append(s0, 1)
	//长度 --1   capacity容量 --1
	t.Log(len(s0), cap(s0))
	t.Log("--------------------")

	//声明并初始化切片
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	t.Log("--------------------")
	//另外一种声明方式
	//make(类型,长度，容量)
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//越界s2[3]
	t.Log(s2[0], s2[1])
	t.Log("--------------------")
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[3])
	t.Log(len(s2), cap(s2))
}

//切片共享存储结构
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	//只要第二季度
	Q2 := year[3:6]
	//len:3 cap:9 从3开始到year的连续存储空间结束
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	//将summer值修改 再看Q2是否变化
	summer[0] = "June"
	//可以发现summer[0]修改了，Q2中共享的数据也受到了修改
	t.Log(Q2)
	t.Log(year)
}
func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4}
	//切片只能和nil比较
	//slice can only be compared to nil
	//if a == b {
	//	t.Log("equal")
	//}
	t.Log(a, b)
	c := []int{}
	if c == nil {
		t.Log("hello world")
	}

}
