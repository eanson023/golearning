package ch24

import (
	"errors"
	"fmt"
	"testing"
)

/*
panic用于不可以恢复的错误
panic退出前会执行defer指定的内容
*/

func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
	}()
	fmt.Println("Start")
	// os.Exit(-1)
	//空接口
	panic(errors.New("Something wrong"))
	fmt.Println("Exit")
}

/*
recover
*/

func TestRecover(t *testing.T) {
	defer func() {
		// 程序恢复 相当于java catch
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()
	fmt.Println("Start")
	//空接口
	panic(errors.New("Something wrong"))
}
