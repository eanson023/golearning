// 练习1.10 在一个产生大量数据的网站。连续两次运行fetchall,看报告时间是否会有大的变化，检查缓存情况。每一次获取的内容一样吗？
// 修改fetchall将内容输出到文件，这样可以检查它是否一致
package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
	f, _ := os.Create(strconv.Itoa(rand.Intn(1000)) + ".txt")
	// 写入输出流进行丢弃，所以只记字节数
	nbytes, err := io.Copy(f, resp.Body)
	// 不要泄露资源
	defer func() {
		resp.Body.Close()
		f.Close()
	}()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
