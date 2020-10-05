package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask is done.")
}

// 串行化
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// futureTask
// 返回channel当调用者需要时才等待
func AsyncService() chan string {
	// 使用make函数 指定channel类型
	retCh := make(chan string, 1)
	// service启动协程去运行
	go func() {
		ret := service()
		fmt.Println("return result.")
		// 如果channel未指定大小那么这里会阻塞 因为要等到接收方接收数据才会关闭
		// 指定大小为1 buffer channel获取了传入就不会阻塞了
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}
func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	// 当需要的时候再拿出来 用箭头<-
	t.Log(<-retCh)
}
