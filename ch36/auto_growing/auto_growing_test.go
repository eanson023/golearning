package auto_growing

import "testing"

//BenchmarkAutoGrow
//BenchmarkAutoGrow-8         	    1610	    628937 ns/op
//BenchmarkProperInit
//BenchmarkProperInit-8       	    7592	    152684 ns/op
//BenchmarkProperSizeInit
//BenchmarkProperSizeInit-8   	    1623	    673704 ns/op

const times int = 1000
const numofElems int = 100000

func TestAutoGrow(t *testing.T) {
	for i := 0; i < times; i++ {
		var s []int
		for j := 0; j < numofElems; j++ {
			s = append(s, j)
		}
	}
}

func TestProperInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, numofElems)
		for j := 0; j < numofElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int
		for j := 0; j < numofElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numofElems)
		for j := 0; j < numofElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperSizeInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numofElems*8)
		for j := 0; j < numofElems; j++ {
			s = append(s, j)
		}
	}
}
