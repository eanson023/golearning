package main

import (
	"context"
	"flag"
	"github.com/eanson023/golearning/microservices/product-api/data"
	"github.com/eanson023/golearning/microservices/product-api/handlers"
	"github.com/go-openapi/runtime/middleware"
	rillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var port string

func init() {
	flag.StringVar(&port, "BIND_ADDRESS", ":8080", "server port")
}

func main() {
	flag.Parse()

	// 日志器
	logger := log.New(os.Stdout, "product-api ", log.LstdFlags)
	// 校验器
	validator := data.NewValiation()

	// productHandler  该结构体的方法 是HanlderFunc
	pd := handlers.NewProductsHandler(logger, validator)
	// 使用gorilla的mux HTTP多路复用器 它实现了http.Hanlder接口所以和 http.ServeMux完全兼容
	// http包中的defauleServeMux无法进行正则匹配 不能很好的构建RESTful service
	router := mux.NewRouter()
	// 路由

	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", pd.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", pd.GetProductSingle)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	// 利用正则匹配
	putRouter.HandleFunc("/products/{id:[0-9]+}", pd.UpdateProduct)
	// 使用中间件 检验json数据
	putRouter.Use(pd.MidllewareProductValidation)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", pd.AddProduct)
	postRouter.Use(pd.MidllewareProductValidation)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", pd.DeleteProduct)

	// 整合swagger
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS 跨域资源访问
	corsHanler := rillaHandlers.CORS(rillaHandlers.AllowedOrigins([]string{"http://localhost:4200"}))

	// 自定义server 我们可以做一些我们想做的东西（自定义参数）
	s := &http.Server{
		Addr:         port,               //configure the bind address
		Handler:      corsHanler(router), //使用corsHandler之后再set the my router handler这个语法有点打脑壳
		ErrorLog:     logger,             //set the logger for the server
		IdleTimeout:  120 * time.Second,  //max time for connections using TCP Kepp-Alice
		ReadTimeout:  1 * time.Second,    //max time to reead request from the client
		WriteTimeout: 1 * time.Second,    //max time to write response to the client
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
