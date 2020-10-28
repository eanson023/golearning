package ch51

import (
	"net/http"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HelloHanlderStruct struct {
	content string
}

func (handler *HelloHanlderStruct) ServeHTTP(http.ResponseWriter, *http.Request) {

}

// 实现了Handler接口
type HanlderFunc func(http.ResponseWriter, *http.Request)

// 我自己调自己
func (f HanlderFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	f(rw, r)
}

func Handle(pattern string, handler Handler) {
	DefaultServeMux.Handle(pattern, handler)
}

// 交给默认的路由复用器
func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

type ServeMux struct {
}

var DefaultServeMux ServeMux

func (mux *ServeMux) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}

func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// 这种叫类型强转？
	mux.Handle(pattern, HanlderFunc(handler))
}

func (mux *ServeMux) Handle(pattern string, handler Handler) {

}
