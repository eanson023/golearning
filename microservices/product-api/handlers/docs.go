// Package handlers classification of Product API
//
// Documentation for Product API
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"github.com/eanson023/golearning/microservices/product-api/data"
)

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of products return in response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	// required: true
	Body []data.Product
}

// noContent to response
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// 之后的tag是指明引用处
// swagger:parameters listSingleProduct deleteProduct
type productIdParamsWarpper struct {
	// The id of the product for which the operation relates
	// in:path
	// required:true
	ID int `json:"id"`
}
