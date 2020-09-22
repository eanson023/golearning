package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", helloworld)
	// 静态文件管理
	files := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
