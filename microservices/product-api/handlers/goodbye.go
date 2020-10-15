package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{
		l: l,
	}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	g.l.Println("goodbye")
	fmt.Fprint(rw, "goodbye")
	for {

	}
}
