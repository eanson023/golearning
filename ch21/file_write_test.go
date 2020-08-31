package ch21

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

//文件的基本操作
// 创建文件
/*
使用Create()创建文件时：

- 如果文件不存在，则创建文件。
- 如果文件存在，则清空文件内内容。
- Create创建的文件任何人都可以读写。
*/
func TestCreateFile(t *testing.T) {
	f, err := os.Create("test.txt")
	if err != nil {
		t.Log(err)
		return
	}
	//打印文件指针
	t.Log(f)
	//打开的资源在不需要是必须关闭
	f.Close()
}

// 打开文件，写入内容
/*
打开文件有两种方式：

- Open()：以只读的方式打开文件，若文件不存在则会打开失败
- OpenFile()：打开文件时，可以传入打开方式，该函数的三个参数：
  - 参数1：要打开的文件路径
  - 参数2：文件打开模式，如 `O_RDONLY`，`O_WRONGLY`，`O_RDWR`，还可以通过管道符来指定文件不存在时创建文件
  - 参数3：文件创建时候的权限级别，在0-7之间，常用参数为6
*/
func TestOpen(t *testing.T) {
	_, err := os.Open("null.txt")
	if err != nil {
		t.Log(err)
		return
	}
}

func TestOpenFile(t *testing.T) {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		t.Log(err)
		return
	}
	// 写文件 n:代表个数
	n, err := f.WriteString("你还好吗")
	if err != nil {
		t.Log("write err:", err)
		return
	}
	t.Log("write number = ", n)
	f.Close()
}

/*
修改文件的读写指针位置 `Seek()`，包含两个参数：

- 参数1：偏移量，为正数时向文件尾偏移，为负数时向文件头偏移
- 参数2：偏移的开始位置，包括：
  - io.SeekStart：从文件起始位置开始
  - io.SeekCurrent：从文件当前位置开始
  - io.SeekEnd：从文件末尾位置开始

`Seek()`函数返回
*/
func TestSeek(t *testing.T) {
	f, _ := os.OpenFile("test.txt", os.O_RDWR, 6)
	offset, _ := f.Seek(5, io.SeekStart)
	t.Log("offset:", offset)
	n, _ := f.WriteAt([]byte("阿U欧克？"), offset)
	t.Log(n)
	f.Close()
}

// 获取文件描述信息 os.Stat()
// Go的os包中定义了file类，封装了文件描述信息，同时也提供了Read、Write的实现。
func TestStat(t *testing.T) {
	fileInfo, err := os.Stat("./test.txt")
	if err != nil {
		t.Log(err)
		return
	}
	// *os.fileStat
	t.Logf("%T\n", fileInfo)
	t.Log(fileInfo.Name())
}

// 路径、目录操作
func TestPathOrDir(t *testing.T) {
	//路径操作
	t.Log(filepath.IsAbs("test.txt")) //false:判断是否为绝对值
	t.Log(filepath.Abs("test.txt"))   //转换为绝对值

	//创建目录
	err := os.Mkdir("./test", os.ModePerm)
	if err != nil {
		t.Log(err)
	}

	// 创建多级目录 mkdir
	err = os.MkdirAll("test/all/a/b/c", os.ModePerm)
	if err != nil {
		t.Log(err)
	}
}

// 贴士：Openfile()可以用于打开目录。

func Test1(t *testing.T) {
	pathPreffix := "../ch"
	var path string
	for i := 23; i < 120; i++ {
		path = pathPreffix + strconv.Itoa(i)
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			t.Log(err)
			return
		}
		abs, _ := filepath.Abs(path)
		t.Log("绝对路径:", abs)
	}
}

//删除文件
// 该函数也可用于删除目录（只能删除空目录）。如果要删除非空目录，需要使用 `RemoveAll()` 函数
func TestDelete(t *testing.T) {
	err := os.Remove("test.txt")
	if err != nil {
		t.Log("remove err:", err)
		return
	}
}
