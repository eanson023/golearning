package util_all_done

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

// 当所有任务都完成才返回
func AllResponse() string {
	numOfRounner := 10
	channel := make(chan string)
	for i := 0; i < numOfRounner; i++ {
		go func(i int) {
			ret := runTask(i)
			channel <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRounner; j++ {
		finalRet += <-channel + "\n"
	}
	return finalRet
}

func runTask(id int) string {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}
func TestFirstResponse(t *testing.T) {
	// 输出当前系统中协程数
	t.Log("Before", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After", runtime.NumGoroutine())
}
