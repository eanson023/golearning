package ch31

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 多线程情况下一段代码只执行一次
// 当然也可以在init函数里面写 但那就不是懒汉式的了
var once sync.Once

type Singleton struct {
}

var singleInstance *Singleton

func GetSingletonObj() *Singleton {
	// 只执行一次
	once.Do(func() {
		fmt.Println("create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			// 地址都一样
			fmt.Println(unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
