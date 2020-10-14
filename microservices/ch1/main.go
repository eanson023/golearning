package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 贪婪匹配
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(req.Body)
		if err != nil {
			// Nick suggest to use status.go 的http status
			http.Error(resp, "Ooops", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)
		// back to response writer
		fmt.Fprintf(resp, "Hello %s", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})
	http.ListenAndServe(":9090", nil)
}
