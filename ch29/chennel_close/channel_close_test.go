package chennel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//channel的关闭和广播
//问题1：Recevier不知道Producer有多少数据 信道何时关闭
// 问题2: 假如有多个Recevier, 及时Producer发出了关闭消息 但只有一个收到 其它的接收信道该怎么办

// 所以go诞生关闭channel
// close(channel) 及时接收信道在等待时 它也会返回

///生产者消费者

///发出数据
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	time.Sleep(time.Second * 5)
	ch <- 99
	close(ch)
	//往关闭的信道发送消息会产生panic
	// panic: send on closed channel
	// ch <- 99
	wg.Done()
}

//接收数据
func dataRecevier(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			//多了会阻塞 一直等信道另外一边发送
			// 两个返回 第一个数据 第二个布尔值  当为false时代表channel是关闭状态了
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				wg.Done()
				break
			}
		}
	}()
}
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	// 必须在这家 在函数内加 可能还未执行就已经调用下面的Wait()函数了
	wg.Add(1)
	go dataProducer(ch, &wg)
	// 两个recivier
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go dataRecevier(ch, &wg)
	}
	wg.Wait()
}
