// 练习1.3: 尝试测量可能低效的程序和使用strings.Join的程序在执行时间上的差异
package practice

import (
	"strings"
	"testing"
)

// 测试结果
// goos: darwin
// goarch: amd64
// pkg: gopl/ch1/practice
// BenchmarkStringAdd
// BenchmarkStringAdd-8    	 6810889	       169 ns/op	      40 B/op	       3 allocs/op
// BenchmarkStringJoin
// BenchmarkStringJoin-8   	18003474	        66.7 ns/op	      16 B/op	       1 allocs/op
// PASS
// ok  	gopl/ch1/practice	2.610s

func BenchmarkStringAdd(b *testing.B) {
	args := []string{"are", "you", "ok", "?"}
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range args {
			s += sep + arg
			sep = " "
		}
	}
}

func BenchmarkStringJoin(b *testing.B) {
	args := []string{"are", "you", "ok", "?"}
	for i := 0; i < b.N; i++ {
		_ = strings.Join(args, " ")
	}
}
