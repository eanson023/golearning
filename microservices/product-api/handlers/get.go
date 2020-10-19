package handlers

import (
	"fmt"
	"github.com/eanson023/golearning/microservices/product-api/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route GET /products listProducts
//  Returns a list of products
//
// responses:
// 	200: productsResponse

// GetProducts 获取所有产品
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	// 封装一些实现细节 将lp数据 写到writer里
	err := data.ToJSON(lp, rw)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingle
//  Returns a  product
//
// responses:
// 	200: productsResponse
//  404: errorResponse
//  500: errorResponse

// GetProductSingle 根据id获取商品
func (p *Products) GetProductSingle(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Hanlde GET Products")
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, fmt.Sprintf("the uri param must be number:%s", idStr), http.StatusBadRequest)
	}
	prod := &data.Product{ID: id}
	err = data.GetProduct(prod)
	if err == data.ErrorProductNotFound {
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{err.Error()}, rw)
		return
	}
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{err.Error()}, rw)
		return
	}
	// 回写数据
	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
