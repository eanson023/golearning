package ch29

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	// 加waitGroup 证明下其它协程未退出
	var wg sync.WaitGroup
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
			wg.Done()
		}(i, cancelChan)
	}
	// 这个只有一个信道能收到然后结束该协程 其它的都收不到
	// cancel1(cancelChan)
	cancel2(cancelChan)
	time.Sleep(time.Second * 1)
	wg.Wait()
}
