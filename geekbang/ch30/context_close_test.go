package ch30

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// Context与任务关闭

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	// 加waitGroup 证明下其它协程未退出
	var wg sync.WaitGroup
	// context.Backgroud()表示父节点
	// cancel是返回的函数
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
			wg.Done()
		}(i, ctx)
	}
	time.Sleep(time.Second * 1)
	cancel()
	wg.Wait()
}
