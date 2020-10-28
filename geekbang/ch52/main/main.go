package main

import (
	"github.com/eanson023/golearning/geekbang/ch52"
	"github.com/eanson023/golearning/geekbang/ch52/route"
	"net/http"
)

func main() {
	r := ch52.GetGlobalRouter()
	r.RegisterRoute(new(route.User))
	r.RegisterRoute(new(route.Stuff))
	http.ListenAndServe(":8080", r.Router)
}
