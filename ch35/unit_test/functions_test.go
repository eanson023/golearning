package unit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	// 表格测试法
	inputs := [...]int{1, 2, 3}
	excepted := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != excepted[i] {
			t.Errorf("input is %d, the excepted is %d, the actual %d", inputs[i], excepted[i], ret)
		}
	}
}

func TestErrorIncode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal()
	// 直接终止 不输出了
	fmt.Println("End")
}

func TestSquareWithAssert(t *testing.T) {
	assert := assert.New(t)
	// 表格测试法
	inputs := [...]int{1, 2, 3}
	excepted := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(excepted[i], ret)
	}
}
