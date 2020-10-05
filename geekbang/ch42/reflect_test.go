package ch42

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
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
type scoreForm struct {
	__VIEWSTATE string
	// 学年
	ddlXN string
	// 学期
	ddlXQ string
	// 安学期查询
	Button1 string
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
func addToRequestForm(form interface{}) {
	getValue := reflect.ValueOf(form).Elem()
	getType := reflect.TypeOf(form).Elem()
	// 反射添加到map
	// 反射添加进条件中
	params := make([]string, 4)
	for i := 0; i < getValue.NumField(); i++ {
		str := getType.Field(i).Name + "=" + getValue.Field(i).String()
		params = append(params, str)
	}
	ret := url.QueryEscape(strings.Join(params, "&"))
	fmt.Println(ret)
}

func Test(t *testing.T) {
	var scoreForm *scoreForm = &scoreForm{
		ddlXN: "1",
	}
	addToRequestForm(scoreForm)
}

func TestUrlEncode(t *testing.T) {
	str := "\u4e2d\u6587"
	t.Log(url.QueryEscape(str))
	sText := "中文"
	textQuoted := strconv.QuoteToASCII(sText)
	t.Log(textQuoted)
}

func TestTTTT(t *testing.T) {
	str := "大家好：12345"
	t.Log(strings.Split(str, "：")[1])
}

func TestGoquery(t *testing.T) {
	html := `<body>

				<div>DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	dom.Find("div[class=name]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	t.Log(dom.Find("div[class=name]").Text())
}
