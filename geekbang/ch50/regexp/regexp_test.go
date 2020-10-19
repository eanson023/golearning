package regexp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
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

func TestRegexp5(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}

func TestRegexp6(t *testing.T) {
	a := "I am learning Go language"
	re, _ := regexp.Compile("[a-z]{2,4}")

	// 查找符合正则的第一个
	one := re.Find([]byte(a))
	assert.Equal(t, "am", string(one))

	// 查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAllString(a, -1)
	// [am lear ning lang uage]
	t.Log(all)

	// 查找符合条件的index位置,开始位置和结束位置
	index := re.FindIndex([]byte(a))
	assert.Equal(t, []int{2, 4}, index)

	//查找符合条件的所有的index位置，n同上
	allindex := re.FindAllStringIndex(a, -1)
	// [[2 4] [5 9] [9 13] [17 21] [21 25]]
	t.Log(allindex)
}

// 查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
//下面的输出第一个元素是"am learning Go language"
//第二个元素是" learning Go "，注意包含空格的输出
//第三个元素是"uage"
func TestRegexp7(t *testing.T) {
	a := "I am learning Go language"
	re2, _ := regexp.Compile("am(.*)lang(.*)")
	submatch := re2.FindSubmatch([]byte(a))
	t.Log(submatch)
	for _, v := range submatch {
		t.Log(string(v))
	}
}
