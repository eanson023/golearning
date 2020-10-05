package midinum

import (
	"math"
	"testing"
)

/*
public int myAtoi(String str) {
        int forIdx = 0;
        boolean isFushu = false, haveOpt = false;
        char c;
        for (int i = 0; i < str.length(); i++) {
            c = str.charAt(i);
            forIdx = i;
            if (c == ' ') {
                continue;
            }
            if (c >= '0' && c <= '9') {
                break;
            } else if (c == '-') {
                isFushu = true;
                haveOpt = true;
                break;
            } else if (c == '+') {
                haveOpt = true;
                break;
            } else return 0;

        }
        long num = 0;
        for (int i = haveOpt ? forIdx + 1 : forIdx; i < str.length(); i++) {
            c = str.charAt(i);
            if (c >= '0' && c <= '9') {
                num = num * 10 + c - 48;
                if (num >= Integer.MAX_VALUE)
                    break;
            } else break;
        }
        num = isFushu ? -num : num;
        if (num > Integer.MAX_VALUE) {
            num = Integer.MAX_VALUE;
        } else if (num < Integer.MIN_VALUE) {
            num = Integer.MIN_VALUE;
        }
        return (int) num;
    }
*/
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
