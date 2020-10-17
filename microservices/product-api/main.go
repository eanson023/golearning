package main

import (
	"context"
	"github.com/eanson023/golearning/microservices/product-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api ", log.LstdFlags)
	pd := handlers.NewProductsHandler(logger)
	// 不使用http里的defaulrServeMux,我们要自己创建一个ServeMux来处理http请求
	sm := http.NewServeMux()
	sm.Handle("/", pd)
	// 自定义server 我们可以做一些我们想做的东西（自定义参数）
	s := &http.Server{
		Addr:         ":8080",           //configure the bind address
		Handler:      sm,                //set the my handler
		ErrorLog:     logger,            //set the logger for the server
		IdleTimeout:  120 * time.Second, //max time for connections using TCP Kepp-Alice
		ReadTimeout:  1 * time.Second,   //max time to reead request from the client
		WriteTimeout: 1 * time.Second,   //max time to write response to the client
	}
	go func() {
		logger.Printf("starting server on port: %s", s.Addr)
		// block
		err := s.ListenAndServe()
		if err != nil {
			logger.Printf("error starting server:%s\n", err)
			os.Exit(1)
		}
	}()
	// 1:bind addr 2:http.Handler
	// http.ListenAndServe(":9090", sm)

	// 使用os/signal包里 通知某种信号来告知程序来关闭服务器
	sigChan := make(chan os.Signal)
	// 当收到终止或kill命令时 会向sigChan发送
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	// 在未收到信号前 这里是阻塞的
	sig := <-sigChan
	logger.Println("Recived terminate,graceful shutdown", sig)
	// 截止时间是现在+设置的绝对时间
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// 如果没有处理程序，则正常关机，如果30s后任然有请求发生。则强制关闭
	s.Shutdown(tc)
}
