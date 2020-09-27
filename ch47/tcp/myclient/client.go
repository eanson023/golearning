package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"sync/atomic"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("create dial error:", err)
	}
	defer conn.Close()
	var strings = [...]string{"我是eanson", "i'm eanson", "wdnmd", "are you kidding me", "hello world"}
	buf := make([]byte, 128)
	var wg sync.WaitGroup
	var errorFlag = false
	var count int32 = 0
	for {
		// 全双工通信验证
		wg.Add(2)
		go func() {
			defer wg.Done()
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("client reading data error:", err)
				errorFlag = true
				return
			}
			fmt.Printf("from server:%s\tdata:%s\tcount:%d\n", conn.RemoteAddr(), string(buf[:n]), atomic.AddInt32(&count, 1))
		}()
		go func() {
			defer wg.Done()
			fmt.Println("write count:", atomic.AddInt32(&count, 1))
			_, err = conn.Write([]byte(strings[rand.Intn(len(strings))]))
			if err != nil {
				fmt.Println("client writing data error:", err)
				errorFlag = true
				return
			}
		}()
		wg.Wait()
		if errorFlag {
			return
		}
		//time.Sleep(time.Second)
	}
}
