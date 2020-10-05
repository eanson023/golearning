package midinum

import (
	"strings"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	max := 1
	for pointer1, pointer2 := 0, 1; pointer2 < len(s); pointer2++ {
		if samePos := strings.Index(s[pointer1:], s[pointer2:pointer2+1]); samePos < pointer2-pointer1 {
			pointer1 = samePos + pointer1 + 1
		}
		max = Max(max, pointer2-pointer1+1)
	}
	return max
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func TestName(t *testing.T) {
	var s = "bbbbb"
	t.Log(lengthOfLongestSubstring(s))
}
