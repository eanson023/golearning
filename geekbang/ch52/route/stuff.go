package route

import (
	"github.com/eanson023/golearning/geekbang/ch52"
	"net/http"
)

type Stuff struct {
}

func (s *Stuff) Get(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("hello world2"))
}

func (s *Stuff) GetRoutes() []*ch52.RouteEntry {
	return []*ch52.RouteEntry{ch52.NewRouteEntry("/stuff", http.MethodGet, s.Get)}
}
