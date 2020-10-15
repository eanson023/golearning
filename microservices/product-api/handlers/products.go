package handlers

import (
	"encoding/json"
	"github.com/eanson023/golearning/microservices/product-api/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	rw.Write(d)
}
