package main

import (
	"fmt"
)

// 就地反转一个整型slice中的元素
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	// 反转整个数组
	a := [...]int{0, 1, 2, 3, 4, 5}
	// 转切片
	reverse(a[:])
	fmt.Println(a)

	// 将一个slice左移n个元素的简单方法就是连续调用reverse函数3次
	s := []int{1, 2, 3, 4, 5}
	// 向左移动2个元素
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
	var runes []rune
	for _, r := range "Hello,世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
