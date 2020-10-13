package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "asas老师啊  hfgh"
	s = string(reverseUTF8([]byte(s)))
	fmt.Println(s)
}

func rev(bs []byte) {
	l := 0
	r := len(bs) - 1
	for l < r {
		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}
}

// 反正utf8 太巧妙 首先局部反转每个rune 最后全局反转 就得到了 结果
func reverseUTF8(s []byte) []byte {
	for i := 0; i < len(s); {
		_, n := utf8.DecodeRune(s[i:])
		rev(s[i : i+n])
		i += n
	}
	rev(s)
	return s
}
