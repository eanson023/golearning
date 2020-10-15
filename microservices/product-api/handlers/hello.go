package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{
		l: l,
	}
}

// ServeHTTP implements the go http.Hanlder interface
func (h *Hello) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// Nick suggest to use status.go 的http status
		http.Error(resp, "Ooops", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s\n", d)
	// back to response writer
	fmt.Fprintf(resp, "Hello %s", d)
}
