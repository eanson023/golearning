package main

import (
	"fmt"
	"net"
	"time"
)

/**
Go语言包中处理UDP Socket和TCP Socket不同的地方就是在服务器端处理多个客户端请求数据包的方式不同,UDP缺少了对客户端连接请求的Accept函数。其他基本几乎一模一样，只有TCP换成了UDP而已。
*/
func main() {
	udp, err := net.ResolveUDPAddr("udp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.ListenUDP("udp", udp)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf = make([]byte, 128)
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("from:%s\tdata:%s", addr.String(), string(buf[:n]))
	conn.WriteToUDP([]byte(time.Now().String()), addr)
}
