// server1 是一个迷你回声服务器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 回声请求调用处理程序
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 处理程序回显请求URL r路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	// %q 带引号的字符
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}
