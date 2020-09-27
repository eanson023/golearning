package main

import (
	"fmt"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	var buf = make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[0:n]))
}
