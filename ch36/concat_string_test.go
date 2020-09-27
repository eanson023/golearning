package ch36

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/**
高效字符串连接测试
BenchmarkSprintf-8         	   45663	     26161 ns/op
BenchmarkStringAdd
BenchmarkStringAdd-8       	  173791	      6693 ns/op
BenchmarkStringBuilder
BenchmarkStringBuilder-8   	 1000000	      1052 ns/op
BenchmarkBytesBuf
BenchmarkBytesBuf-8        	  261390	      5070 ns/op
*/
const numbers int = 100

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < numbers; j++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < numbers; j++ {
			s += strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j := 0; j < numbers; j++ {
			buf.WriteString(strconv.Itoa(i))
		}
		_ = buf.String()
	}
	b.StopTimer()
}
