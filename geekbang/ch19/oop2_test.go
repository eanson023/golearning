package ch19

import (
	"testing"

	"github.com/eanson023/golearning/geekbang/ch19/person"
)

// 面向对象三大特性

// 封装：把抽象出的字段和对字段的操作封装在一起,数据被保护在内部,程序的其它包只有通过被授权的操作(方法),才能对字段进行修改，其作用有：

// - 隐藏实现细节
// - 可以对数据进行验证，保证安全合理

// Golang对面向对象做了极大简化，并不强调封装特性，下列示例进行模拟实现：
//在person包下新建person.go文件

//oop2_test.go文件操作person
func TestUsePerson(t *testing.T) {
	p := person.NewPerson("Eanson")
	p.SetAge(18)
	t.Log(p)
}

// 继承

// 在 Golang 中，如果一个 struct 嵌套了另一个匿名结构体，那么这个结构体可以直接访 问匿名结构体的字段和方法，从而实现了继承特性。

type Father struct {
	Name string
	age  int
}

type Son struct {
	Father
}

func TestExtend(t *testing.T) {
	var s Son
	s.Father.Name = "Tom"
	s.Father.age = 10

	//上面简述为
	s.Name = "Tom"
	s.age = 222
}

/*
注意：

- 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。
- 结构体嵌入多个匿名结构体，如果两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，访问时必须明确指定匿名结构体名字，否则编译报错。
- 如果一个 struct 嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字。
*/
