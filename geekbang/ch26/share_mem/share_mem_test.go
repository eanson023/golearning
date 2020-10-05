package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	// 5000个协程
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	// 不等于5000
	t.Logf("counter=%d", counter)
}
func TestCounterSafe(t *testing.T) {
	// 同步锁
	var mut sync.Mutex
	// join
	var wg sync.WaitGroup
	counter := 0
	// 5000个协程
	for i := 0; i < 5000; i++ {
		go func() {
			wg.Add(1)
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			wg.Done()
			counter++
		}()
	}
	wg.Wait()
	//5000加锁的方式
	t.Logf("counter=%d", counter)
}
