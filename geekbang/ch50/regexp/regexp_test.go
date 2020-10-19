package regexp

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRegexp1(t *testing.T) {
	var s = "A ghost boooooed"
	// *等价于{0,}0到多次
	if b, err := regexp.Match(`bo*`, []byte(s)); err != nil {
		t.Fatal(err)
	} else {
		assert.True(t, b)
	}

	// 会匹配b字符
	s = "A bird warbled"
	b, _ := regexp.MatchString(`bo*`, s)
	assert.True(t, b)

	s = "an A"
	// ^:匹配输入的开始
	b, _ = regexp.MatchString(`^A`, s)
	assert.False(t, b)
	s = "An E"
	b, _ = regexp.MatchString(`^A`, s)
	assert.True(t, b)
	b, _ = regexp.MatchString(`E$`, s)
	// 匹配输入的结束。
	assert.True(t, b)
}

func TestRegexp3(t *testing.T) {
	s := "candy"
	// + 代表匹配一个或多个
	b, _ := regexp.MatchString(`a+.*y$`, s)
	assert.True(t, b)
}

func TestRegexp4(t *testing.T) {
	// 单用:匹配前面一个表达式 0 次或者 1 次。等价于 {0,1}。
	s := "angle"
	// 匹配 le
	b, _ := regexp.MatchString(`e?le?`, s)
	assert.True(t, b)
	s = "oslo"
	b, _ = regexp.MatchString(`e?le?`, s)
	// 匹配l
	assert.True(t, b)

	// 非贪婪模式: 紧跟在任何量词 *、 +、? 或 {} 的后面，将会使量词变为非贪婪（匹配尽量少的字符），和缺省使用的贪婪模式（匹配尽可能多的字符）正好相反
	s = "123abc"
	// 匹配出123
	b, _ = regexp.MatchString(`\d+`, s)
	assert.True(t, b)

	s = "123abc"
	// 匹配出1
	b, _ = regexp.MatchString(`\d+?`, s)
	assert.True(t, b)
}
