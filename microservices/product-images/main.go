package main

import (
	"context"
	"flag"
	"github.com/eanson023/golearning/microservices/product-images/files"
	"github.com/eanson023/golearning/microservices/product-images/handlers"
	rillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	bindAddress = flag.String("BIND_ADDRESS", ":8080", "Bind address for the server")
	logLevel    = flag.String("LOG_LEVEL", "debug", "Log output level for the server [debug,info,trace]")
	basePath    = flag.String("BASE_PATH", "tmp/images", "Base path to save images")
)

// 使用下面的命令上传文件
// curl -vv localhost:8080/images/1/test.png -X POST --data-binary @maozi.png
func main() {
	flag.Parse()

	logger := log.New(os.Stdout, "[product-images] ", -1)
	// 创建本地存储对象
	local, err := files.NewLocal(*basePath, 1024*1024*2)
	if err != nil {
		logger.Fatal(err)
	}
	files := handlers.NewFiles(local, logger)

	router := mux.NewRouter()
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", files.UploadREST)
	postRouter.HandleFunc("/", files.UploadMulipart)

	// 文件系统
	getRouter := router.Methods(http.MethodGet).Subrouter()
	// StripPrefix返回一个处理器，该处理器会将请求的URL.Path字段中给定前缀prefix去除后再交由h处理。
	// StripPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found。
	fileHandler := http.StripPrefix("/images/", http.FileServer(http.Dir(*basePath)))
	getRouter.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fileHandler)

	// CORS 跨域资源访问
	corsHanler := rillaHandlers.CORS(rillaHandlers.AllowedOrigins([]string{"http://localhost:4200"}))

	server := &http.Server{
		Addr:         *bindAddress,
		Handler:      corsHanler(router),
		ErrorLog:     logger,
		IdleTimeout:  120 * time.Second, //max time for connections using TCP Kepp-Alice
		ReadTimeout:  1 * time.Second,   //max time to reead request from the client
		WriteTimeout: 1 * time.Second,   //max time to write response to the client
	}
	go func() {
		logger.Printf("starting server on port: %s", server.Addr)
		// block
		err := server.ListenAndServe()
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
	server.Shutdown(tc)

}
