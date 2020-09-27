package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	// 据我分析 应该是在连接关闭时 向服务端发送的结束标识 这样服务端才有io.EOF
	defer conn.Close()
	for {
		//	 主动向服务器发送数据
		conn.Write([]byte("are you kidding me?"))

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client read err:", err)
			return
		}
		fmt.Println("client receive:", string(buf[:n]))
		time.Sleep(time.Second * 1)
	}
}
