// server3 更详细的显示请求信息
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/eanson023/golearning/gopl/ch1/lissajous/lissa"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissa", lissajous)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 处理程序回显请求的URL的请求部分
func handler(w http.ResponseWriter, r *http.Request) {
	countAdd()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter回显目前为止调用的次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count = %d\n", count)
	mu.Unlock()
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	countAdd()
	lissa.Lissajous(w)
}

func countAdd() {
	mu.Lock()
	count++
	mu.Unlock()
}
