package ch21

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// 直接读取read()
// read() 实现的是按字节数读取
func TestRead(t *testing.T) {
	f, _ := os.Open("test.txt")
	// 指定要读取的长度
	readByte := make([]byte, 128)
	for {
		//将数据读取如切片，返回值 n 是实际读取到的字节数
		n, err := f.Read(readByte)
		//如果读到了文件末尾:EOF即 end of file
		if err != nil && err != io.EOF {
			t.Log("read file:", err)
			break
		}
		t.Log("read: \n", string(readByte))
		t.Log("--------------------------------")
		if n < 128 {
			t.Log("read end")
			break
		}
	}
}

// bufio的写操作
/*
bufio封装了io.Reader、io.Writer接口对象，并创建了另一个也实现了该接口的对象：bufio.Reader、bufio.Writer。通过该实现，bufio实现了文件的缓冲区设计，可以大大提高文件I/O的效率。

使用bufio读取文件时，先将数据读入内存的缓冲区（缓冲区一般比要比程序中设置的文件接收对象要大），这样就可以有效降低直接I/O的次数。

`bufio.Read([]byte)`相当于读取大小`len(p)`的内容：

- 当缓冲区有内容时，将缓冲区内容全部填入p并清空缓冲区
- 当缓冲区没有内容且`len(p)>len(buf)`，即要读取的内容比缓冲区还要大，直接去文件读取即可
- 当缓冲区没有内容且`len(p)<len(buf)`，即要读取的内容比缓冲区小，读取文件内容并填满缓冲区，并将p填满
- 以后再次读取时，缓冲区有内容，将缓冲区内容全部填入p并清空缓冲区（和第一步一致）
*/
func TestBufio(t *testing.T) {
	//创建读对象
	f, _ := os.Open("test.txt")
	//创建读对象
	reader := bufio.NewReader(f)

	//读一行数据
	// ReadBytes读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的切片。如果ReadBytes方法在读取到delim之前遇到了错误，
	//它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。当且仅当ReadBytes方法返回的切片不以delim结尾时，会返回一个非nil的错误。
	byt, _ := reader.ReadBytes('\n')
	t.Log(string(byt))
	// ReadString() 函数也具有同样的功能，且能直接读取到字符串数据，无需转换，示例：读取大文件的全部数据
	// 按照缓冲区读取：读取到特定字符结束
	for {
		//按行读取 读取出来的直接就是字符
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Log("read err:", err)
			break
		}
		t.Log("str = ", str)
		if err == io.EOF {
			t.Log("read end")
			break
		}
	}
	// 在Unix设计思想中，一切皆文件，命令行输入也可以作为文件读入：
	reader = bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('-') //假设命令已 - 开始
	// 缓冲的思想：通过bufio，数据被写入用户缓冲，再进入系统缓冲，最后由操作系统将系统缓冲区的数据写入磁盘。
	t.Log(s)
}

// io/ioutil 包文件读取
func TestIoUtil(t *testing.T) {
	ret, err := ioutil.ReadFile("test.txt")
	if err != nil {
		t.Log("read err:", err)
		return
	}
	t.Log(string(ret))
}

//文件写操作

// 直接写
func TestFileWrite(t *testing.T) {
	f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Log(err)
		return
	}
	defer f.Close()

	n, err := f.Write([]byte("hello world"))
	if err != nil {
		t.Log("write err", err)
	}
	t.Log(n)
}

//bufio的写操作
func TestBufWrite(t *testing.T) {
	f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Log(err)
		return
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	_, err = writer.WriteString("hello world")
	if err != nil {
		t.Log("write err:", err)
		return
	}
	// 必须刷新缓冲区：将缓冲区的内容写入文件中。如果不刷新，则只会在内容超出缓冲区大小时写入
	writer.Flush()
}

// io/ioutil 包文件写入
func TestWriteIoUtil(t *testing.T) {
	s := "你好世界"
	//不能追加
	err := ioutil.WriteFile("test.txt", []byte(s), os.ModePerm)
	if err != nil {
		t.Log(err)
	}
}

//文件读取偏移量
func Test2(t *testing.T) {
	f, err := os.OpenFile("test.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f.Close()
	//读取前三个字节
	bs := make([]byte, 3)
	_, err = f.Read(bs)
	if err != nil {
		t.Log("read err:", err)
		return
	}
	t.Log("读取到的数据是:", string(bs))

	//移动光标
	_, err = f.Seek(12, io.SeekStart) //光标从最开始位置，移动12位
	if err != nil {
		t.Log("seek err:", err)
		return
	}
	_, err = f.Read(bs)
	if err != nil {
		t.Log("read err:", err)
		return
	}
	t.Log("读到的数据是：", string(bs))
}

// 通过记录光标的位置，可以实现断点续传：假设已经下载了1KB文件，即本地临时文件存储了1KB，
// 此时断电，重启后通过本地文件大小、Seek()方法获取到上次读取文件的光标位置即可实现继续下载！
