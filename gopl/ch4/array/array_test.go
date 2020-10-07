package array

import (
	"fmt"
	"testing"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func TestWTF(t *testing.T) {
	// 没想到数组还能指定下标初始化
	symbol := [...]string{USD: "$", GBP: "£", RMB: "￥", EUR: "€"}
	fmt.Println(RMB, symbol[RMB])
	// 定义100个元素的数组r,除了最后个元素是-1外，其它的都是0
	r := [...]int{99: -1}
	// [100]int,-1
	fmt.Printf("%T,%d\n", r, r[99])
}

/*
测试结果：利用取值符取值range更快,但直接申明是最快的
BenchmarkA
BenchmarkA-8   	256686823	         4.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkB
BenchmarkB-8   	68595060	        17.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkC
BenchmarkC-8   	1000000000	         0.299 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bt [32]byte
		zero1(&bt)
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bt [32]byte
		zero2(&bt)
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bt [32]byte
		zero3(&bt)
	}
}

func zero1(ptr *[32]byte) {
	for i := range *ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
func zero3(ptr *[32]byte) {
	*ptr = [32]byte{}
}
