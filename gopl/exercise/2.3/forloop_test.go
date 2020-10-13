// forloop 测试两种循环的效率 结论 差不多 可能做过编译器优化
package forloop

import (
	"testing"
)

// pc[i] 是i的种群统计
var pc [256]byte

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range pc {
			pc[j] = pc[j/2] + byte(j&1)
		}
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, _ := range pc {
			pc[j] = pc[j/2] + byte(j&1)
		}
	}
}
