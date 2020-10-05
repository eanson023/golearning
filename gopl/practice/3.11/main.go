package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	a := "-12345.0123"
	s := comma(a)
	fmt.Println(s)
}

func comma(s string) string {
	var buf bytes.Buffer
	if s[0] == '-' {
		buf.WriteByte('-')
		s = s[1:]
	}
	arr := strings.Split(s, ".")
	for _, v := range arr {
		l := len(v)
		i := l % 3
		// 写第一组 最多3个数
		if i == 0 {
			i = 3
		}
		buf.WriteString(v[:i])
		// 处理剩下的
		for ; i < l; i += 3 {
			buf.WriteByte(',')
			buf.WriteString(v[i : i+3])
		}
		buf.WriteByte('.')
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}
