package leetcode

import (
	"strconv"
	"testing"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var res, tmp = 0, x
	for tmp != 0 {
		res = res*10 + tmp%10
		tmp = tmp / 10
	}
	return res-x == 0
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	var s2 []byte
	for i := len(s) - 1; i >= 0; i-- {
		s2 = append(s2, s[i])
	}
	return s == string(s2)
}

func Test9(t *testing.T) {

}
