package ch36

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatByStringAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5", "6"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("123456", ret)
}

func TestConcatStringByBuffer(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5", "6"}
	var buf bytes.Buffer
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("123456", buf.String())
}

//测试性能
func BenchmarkConcatByStringAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5", "6"}
	b.ResetTimer()
	ret := ""
	for i := 0; i < b.N; i++ {
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

//测试性能
func BenchmarkConcatStringByBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5", "6"}
	b.ResetTimer()
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		for _, elem := range elems {
			buf.WriteString(elem)
		}
		buf.Reset()
	}
	b.StopTimer()
}

/*
goos: darwin
goarch: amd64
pkg: go_learning/ch36
BenchmarkConcatByStringAdd
BenchmarkConcatByStringAdd-8      	   48384	    111075 ns/op	  893480 B/op	       6 allocs/op //-8:8个CPU执行 执行了48384次 111075 ns/op:每次操作所耗时间111075纳秒 893480 B/op:每次分配了893480字节内存 6 allocs/op：每次分配6个对象
BenchmarkConcatStringByBuffer
BenchmarkConcatStringByBuffer-8   	23901672	        50.4 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: [no statements]
ok  	go_learning/ch36	6.897s

*/
