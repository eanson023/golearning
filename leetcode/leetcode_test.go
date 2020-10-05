package leetcode

import (
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	counter := 0
	for prev, next := 0, 1; next < len(nums); next++ {
		if nums[prev] != nums[next] {
			nums[prev+1] = nums[next]
			prev = prev + 1
			counter++
		}
	}
	return counter + 1
}

func TestRemoveDuplicates(t *testing.T) {
	t.Log(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
