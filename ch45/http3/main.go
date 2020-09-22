package main

import (
	"fmt"
	"net/http"
)

// 很多场景中，路由的处理函数在执行前，要先进行一些校验，比如安全检查，错误处理等等，这些行为需要在路由处理函数执行前有限执行。

// 小小闭包 责任链模式?
func before(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("执行前置处理")
		handle(w, r)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test1")
}

func main() {
	http.HandleFunc("/", before(test))
	http.ListenAndServe(":8080", nil)
}
