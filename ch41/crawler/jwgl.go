package crawler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

/*
爬取我们学校教务系统的信息
*/
const URL = "http://jwgl.sanxiau.edu.cn"

// 发送请求的客户端
var client *http.Client = new(http.Client)

//RequestGeneral ...
type RequestGeneral struct {
	cookies []*http.Cookie
	headers map[string]string
}

//LoginForm 登录表单
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

func newRequestGeneral() *RequestGeneral {
	return &RequestGeneral{
		cookies: []*http.Cookie{},
		headers: map[string]string{
			"User-Agent":      " Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36",
			"Accept-Language": " zh-CN,zh;q=0.9,en;q=0.8",
		},
	}
}

//添加cookie进ReqHeader
func (reqGeneral *RequestGeneral) addCookie(cookies []*http.Cookie) {
	for _, cookie := range cookies {
		cookie.MaxAge = 1800
		reqGeneral.cookies = append(reqGeneral.cookies, cookie)
	}
}
func (reqGeneral *RequestGeneral) addHeader(key string, value string) {
	reqGeneral.headers[key] = value
}

// 将RequestHeader导入到http请求中
func (reqGeneral *RequestGeneral) headerIntoRequest(request *http.Request) {
	// 添加requestHeader
	for key, value := range reqGeneral.headers {
		request.Header.Add(key, value)
	}
}

//将cookie加到请求中
func (reqGeneral *RequestGeneral) cookieIntoRequest(request *http.Request) {
	// 添加cookie
	for _, cookie := range reqGeneral.cookies {
		request.AddCookie(cookie)
	}
}

// fmt.Printf("", var)
func (reqGeneral *RequestGeneral) newEmptyBodyRequest(method string, url string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	// 将基本请求头加入到请求中
	reqGeneral.headerIntoRequest(request)
	return request
}

func (reqGeneral *RequestGeneral) newRequest(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	// 将基本请求头加入到请求中
	reqGeneral.headerIntoRequest(request)
	return request
}

func (lf *LoginForm) generateForm(reqGeneral *RequestGeneral) *bytes.Buffer {
	// x-www-form-urlencoded
	payload := &bytes.Buffer{}
	getValue := reflect.ValueOf(lf).Elem()
	getType := reflect.TypeOf(lf).Elem()
	// 反射添加进条件中
	params := []string{}
	for i := 0; i < getValue.NumField(); i++ {
		str := getType.Field(i).Name + "=" + url.QueryEscape(getValue.Field(i).String())
		params = append(params, str)
	}
	ret := strings.Join(params, "&")
	fmt.Println(ret)
	payload.ReadFrom(strings.NewReader(ret))
	reqGeneral.addHeader("Content-Type", "application/x-www-form-urlencoded")
	return payload
}

func newLoginForm(username string, password string, checkcode string) *LoginForm {
	return &LoginForm{
		// 固定值
		__VIEWSTATE: "dDw3OTkxMjIwNTU7Oz6bmpbeSO1k01TBeZU9nxNbmYM4aw==",
		TextBox1:    username,
		TextBox2:    password,
		TextBox3:    checkcode,
		// 学生
		RadioButtonList1: "学生",
	}
}

//Start ... 程序入口
func Start(username string, password string) {
	reqGeneral := newRequestGeneral()
	cookies, err := getJwglCookies(reqGeneral)
	if err != nil {
		panic(err)
	}
	// 添加cookie
	reqGeneral.addCookie(cookies)
	// 随便
	time.Sleep(time.Second * 1)
	checkcode := inputCheckCode(reqGeneral)
	// 制造表单
	loginForm := newLoginForm(username, password, checkcode)
	// 再次睡眠
	time.Sleep(time.Second * 1)
	success, _ := login(reqGeneral, loginForm)
	if success {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}
}

// 获取教务系统cookie 里面包含sessionID
func getJwglCookies(reqGeneral *RequestGeneral) ([]*http.Cookie, error) {
	request := reqGeneral.newEmptyBodyRequest("GET", URL)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp.Cookies(), nil
}

// 1. 获取cookie,通过cookie获取该验证码
func inputCheckCode(reqGeneral *RequestGeneral) string {
	var checkcode string
	checkCodeURL := "http://jwgl.sanxiau.edu.cn/CheckCode.aspx?"
	request := reqGeneral.newEmptyBodyRequest("GET", checkCodeURL)
	reqGeneral.cookieIntoRequest(request)
	reqGeneral.headerIntoRequest(request)
	if resp, err := client.Do(request); err != nil {
		panic(err)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fileName := "checkcode.gif"
		err = ioutil.WriteFile(fileName, body, os.ModePerm)
		if err != nil {
			fmt.Println("写入验证码文件出错")
			panic(err)
		}
		fmt.Println("请查看文件夹目录里的" + fileName + "文件")
		for checkcode == "" {
			fmt.Print("请输入验证码(回车结束):")
			fmt.Scanln(&checkcode)
		}
	}
	return checkcode
}

// 2. 模拟表单登录，随后根据Location跳转获取信息
// return 是否登录成功和响应结果
func login(reqGeneral *RequestGeneral, loginForm *LoginForm) (bool, *http.Response) {
	loginURL := "http://jwgl.sanxiau.edu.cn/default2.aspx"
	// 添加
	param := loginForm.generateForm(reqGeneral)
	request := reqGeneral.newRequest("POST", loginURL, param)
	reqGeneral.cookieIntoRequest(request)
	reqGeneral.headerIntoRequest(request)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	// 用goquery解析html 判断是否登录成功
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	scriptText := doc.Find("script").Last().Text()
	fmt.Println(scriptText)
	wrong := strings.Contains(scriptText, "验证码不正确")
	wrong = wrong || strings.Contains(scriptText, "密码错误")
	return !wrong, resp
}
