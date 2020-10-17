package handlers

import (
	"regexp"
	"strconv"
	// "encoding/json"
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

// 实现server.go http.Hanlder接口 构建RESTful服务
func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		p.getProducts(rw, req)
	case http.MethodPost:
		p.addProduct(rw, req)
	case http.MethodPut:
		// 使用正则判断uri数字
		r := regexp.MustCompile(`/([0-9]+)`)
		g := r.FindAllStringSubmatch(req.URL.Path, -1)

		if len(g) != 1 || len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		p.l.Println("got id", id)
		p.updateProduct(id, rw, req)
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Hanlde GET Products")
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	// 封装一些实现细节
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Hanlde POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	p.l.Printf("Prod: %#v\n", prod)
	data.AddProduct(prod)
}

// 更新产品 我们希望能从URI中得到ID
func (p *Products) updateProduct(id int, rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Hanlde PUT Product")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	prod.ID = id
	p.l.Printf("Prod: %#v\n", prod)
	err = data.UpdateProduct(prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
	}
}
