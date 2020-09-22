package pipe_filter

import (
	"testing"
)

func TestStraightPipeline(t *testing.T) {
	spfer := NewSplitFilter(",")
	tifer := NewToIntFilter()
	sumfer := NewSumFilter()
	straghtPipeline := NewStraightPipeline("MyPipeline", spfer, tifer, sumfer)
	// 这个其实也可以理解为大的filter里面装着一些小的filter 牛批
	if resp, err := straghtPipeline.Process("1,2,3,4,5,6,7,8,9"); err != nil {
		t.Log(err)
	} else {
		t.Log(resp)
	}
}
