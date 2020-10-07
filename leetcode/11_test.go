package leetcode

import "testing"

func maxArea(height []int) int {
	left, right, max := 0, len(height)-1, 0
	for left != right {
		curr := height[right]
		if curr > height[left] {
			curr = height[left]
		}
		curr *= right - left
		if max < curr {
			max = curr
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func Test11(t *testing.T) {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(height))
}
