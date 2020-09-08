package crawler

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

/*
爬取我们学校教务系统的信息
*/
const URL = "http://jwgl.sanxiau.edu.cn"

//RequestHeader ... 请求header
type RequestHeader struct {
	cookies   []*http.Cookie
	headerMap map[string]string
}

func newRequestHeader() *RequestHeader {
	return &RequestHeader{
		cookies: []*http.Cookie{},
		headerMap: map[string]string{
			"User-Agent":      " Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36",
			"Accept-Language": " zh-CN,zh;q=0.9,en;q=0.8",
		},
	}
}

//添加cookie进ReqHeader
func (reqHeader *RequestHeader) addCookie(cookies []*http.Cookie) {
	for _, cookie := range cookies {
		reqHeader.cookies = append(reqHeader.cookies, cookie)
	}
}

// 将RequestHeader导入到http请求中
func (reqHeader *RequestHeader) addAllToRequest(request *http.Request) {
	// 添加cookie
	for _, cookie := range reqHeader.cookies {
		request.AddCookie(cookie)
	}
	// 添加requestHeader
	for key, value := range reqHeader.headerMap {
		request.Header.Add(key, value)
	}
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

func (lf *LoginForm) addToRequestForm(form url.Values) {
	getValue := reflect.ValueOf(lf).Elem()
	getType := reflect.TypeOf(lf).Elem()
	// 反射添加到map
	for i := 0; i < getValue.NumField(); i++ {
		form[getType.Field(i).Name] = []string{getValue.Field(i).String()}
	}
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

func makeNilParamRequest(method string, url string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return request
}

//Login ... 登录教务管理系统
// 1. 获取cookie,通过cookie获取该验证码,随后制造表单登录
// 2. 模拟表单登录，随后根据Location跳转获取信息
func Login(username string, password string) {
	requestHeader := newRequestHeader()
	if resp, err := http.Get(URL); err != nil {
		panic(err)
	} else {
		requestHeader.addCookie(resp.Cookies())
	}
	client := new(http.Client)
	checkCodeURL := "http://jwgl.sanxiau.edu.cn/CheckCode.aspx?"
	request := makeNilParamRequest("GET", checkCodeURL)
	// 将所有请求头加到请求对象里面
	requestHeader.addAllToRequest(request)
	// 执行
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
		var checkcode string
		for checkcode == "" {
			fmt.Print("请输入验证码(回车结束):")
			fmt.Scanln(&checkcode)
		}
		loginURL := "http://jwgl.sanxiau.edu.cn/default2.aspx"
		request := makeNilParamRequest("POST", loginURL)
		// 将cookie和请求头加到请求中
		requestHeader.addAllToRequest(request)
		loginForm := newLoginForm(username, password, checkcode)
		// 竟然要自己创建
		request.PostForm = make(url.Values)
		// 添加
		loginForm.addToRequestForm(request.PostForm)
		resp, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		// 登录后重定向url
		if location, ok := resp.Header["Location"]; !ok {
			printReader(resp.Body)
		} else {
			redirectURL := URL + location[0]
			fmt.Printf("跳转到:%v\n", redirectURL)
			request = makeNilParamRequest("GET", redirectURL)
			requestHeader.addAllToRequest(request)
			if resp, err = client.Do(request); err != nil {
				panic(err)
			} else {
				printReader(resp.Body)
			}
		}
	}
}

func printReader(reader io.ReadCloser) {
	defer reader.Close()
	data, _ := ioutil.ReadAll(reader)
	enc := mahonia.NewDecoder("gbk")
	str := enc.ConvertString(string(data))
	fmt.Println(str)
}
