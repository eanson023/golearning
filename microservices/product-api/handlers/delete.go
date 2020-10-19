package handlers

import (
	"fmt"
	"github.com/eanson023/golearning/microservices/product-api/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route DELETE /products/{id} products deleteProduct
//
// responses:
// 	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, fmt.Sprintf("can't format the var:%s we need number", idStr), http.StatusBadRequest)
		return
	}
	p.l.Println("[DEBUG] deleting record id", id)

	// 从上下文中获取
	prod := req.Context().Value(ContextKeyProduct{}).(*data.Product)
	prod.ID = id
	p.l.Println("[DEBUG] Prod:", prod)
	err = data.DeleteProduct(prod)
	if err == data.ErrorProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{err.Error()}, rw)
		return
	}
	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
