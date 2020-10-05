package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type filePool struct {
	bufChannel chan *bufio.Reader
}

const (
	poolSize = 5
)

var filePath string

func init() {
	if path, err := filepath.Abs("./index.htm"); err != nil {
		panic(err)
	} else {
		filePath = path
	}
}

// 创建文件池
func NewFilePool() *filePool {
	filePool := &filePool{
		bufChannel: make(chan *bufio.Reader, poolSize),
	}
	for i := 0; i < poolSize; i++ {
		filePool.appendOneBuf()
	}
	return filePool
}

func (pool *filePool) appendOneBuf() {
	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else {
		pool.bufChannel <- bufio.NewReader(file)
	}
}

func main() {
	var pool *filePool = NewFilePool()
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Printf("recive request from %s\n", req.RemoteAddr)
		buf := <-pool.bufChannel
		buf.WriteTo(writer)
		go pool.appendOneBuf()

	})
	http.HandleFunc("/upload", func(writer http.ResponseWriter, req *http.Request) {
		fileName := "./pic.jpg"
		file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
		defer file.Close()
		bufWriter := bufio.NewWriter(file)
		// bufWriter.ReadFrom(req.Body)
		// bufWriter.Flush()
		for {
			buf := make([]byte, 1024)
			_, err := req.Body.Read(buf)
			if err != nil && err != io.EOF {
				os.Remove(fileName)
				writer.Write([]byte("{\"error\":\"上传过程出错\"}"))
			}
			bufWriter.Write(buf)
			bufWriter.Flush()
			if err == io.EOF {
				break
			}
		}
		writer.Write([]byte("{\"msg\":\"ojbk\"}"))
	})
	var port string = ":8080"
	fmt.Println("server is listening on port" + port)
	http.ListenAndServe(port, nil)
}
