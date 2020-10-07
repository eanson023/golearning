// leetbook leetcode上的探索页
package leetbook

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode"
)

/*
买卖股票的最佳时机 II
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func maxProfit(prices []int) int {
	max := 0
	for left, right := 0, 1; right < len(prices); right++ {
		if prices[left] > prices[right] {
			left++
			continue
		}
		if prices[right]-prices[left] > max {
			max = prices[right] - prices[left]
		}
	}
	return max
}

func TestMaxProfit(t *testing.T) {
	s := []int{7, 1, 5, 3, 6, 4}
	assert.Equal(t, 5, maxProfit(s))
	s = []int{1, 2, 3, 4, 5}
	assert.Equal(t, 4, maxProfit(s))
	s = []int{7, 6, 4, 3, 1}
	assert.Equal(t, 0, maxProfit(s))
}
func reverseString(s []byte) {
	var s2 = make([]byte, len(s))
	copy(s2, s)
	for idx, i := 0, len(s)-1; i >= 0; i-- {
		s[idx] = s2[i]
		idx++
	}
}

func reverseString2(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func TestReverseString(t *testing.T) {
	var s = []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString2(s)
	t.Logf("%s", s)
}

/*
整数反转

重新做题
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnx13t/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func reverse(x int) int {
	res := 0
	for x != 0 {
		b := x % 10
		x = x / 10
		res = res*10 + b
	}
	if res > 1<<31-1 || res < (-1<<31) {
		res = 0
	}
	return res
}

func TestReverse(t *testing.T) {
	var x int32 = 1534236469
	res := reverse(int(x))
	// overflow
	assert.Equal(t, 0, res)
	x = -1234
	res = reverse(int(x))
	assert.Equal(t, -4321, res)
}

/*
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func firstUniqChar(s string) int {
	var mp [26]int
	for _, v := range s {
		i := v - 'a'
		mp[i]++
	}
	for i, v := range s {
		if mp[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

func TestFirstUniqChar(t *testing.T) {
	s := "leetcode"
	assert.Equal(t, 0, firstUniqChar(s))
	s = "loveleetcode"
	assert.Equal(t, 2, firstUniqChar(s))
}

/*
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var mp = make(map[rune]int)
	for _, v := range s {
		mp[v]++
	}
	for _, v := range t {
		mp[v]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}

// 不支持unicode
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var mp [26]int
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
		mp[t[i]-'a']--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	tt := "nagaram"
	assert.Equal(t, true, isAnagram2(s, tt))
	s = "a"
	assert.Equal(t, false, isAnagram2(s, tt))
}

/*
验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	c := []rune(s)
	l, r := 0, len(c)-1
	for l < r {
		if !unicode.IsLetter(c[l]) && !unicode.IsDigit(c[l]) {
			l++
			continue
		}
		if !unicode.IsLetter(c[r]) && !unicode.IsDigit(c[r]) {
			r--
			continue
		}
		if c[l] != c[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	assert.Equal(t, true, isPalindrome(s))
	s = "race a car"
	assert.Equal(t, false, isPalindrome(s))
}

/*
字符串转换整数 (atoi)

请你来实现一个 atoi 函数，使其能将字符串转换成整数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：

如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnoilh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if s == "" {
		return 0
	}
	runes := []rune(s)
	idx := 0
	var res int
	var isNeg bool
	if runes[0] == '-' {
		isNeg = true
		idx = 1
	}
	if runes[0] == '+' {
		idx = 1
	}
	for ; idx < len(runes); idx++ {
		if !unicode.IsDigit(runes[idx]) {
			break
		}
		res = res*10 + int(runes[idx]) - '0'
		if res > 1<<31-1 {
			if isNeg {
				return -1 << 31
			}
			return 1<<31 - 1
		}
	}
	if isNeg {
		res = -res
	}
	return res
}

func TestMyAtoi(t *testing.T) {
	s := "42 9"
	assert.Equal(t, 42, myAtoi(s))
	s = "   -42 "
	assert.Equal(t, -42, myAtoi(s))
	s = "4193 with words"
	assert.Equal(t, 4193, myAtoi(s))
	s = "words and 987"
	assert.Equal(t, 0, myAtoi(s))
	s = "-91283472332"
	assert.Equal(t, -1<<31, myAtoi(s))
	s = "91283472332"
	assert.Equal(t, 1<<31-1, myAtoi(s))
	s = "9223372036854775808"
	assert.Equal(t, 2147483647, myAtoi(s))
	s = "-9223372036854775808"
	assert.Equal(t, -1<<31, myAtoi(s))
}

/*
实现 strStr()
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnr003/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func TestStrStr(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	assert.Equal(t, 2, strStr(haystack, needle))
}

/*
外观数列
给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。

注意：整数序列中的每一项将表示为一个字符串。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
第一项是数字 1

描述前一项，这个数是 1 即 “一个 1 ”，记作 11

描述前一项，这个数是 11 即 “两个 1 ” ，记作 21

描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211

描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221



示例 1:

输入: 1
输出: "1"
解释：这是一个基本样例。
示例 2:

输入: 4
输出: "1211"
解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnpvdm/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func countAndSay(n int) string {
	var sb strings.Builder
	res := "1"
	for i := 2; i <= n; i++ {
		slice, target := make([]int, 1), make([]byte, 1)
		idx := 0
		for j := range res {
			slice[idx]++
			target[idx] = res[j]
			if j+1 < len(res) && res[j+1] != res[j] {
				slice, target = append(slice, 0), append(target, 0)
				idx++
			}
		}
		for i, v := range slice {
			sb.WriteByte(byte(v + '0'))
			sb.WriteByte(target[i])
		}
		res = sb.String()
		sb.Reset()
	}
	return res
}

func TestCountAndSay(t *testing.T) {
	assert.Equal(t, "312211", countAndSay(6))
}

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnmav1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for _, v := range strs {
		// !=0表示 	子串不是从下标0开始的
		for strings.Index(v, res) != 0 {
			// 不断缩小区间 知道匹配为止
			res = res[:len(res)-1]
		}
	}
	return res
}

func TestLongestCommonPrefix(t *testing.T) {
	var strs = []string{"flower", "flow", "flight"}
	assert.Equal(t, "fl", longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	assert.Equal(t, "", longestCommonPrefix(strs))
	strs = []string{"c"}
	assert.Equal(t, "c", longestCommonPrefix(strs))
	strs = nil
	assert.Equal(t, "", longestCommonPrefix(strs))
}
