// fetchall 并发获取URL并报告它们的时间和大小
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 启动一个goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 从通道ch接收
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 发送到通道ch
		ch <- fmt.Sprint(err)
		return
	}
	// 写入输出流进行丢弃，所以只记字节数
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// 不要泄露资源
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
