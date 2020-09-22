package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "this is / pattern")
	})
	http.HandleFunc("/hello", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "this is /hello/ pattern")
	})
	http.HandleFunc("/world", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "this is /world/ pattern")
	})
	http.ListenAndServe(":8080", nil)
}
