package leetcode

import (
	"math"
	"testing"
)

func myAtoi(str string) int {
	forIdx := 0
	isFushu, haveOpt := false, false
	var c byte
	for i := 0; i < len(str); i++ {
		c = str[i]
		forIdx = i
		if c == ' ' {
			continue
		}
		if c >= '0' && c <= '9' {
			break
		} else if c == '-' {
			isFushu = true
			haveOpt = true
			break
		} else if c == '+' {
			haveOpt = true
			break
		} else {
			return 0
		}
	}
	num := 0
	if haveOpt {
		forIdx++
	}
	for i := forIdx; i < len(str); i++ {
		c = str[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c) - 48
			if num >= math.MaxInt32 {
				break
			}
		} else {
			break
		}
	}
	if isFushu {
		num = -num
	}
	if num > math.MaxInt32 {
		num = math.MaxInt32
	} else if num < math.MinInt32 {
		num = math.MinInt32
	}
	return num
}
func Test8(t *testing.T) {
	t.Log(myAtoi("   -42"))
}
