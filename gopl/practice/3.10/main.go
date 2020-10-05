package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234"))
}

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)
	i := l % 3
	// 写第一组 最多3个数
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])
	// 处理剩下的
	for ; i < l; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
