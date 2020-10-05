package util_response_reply

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

// 仅需任意任务完成
// 比如说:在各个搜索引擎中返回词条 其中任意一个响应则返回结果
func FirstResponse() string {
	numOfRounner := 10
	// 缓冲信道 防止协程泄露 理解 将数据缓存到信道内 这样不产生阻塞 协程就会退出
	channel := make(chan string, numOfRounner-1)
	for i := 0; i < numOfRounner; i++ {
		go func(i int) {
			ret := runTask(i)
			channel <- ret
		}(i)
	}
	return <-channel
}

func runTask(id int) string {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}
func TestFirstResponse(t *testing.T) {
	// 输出当前系统中协程数
	t.Log("Before", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After", runtime.NumGoroutine())
}
