package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("net listener error:", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("create connection error:", err)
			return
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	var wg sync.WaitGroup
	var errorFlag = false
	//与客户端持续交互
	for {
		// 全双工通信验证
		wg.Add(2)
		go func() {
			defer wg.Done()
			n, err := conn.Read(buf)
			if err == io.EOF {
				fmt.Println("read EOF")
				return
			}
			if err != nil {
				fmt.Println("read error:", err)
				errorFlag = true
				return
			}
			fmt.Printf("from client:%s\tdata:%s\n", conn.RemoteAddr().String(), string(buf[:n]))
		}()
		go func() {
			defer wg.Done()
			if _, err := conn.Write([]byte("done")); err != nil {
				fmt.Println("writing data  error:", err)
				errorFlag = true
			}
		}()
		wg.Wait()
		if errorFlag {
			return
		}
	}
}
