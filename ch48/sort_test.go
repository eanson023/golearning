package ch48

import (
	"sort"
	"testing"
)

// 排序函数测试

type Person struct {
	Name string
	Age  int
}
type PersonSlice []Person

// 重写Len()方法
func (a PersonSlice) Len() int {
	return len(a)
}

// 重写Swap方法
func (a PersonSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 重写Less()方法 从大到小排序
func (a PersonSlice) Less(i, j int) bool {
	return a[j].Name < a[i].Name
}
func TestSort(t *testing.T) {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}
	t.Log(people)
	sort.Sort(PersonSlice(people)) // 按照 Age 的逆序排序
	t.Log(people)
	sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 的升序排序
	t.Log(people)

	s := "2017-2018 学年 第 2 学期"
	s2 := "2017-2018 学年 第 1 学期"
	s3 := "2018-2019 学年 第 2 学期"
	t.Log(s < s2)
	t.Log(s < s3)

}
