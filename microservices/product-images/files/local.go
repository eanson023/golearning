package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Local 本地磁盘存储
type Local struct {
	maxFileSize int //文件最大字节数
	basePath    string
}

// NewLocal creates a new Local filesytem with the given base path
// basePath is the base directory to save files to
// maxSize is the max number of bytes that a file can be
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{maxSize, p}, nil
}

// Save the contents of the Writer to the given path
// path is a relative path, basePath will be appended
func (l *Local) Save(path string, contents io.Reader) error {
	fullPath := l.fullPath(path)
	// 看文件夹是否存在
	dirPath := filepath.Dir(fullPath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to create directory: %v", err)
	}

	// 获取文件信息 存在着删除它
	_, err = os.Stat(fullPath)
	// 逻辑:如果错误为空则代表能检测到该文件
	if err == nil {
		err = os.Remove(fullPath)
		if err != nil {
			return fmt.Errorf("Unable to delete file:%v", err)
		}
	} else if !os.IsNotExist(err) {
		// 如果是除了文件不存在的错误
		return fmt.Errorf("Unable to get file:%v", err)
	}

	// 创建新文件
	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("Unable to create file: %v", err)
	}
	defer f.Close()

	_, err = io.Copy(f, contents)
	if err != nil {
		return fmt.Errorf("Unable to open file:%v", err)
	}
	// 终于上传成功了
	return nil
}

// 返回绝对路径
func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
