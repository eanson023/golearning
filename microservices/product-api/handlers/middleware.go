package handlers

import (
	"context"
	"github.com/eanson023/golearning/microservices/product-api/data"
	"net/http"
)

type ValidationError struct {
	Messages []string `json:"messages"`
}

type GenericError struct {
	Message string `json:"message"`
}

// MidllewareProductValidation mux包中type MiddlewareFunc func(http.Handler) http.Handler方法
// 定义中间件 来验证json数据 他会在实际的hanlder执行之前执行
func (p *Products) MidllewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{err.Error()}, rw)
			return
		}

		// validate the product
		if errs := p.v.Validate(prod); len(errs) != 0 {
			p.l.Println("[ERROR] validating product", err)
			// 422 用于校验错误
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{errs.Errors()}, rw)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyProduct{}, prod)
		// 将数据用上下文的防暑放到请求结构体中
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
