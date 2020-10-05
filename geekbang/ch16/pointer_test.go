package ch16

import "testing"

// ### 1.1 指针的创建
// Go保留了指针，代表某个内存地址，默认值为 `nil` ，使用 `&` 取变量地址，通过 `*` 访问目标对象。
func TestPointerCreate1(t *testing.T) {
	var a int = 10
	// 0xc00001a270 一个十六进制数
	t.Log("&a=", &a)

	var p *int = &a
	//10
	t.Log("*p=", *p)
}

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

/*
注意：
- Go同样支持多级指针，如 **T
- 空指针：声明但未初始化的指针
- 野指针：引用了无效地址的指针，如：`var p *int = 0`，`var p *int = 0xff00`(超出范围)
- Go中直接使用` . `访问目标成员
*/

type User struct {
	name string
	age  int
}

func TestStructPointer(t *testing.T) {
	var u = User{
		name: "lisi",
	}
	p := &u
	t.Log(u.name)
	//语法糖 c语言中 结构体指针类型用-> 访问属性 这里用.
	t.Log(p.name)
}

// ### 1.4 Go不支持指针运算
// 由于垃圾回收机制的存在，指针运算造成许多困扰，所以Go直接禁止了指针运算
func TestPointer2(t *testing.T) {
	a := 1
	p := &a
	t.Log(p)
	// p++ //报错 non-numeric type *int

	// new()函数使用
	// new()函数可以在 heap堆 区申请一片内存地址空间：
	var bp *bool
	bp = new(bool)
	t.Log(*bp) //false
}
