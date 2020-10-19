package handlers

import (
	"github.com/eanson023/golearning/microservices/product-api/data"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	201: noContentResponse
//  400: errorResponse
//  422: errorValidation

// AddProduct 添加产品
func (p *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Hanlde POST Product")

	// 从上下文中获取
	prod := req.Context().Value(ContextKeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v\n", prod)
	data.AddProduct(prod)
	rw.WriteHeader(http.StatusCreated)
}
