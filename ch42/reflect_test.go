package ch42

import (
	"net/url"
	"reflect"
	"testing"
)

type LoginForm struct {
	// 不知道是啥
	__VIEWSTATE string
	// 学号
	TextBox1 string
	// 密码
	TextBox2 string
	// 验证码
	TextBox3 string
	//
	RadioButtonList1 string
	Button1          string
}

func newLoginForm(username string, password string, checkcode string) *LoginForm {
	return &LoginForm{
		__VIEWSTATE:      "dDw3OTkxMjIwNTU7Oz6bmpbeSO1k01TBeZU9nxNbmYM4aw==",
		TextBox1:         username,
		TextBox2:         password,
		TextBox3:         checkcode,
		RadioButtonList1: "学生",
	}
}
func (lf *LoginForm) addToRequestForm(form url.Values) {
	getValue := reflect.ValueOf(lf).Elem()
	getType := reflect.TypeOf(lf).Elem()
	// 反射添加到map
	for i := 0; i < getValue.NumField(); i++ {
		form[getType.Field(i).Name] = []string{getValue.Field(i).String()}
	}
}

func Test(t *testing.T) {
	lf := newLoginForm("1", "2", "3")
	var form url.Values = make(url.Values)
	lf.addToRequestForm(form)
	for key, value := range form {
		t.Log(key, value)
	}
}
