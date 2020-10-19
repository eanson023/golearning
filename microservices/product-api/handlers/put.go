package handlers

import (
	"fmt"
	"github.com/eanson023/golearning/microservices/product-api/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  400: errorResponse
//  404: errorResponse
//  422: errorValidation

// UpdateProduct 更新产品
func (p *Products) UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, fmt.Sprintf("can't format the var:%s we need number", vars["id"]), http.StatusBadRequest)
		return
	}
	p.l.Println("Hanlde PUT Product")
	// 从上下文中获取
	prod := req.Context().Value(ContextKeyProduct{}).(*data.Product)

	prod.ID = id
	p.l.Printf("Prod: %#v\n", prod)
	err = data.UpdateProduct(prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
	}
}
