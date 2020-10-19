package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator" //验证
	"io"
	"regexp"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"` //自定义验证函数
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"` //使用'-'忽略该字段
}

var producList = []*Product{
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

// 参数校验
func (p *Product) Validate() error {
	validate := validator.New()
	// 注册validation
	validate.RegisterValidation("sku", skuValidation)
	return validate.Struct(p)
}

// 自定义validation  利用正则表达式
func skuValidation(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
}

// FromJSON deserialize
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

type Products []*Product

// ToJSON 封装json数据 将其写入到writer中
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts 模仿数据库 获取数据
func GetProducts() Products {
	return producList
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
