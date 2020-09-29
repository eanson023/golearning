package ch36

import (
	"testing"
)

const SIZE int = 1000000

// 测试两种for循环 不相上下
// BenchmarkFor1
// BenchmarkFor1-8   	    4446	    271655 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFor2
// BenchmarkFor2-8   	    4360	    268875 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFor1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var arr [SIZE]int
		for j := range arr {
			_ = arr[j]
		}
	}
	b.StopTimer()
}

func BenchmarkFor2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var arr [SIZE]int
		for _, e := range arr {
			_ = e
		}
	}
	b.StopTimer()
}
