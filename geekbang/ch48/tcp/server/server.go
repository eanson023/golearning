package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	/**
	Unix网编程步骤：Server->Bind->Listen->Accepts
	Go语言简化为了：Listen->Accept
	*/

	//此处创建了第一个套接字:设置了通信协议、IP地址、port
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	// 套接字也是文件，需要关闭
	defer listener.Close()
	if err != nil {
		fmt.Println("net listen err", err)
		return
	}
	/**
	服务端接收了一次请求后即关闭了，显然不符合现在服务端能够同时接收大量请求的业务要求。使用for循环不断创建连接，等待带新的请求即可实现多客户端接入。
	具体的业务逻辑则可以交给一个go协程处理，这样服务端就可以专门用于循环等待请求、创建连接，而每个go程则负责具体的业务逻辑！
	*/
	// 监听用户连接请求
	for {
		//此处创建第二个套接字：用于阻塞监听客户端连接请求。注意listener并未监听，accept实现了监听
		conn, err := listener.Accept()
		//defer conn.Close()
		if err != nil {
			fmt.Println("listener accept err:", err)
			return
		}
		//	业务逻辑
		go handler(conn)

	}
	//运行服务端后，使用命令行工具模拟请求：`nc 127.0.0.1 3000`
}

func handler(conn net.Conn) {
	if conn == nil {
		panic("conn is null")
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	for {
		// 读取客户端数据
		n, err := conn.Read(buf)
		if err == io.EOF { //此时n=0
			fmt.Println("read EOF")
			break
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		//业务逻辑
		fmt.Println("Read msg:", string(buf[:n]))
		conn.Write([]byte("i'm  come from server"))
	}

	//
}
