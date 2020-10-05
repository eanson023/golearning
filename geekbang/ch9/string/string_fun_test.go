package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	//分割
	var parts = strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	for i := range parts {
		t.Log(parts[i])
	}
	news := strings.Join(parts, "-")
	t.Log(news)
}

func TestConvert(t *testing.T) {
	//整数转为字符数
	s := strconv.Itoa(10)
	t.Log(s)
	//字符串转数字 会返回两个值 一个是val 一个是error
	if val, err := strconv.Atoi(s); err == nil {
		t.Log(20 + val)
	}
}
