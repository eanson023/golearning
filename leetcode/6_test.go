package leetcode

import (
	"testing"
)

/*
public String convert(String s, int numRows) {
        if (numRows == 1) return s;

        StringBuilder ret = new StringBuilder();
        int n = s.length();
        int cycleLen = (numRows << 1) - 2;

        for (int i = 0; i < numRows; i++) {
            for (int j = 0; j + i < n; j += cycleLen) {
                ret.append(s.charAt(j + i));
                if (i != 0 && i != numRows - 1 && j + cycleLen - i < n)
                    ret.append(s.charAt(j + cycleLen - i));
            }
        }
        return ret.toString();
    }
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ret []byte
	var n, cycleLen = len(s), (numRows << 1) - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			ret = append(ret, s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				ret = append(ret, s[j+cycleLen-i])
			}
		}
	}
	return string(ret)
}

/*
public String convert(String s, int numRows) {
        if (numRows <= 1) {
            return s;
        }
        StringBuilder[] sb = new StringBuilder[numRows];
        for (int i = 0; i < sb.length; i++) {
            sb[i] = new StringBuilder();
        }
        //true向下 false向上
        boolean direction = false;
        for (int i = 0, curr = 0; i < s.length(); i++) {
            sb[curr].append(s.charAt(i));
            if (curr == 0 || curr == numRows - 1) {
                direction = !direction;
            }
            curr = direction ? ++curr : --curr;
        }
        for (int i = 1; i < numRows; i++) {
            sb[0].append(sb[i]);
        }
        return sb[0].toString();
    }
*/
func convert2(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	//多维切片用make函数
	var sb = make([][]byte, numRows)
	var direction = false
	for i, curr := 0, 0; i < len(s); i++ {
		sb[curr] = append(sb[curr], s[i])
		if curr == 0 || curr == numRows-1 {
			direction = !direction
		}
		if direction {
			curr++
		} else {
			curr--
		}
	}
	//切片合并到第一位
	for i := 1; i < numRows; i++ {
		sb[0] = append(sb[0], sb[i]...)
	}
	//转string
	return string(sb[0])
}
func Test6(t *testing.T) {
	t.Log(convert2("PAYPALISHIRING", 3))
}
