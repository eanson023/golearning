package route

import (
	"github.com/eanson023/golearning/geekbang/ch52"
	"net/http"
)

type User struct {
}

func (u *User) Get(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("hello world"))
}

func (u *User) GetRoutes() []*ch52.RouteEntry {
	res := []*ch52.RouteEntry{}
	re := ch52.NewRouteEntry("/users", http.MethodGet, u.Get)
	res = append(res, re)
	return res
}
