package utf8

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func TestA(t *testing.T) {
	s := "Hello, 世界"
	t.Log(s)
	t.Log(utf8.RuneCountInString(s))
	// 解码字符
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		t.Logf("%d\t%c\n", i, r)
		i += size
	}
	var n int
	// 可以用简单的range循环统计字符串中的文字符号数目
	for range s {
		n++
	}
	assert.Equal(t, 9, n)
}

func TestB(t *testing.T) {
	// 日语片假名 “程序”
	s := "プログラム"
	// 第一个Printf里的谓词%x以十六进制形式输出，并在每两个数位间插入空格
	t.Logf("% x\n", s)
	// 当[]rune转换作用用于UTF-8编码的字符串时，返回该字符串的Unicode码点序列
	r := []rune(s)
	t.Logf("%x\n", r)
	for _, v := range r {
		t.Logf("%c\n", v)
	}
	// 将slice转为string
	t.Log(string(r))
}

func TestC(t *testing.T) {
	// "A",而不是"65"
	t.Log(string(rune(65)))
	// "京"
	t.Log(string(rune(0x4eac)))
	// 如果文字符号值非法，将被专门的替换字符取代（'\uFFFD'）
	t.Log(string(rune(1234567)))
}

func TestD(t *testing.T) {
	var s = "你好，Hello"
	r := []rune(s)
	r[0] = '您'
	r[4] = 'a'
	t.Log(string(r))
	b := []byte(s)
	b[0] = '1'
	t.Log(string(b))
}
