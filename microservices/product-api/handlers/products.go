package handlers

import (
	"context"
	"fmt"
	"github.com/eanson023/golearning/microservices/product-api/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts 获取产品
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Hanlde GET Products")
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	// 封装一些实现细节
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct 添加产品
func (p *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Hanlde POST Product")

	// 从上下文中获取
	prod := req.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v\n", prod)
	data.AddProduct(prod)
}

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
	prod := req.Context().Value(KeyProduct{}).(*data.Product)

	prod.ID = id
	p.l.Printf("Prod: %#v\n", prod)
	err = data.UpdateProduct(prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
	}
}

type KeyProduct struct{}

// MidllewareProductValidation mux包中type MiddlewareFunc func(http.Handler) http.Handler方法
// 定义中间件 来验证json数据 他会在实际的hanlder执行之前执行
func (p *Products) MidllewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("Unable to unmarshal json")
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate the product
		if err := prod.Validate(); err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, "Error valiating product "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		// 将数据用上下文的防暑放到请求结构体中
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
