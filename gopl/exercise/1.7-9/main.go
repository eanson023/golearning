package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		const protocol = "http://"
		if !strings.HasPrefix(url, protocol) {
			url = fmt.Sprintf("%v%v", protocol, url)
		}
		// 产生一个HTTP请求
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "response:\n%s %s\n", resp.Proto, resp.Status)
		// 将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
