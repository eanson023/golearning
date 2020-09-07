package groutine

import (
	"fmt"
	"testing"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// go func(i int) {
		// 	fmt.Println("1-----", i)
		// 	//这是采用方法调用的方式 go的参数传递是值传递 所以i不同
		// }(i)
		go func() {
			// 这里是共享变量 存在锁的竞争 所以数据不一致
			fmt.Println("2-----", i)
		}()
	}
	//等待
	// time.Sleep(time.Millisecond * 50)
}
