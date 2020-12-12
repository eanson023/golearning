package client

import (
	// 从GOPATH src后的目录开始

	"testing"

	"github.com/eanson023/golearning/geekbang/ch25/series"
)

/*
package
1. 基本复用模块单元
	以首字母大写来表示可被包外代码访问
2. 代码的package可以和所在的目录不一致
3. 同一目录里的Go代码的package要保持一致
*/

/*
1. 通过go get 来获取远程依赖
 go get -u强制从网络更新远程依赖

 2.注意代码在Github上的组织形式，以适应go get
  直接以代码路径开始，不要有src
*/

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
}
