package ch23

import (
	"errors"
	"github.com/stretchr/testify/assert"
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

// 比较map
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	assert.True(t, reflect.DeepEqual(a, b))
}
func TestWanneng(t *testing.T) {
	stu := new(Student)
	settings := map[string]interface{}{"Age": 18, "Name": "eanson"}
	err := fillBySettings(stu, settings)
	if err != nil {
		t.Log(err)
	}
	t.Log(stu)
}

// 万能方法 填充结构体
func fillBySettings(sc interface{}, settings map[string]interface{}) error {
	//   首先判断sc是不是指针和结构体
	if reflect.TypeOf(sc).Kind() != reflect.Ptr {
		return errors.New("the first param shoud be a pointer to the struct type")
	}
	// Elem()回去指针指向的值
	if reflect.TypeOf(sc).Elem().Kind() != reflect.Struct {
		return errors.New("the first param shoud be a pointer to the struct type")
	}
	if settings == nil {
		return errors.New("settings is nil")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		// 判断类型里面有没有field
		if field, ok = (reflect.ValueOf(sc)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		// 类型一致
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(sc)
			// 解指针
			vstr = vstr.Elem()
			// 找到field设置值
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}
