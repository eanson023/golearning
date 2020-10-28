package data

import (
	"fmt"
	"github.com/go-playground/validator"
	"regexp"
)

// Validation contain
type Validation struct {
	validate *validator.Validate
}

func NewValiation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", skuValidation)
	return &Validation{validate}
}

// 匿名变量 这样做不会将validator暴露出去
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors 错误集合
type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// 参数校验
// Validate the item
// for more detail the returned error can be cast into a
// validator.ValidationErrors collection
func (v *Validation) Validate(i interface{}) ValidationErrors {
	// 有点迷幻
	errs := v.validate.Struct(i).(validator.ValidationErrors)
	if len(errs) == 0 {
		return nil
	}

	var returnErrs ValidationErrors
	for _, err := range errs {
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}

// 自定义validation  利用正则表达式
func skuValidation(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`^[a-z]+-[a-z]+-[a-z]+$`)
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
}
