package main

import (
	"fmt"
)

// 无类型常量 编译器将这些从属类型待定的常量表示层某些值
const (
	KiB = 1024 << (10 * iota)
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB //1<<64
	YiB
)

func main() {
	fmt.Println(KiB)
	fmt.Println(MiB)
	// 正常 虽然看似越界了
	fmt.Println(YiB / ZiB)
}
