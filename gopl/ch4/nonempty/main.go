// nonempty 演示了slice的就地修改算法
package main

import (
	"fmt"
)

// nonempty 返回一个新的slice,slice中的元素都是非空字符串
// 在函数的调用过程中，底层数组的元素发生了改变
func nonempty(strings []string) []string {
	var i int
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings[:i]
}
func nonempty2(strings *[]string) {
	var i int
	for _, v := range *strings {
		if v != "" {
			(*strings)[i] = v
			i++
		}
	}
	*strings = (*strings)[:i]
}

var strs = []string{"One", "", "Two", "Three"}

func main() {
	strings := []string{"One", "", "Two", "Three"}
	fmt.Printf("%q\n", nonempty(strings))
	fmt.Printf("%q\n", strings)
	nonempty2(&strs)
	fmt.Printf("%q\n", strs)
}
