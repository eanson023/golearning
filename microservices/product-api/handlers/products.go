package handlers

import (
	"github.com/eanson023/golearning/microservices/product-api/data"
	"log"
)

// ContextKeyProduct 上下文key 用于取value用
type ContextKeyProduct struct{}

// Products is a http.Hanlder
type Products struct {
	l *log.Logger
	v *data.Validation
}

func NewProductsHandler(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}
