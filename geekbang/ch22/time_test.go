package ch22

import (
	"testing"
	"time"
)

// 时间操作
func Test1(t *testing.T) {
	//当前时间
	nowTime := time.Now()
	// time.Time
	t.Logf("%T\n", nowTime)
	t.Log(nowTime)
	//自定义时间
	customTime := time.Date(2008, 7, 15, 13, 30, 0, 0, time.Local)
	t.Log(customTime)
	t.Log(customTime.String())
}

// 时间格式化与解析
// Go的时间格式化必须传入Go的生日：`Mon Jan 2 15:04:05 -0700 MST 2006`
func Test2(t *testing.T) {
	nowTime := time.Now()
	// stringTime := nowTime.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	stringTime := nowTime.Format("2006年1月2日 15:04:05")
	t.Log(stringTime)

	// 时间解析
	stringTime = "2019-01-01 15:03:01"
	objTime, _ := time.Parse("2006-01-02 15:04:05", stringTime)
	t.Log(objTime)
}

// 获取 年 月 日
func Test3(t *testing.T) {
	nowTime := time.Now()
	year, month, day := nowTime.Date()
	t.Log(year, month, day)

	hour, min, sec := nowTime.Clock()
	t.Log(hour, min, sec)

	t.Log(nowTime.Year(), nowTime.Month(), nowTime.Hour())
	// 指今年一共过了多少天
	t.Log(nowTime.YearDay())
}

// ### 时间戳
// 时间戳是指计算时间距离 1970年1月1日的秒数：

// 时间间隔
func Test4(t *testing.T) {
	nowTime := time.Now()
	t.Log(nowTime.Unix())
	// 20s后        一年一个月加一天后
	t.Log(nowTime.Add(time.Second*20), nowTime.AddDate(1, 1, 1))
	// 贴士：
	// - 传入负数则是往前计算
	// - Sub()函数可以用来计算两个时间的差值

	// 时间睡眠
	time.Sleep(time.Second * 3) //让程序睡眠三秒钟
}

// 二 时间中的通道操作（定时器）
// 标准库中的Timer可以让用户自定义一个定时器，在用对select处理多个channel的超时、单channel读写的超时等情形时很方便：
