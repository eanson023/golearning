package passing_ref

import "testing"

//BenchmarkPassingArrayWithValue-8       	      68	  16302013 ns/op
//BenchmarkPassingArrayWithReference
//BenchmarkPassingArrayWithReference-8   	1000000000	         0.315 ns/op

const NumOfElems int = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NumOfElems]Content) int {
	return 0
}

func withReference(arr *[NumOfElems]Content) int {
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

//go test -bench=BenchmarkPassingArrayWithReference -trace=trace_1.out

func BenchmarkPassingArrayWithReference(b *testing.B) {
	var arr [NumOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}
