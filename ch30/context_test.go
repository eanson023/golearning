package ch30

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
Golang 中上下文的最大作用，在不同 Goroutine 之间对信号进行同步避免对计算资源的浪费，与此同时 Context 还能携带以请求为作用域的键值对信息。
*/

func TestContextWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// 500 2500
	go HelloHandle(ctx, 2500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("Hello Handle", ctx.Err())
	}

}

func HelloHandle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
