package crawler

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"net/http"
	"strings"
)

//为了更好的爬取，写一个格式化输出语句

type Charset byte

const (
	UTF8 Charset = 0 + iota
	GBK
)

func PrintResponseInfo(response *http.Response, charset Charset) {
	//HTTP/1.1 302 Found
	fmt.Printf("%s %s %s\r\n", response.Proto, response.StatusCode, response.Status)
	printHeader(response.Header)
	printlnCookie(response.Cookies())
	if charset == UTF8 {
		printBody(response.Body)
	} else if charset == GBK {
		printBodyGBK(response.Body)
	}
	fmt.Println("*******************************************")
}

func PrintRequestInfo(request *http.Request, charset Charset) {
	// POST /default2.aspx HTTP/1.1
	fmt.Printf("%s %s %s\r\n", request.Method, request.RequestURI, request.Proto)
	printHeader(request.Header)
	printlnCookie(request.Cookies())
	if charset == UTF8 {
		printBody(request.Body)
	} else if charset == GBK {
		printBodyGBK(request.Body)
	}
	fmt.Println("*******************************************")
}

func printHeader(header http.Header) {
	for key, values := range header {
		fmt.Printf("%s: %s\r\n", key, strings.Join(values, ","))
	}
}

func printlnCookie(cookies []*http.Cookie) {
	for _, cookie := range cookies {
		fmt.Printf("%s: %s\r\n", cookie.Name, cookie.String())
	}
}

func printBody(body io.ReadCloser) {
	sb := readBody(body)
	if sb == nil {
		return
	}
	fmt.Printf("\r\n%s\r\n", sb.String())
}

func printBodyGBK(body io.ReadCloser) {
	sb := readBody(body)
	if sb == nil {
		return
	}
	dec := mahonia.NewDecoder("gbk")
	str := dec.ConvertString(sb.String())
	fmt.Printf("\r\n%s\r\n", str)
}

func readBody(body io.ReadCloser) *strings.Builder {
	if body == nil {
		return nil
	}
	defer body.Close()
	var sb strings.Builder
	buf := make([]byte, 1024)
	for {
		n, err := body.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		sb.Write(buf[:n])
		if err == io.EOF {
			break
		}
	}
	return &sb
}
