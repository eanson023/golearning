package ch52

import (
	"github.com/gorilla/mux"
	"net/http"
)

type RouterSubject interface {
	RegisterRoute(observer Observer)
}

type Observer interface {
	GetRoutes() []*RouteEntry
}

type RouteEntry struct {
	pattern, method string
	handler         http.HandlerFunc
}

func NewRouteEntry(pattern, method string, hanlder func(http.ResponseWriter, *http.Request)) *RouteEntry {
	return &RouteEntry{
		pattern: pattern,
		handler: hanlder,
		method:  method,
	}
}

type Router struct {
	Router *mux.Router
}

func NewRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
	}
}

var globalRouter = NewRouter()

func RegisterRoute(observer Observer) {
	globalRouter.RegisterRoute(observer)
}

func (r *Router) RegisterRoute(observer Observer) {
	res := observer.GetRoutes()
	for _, re := range res {
		r.Router.Methods(re.method).Subrouter().Handle(re.pattern, re.handler)
	}
}

func GetGlobalRouter() *Router {
	return globalRouter
}
