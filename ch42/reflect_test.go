package ch42

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
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
	// 反射添加进条件中
	params := make([]string, 4)
	for i := 0; i < getValue.NumField(); i++ {
		str := getType.Field(i).Name + "=" + getValue.Field(i).String()
		params = append(params, str)
		// _ = writer.WriteField(getType.Field(i).Name, getValue.Field(i).String())
	}
	ret := url.QueryEscape(strings.Join(params, "&"))
	fmt.Println(ret)
}

func Test(t *testing.T) {
	lf := newLoginForm("1", "2", "3")
	var form url.Values = make(url.Values)
	lf.addToRequestForm(form)
	for key, value := range form {
		t.Log(key, value)
	}
}

func TestUrlEncode(t *testing.T) {
	str := "\u4e2d\u6587"
	t.Log(url.QueryEscape(str))
	sText := "中文"
	textQuoted := strconv.QuoteToASCII(sText)
	t.Log(textQuoted)
}
