package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 设置访问路由
	http.HandleFunc("/", SayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler会创建一个goroutine来为其提供服务，而且连续请求3次，request的地址也是不同的：
func SayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(&request)
	go func() {
		// 假定请求需要耗时2s，在请求2s后返回，我们期望监控goroutine在打印2次Current request is in progress后即停止。但运行发现，监控goroutine打印2次后，其仍不会结束，而会一直打印下去。\
		// 问题出在创建监控goroutine后，未对其生命周期作控制，下面我们使用context作一下控制，即监控程序打印前需检测request.Context()是否已经结束，若结束则退出循环，即结束生命周期。
		for range time.Tick(time.Second) {
			select {
			case <-request.Context().Done():
				fmt.Println("request is outgoing")
				return
			default:
				fmt.Println("Current request is in progress")
			}
		}
	}()
	time.Sleep(2 * time.Second)
	writer.Write([]byte("hi"))
}
