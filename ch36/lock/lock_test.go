package lock

import (
	"fmt"
	"sync"
	"testing"
)

var cache map[string]string

const NUM_OF_READER int = 40
const READ_TIMES int = 100000

func init() {
	cache = make(map[string]string)
	cache["a"] = "aa"
	cache["b"] = "bb"
}

func lockFreeAccess() {
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	for i := 0; i < NUM_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 只读加锁对协程有影响吗？
// 答案：
//BenchmarkLockFree-8               	     192	   6207104 ns/op
//BenchmarkLocalAccess
//BenchmarkLocalAccess-8            	       5	 230522312 ns/op
func lockAccess() {
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	m := new(sync.RWMutex)
	for i := 0; i < NUM_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				m.RLock()
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
				m.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkLockFree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockFreeAccess()
	}
	b.StopTimer()
}
func BenchmarkLocalAccess(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockAccess()
	}
	b.StopTimer()
}

// 查看性能消耗
// go test -bench=. -cpuprofile=cpu.prof
//  go tool pprof cpu.prof
// top
// list
