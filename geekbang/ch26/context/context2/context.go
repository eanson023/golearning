package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "w1")
	go worker(ctx, "w2")
	go worker(ctx, "w3")
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("canceld")
}

func worker(ctx context.Context, name string) {
	wg.Add(1)
	for range time.Tick(time.Second * 2) {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %s done\n", name)
			wg.Done()
			return
		default:
			fmt.Printf("worker:%s do work\n", name)
		}
	}
}
