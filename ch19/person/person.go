package person

import (
	"fmt"
)

type person struct {
	Name string
	age  int //年龄是隐私，不允许其他包访问
}

//NewPerson 工厂函数(类似构造函数)
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄不合法")
	}
}

func (p *person) GetAge() int {
	return p.age
}
