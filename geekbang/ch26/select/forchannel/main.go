// time.Tick()返回的是一个channel,每隔指定的时间会有数据从channel中出来，for range不仅能遍历map,slice,array还能取出channel中数据
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := NewChannel()
	go func() {
		for b := range ch {
			fmt.Println("中", b)
		}
	}()
	time.Sleep(time.Second * 2)
	ch <- true
	time.Sleep(time.Second * 2)
}

// NewChannel chan作为函数返回值的方式有3种:（chan int）、（<- chan int）、（chan <- int），分别代表（可读可写的管道）、（只读管道）、（只写管道），只读管道不能close()，只写管道可以close()
func NewChannel() chan bool {
	return make(chan bool)
}
