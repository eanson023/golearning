package ch23

import (
	"reflect"
	"testing"
)

// 反射的使用
// 反射操作简单数据类型

func Test1(t *testing.T) {
	var num int64 = 100
	//设置值:指针传递
	ptrValue := reflect.ValueOf(&num)
	// Elem()用于获取原始值的反射对象
	newValue := ptrValue.Elem()
	t.Log("type:", newValue.Type())       //int64
	t.Log(" can set:", newValue.CanSet()) //true
	newValue.SetInt(20)

	// 获取值:值传递
	rValue := reflect.ValueOf(num)
	t.Log(rValue.Int())               //方式1
	t.Log(rValue.Interface().(int64)) //方式2
}
