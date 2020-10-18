package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx上下文对象，cancel是回调函数 context.Background()通常表示最上层的的根节点
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for t := range time.Tick(time.Second * 2) {
			fmt.Println(t)
			select {
			case <-ctx.Done():
				fmt.Println("exit")
				return
			default:
				fmt.Println("do go func work")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("stop the go func work")
	time.Sleep(time.Second)
}
