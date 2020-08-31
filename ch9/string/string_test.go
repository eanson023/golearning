package string

import (
	"testing"
)

func TestString(t *testing.T) {
	//字符串声明
	var s string
	t.Log(s) //初始化为默认零值""
	s = "hello"
	t.Log(len(s))
	//s[1] = '3'         //string是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	// s="\xE4\xBA\xB5\xFF"
	t.Log(s, "长度:", len(s))
	s = "中"
	t.Log(len(s)) //是byte数

	//取出字符串的unicode
	c := []rune(s)
	t.Log(len(c))
	//t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF-8 %x", s)
}

func TestStrongToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		//[1]代表都是和第一个数对应
		t.Logf("%[1]c %[1]d", c)
	}
	for i := range s {
		t.Log(s[i])
	}
}
