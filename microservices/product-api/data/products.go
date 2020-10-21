package data

import (
	"fmt"
	"time"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the user
	//
	// required:true
	// min: 1
	ID int `json:"id"`

	// the name for this product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`
	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`
	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`
	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU       string `json:"sku" validate:"required,sku"` //自定义验证函数
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"` //使用'-'忽略该字段
}

// Products 产品列表
type Products []*Product

var producList Products = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Freonthy milk coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// GetProducts 模仿数据库 获取数据
func GetProducts() Products {
	return producList
}

func GetProduct(p *Product) error {
	pos, err := findProduct(p.ID)
	if err != nil {
		return err
	}
	// 更改指针指向
	*p = *producList[pos]
	return nil
}

// AddProduct 添加商品
func AddProduct(p *Product) {
	p.ID = getNextID()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	producList = append(producList, p)
}

// UpdateProduct 更新商品
func UpdateProduct(p *Product) error {
	pos, err := findProduct(p.ID)
	if err != nil {
		return err
	}
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	producList[pos] = p
	return nil
}

func DeleteProduct(p *Product) error {
	pos, err := findProduct(p.ID)
	if err != nil {
		return err
	}
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	producList = append(producList[:pos], producList[pos+1:]...)
	return nil
}

// ErrorProductNotFound 商品没找到
var ErrorProductNotFound = fmt.Errorf("Product not found")

// return productList index
func findProduct(id int) (int, error) {
	for k, p := range producList {
		if p.ID == id {
			return k, nil
		}
	}
	return -1, ErrorProductNotFound
}

func getNextID() int {
	lp := producList[len(producList)-1]
	return lp.ID + 1
}
