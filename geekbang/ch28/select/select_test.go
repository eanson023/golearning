package select_test

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

func Service2() chan string {
	channel := make(chan string, 1)
	channel <- "你好"
	time.Sleep(time.Second * 2)
	return channel
}

// 多路选择
func TestAsyncService(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log("case1", ret)
	case ret2 := <-Service2():
		t.Log("case2", ret2)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
