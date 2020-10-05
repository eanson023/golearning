package ch20

import (
	"fmt"
	"testing"
)

type BreathAble interface {
	Breath() //动物都具备 呼吸方法
}

type Flyer interface {
	Fly()
}

type Swimer interface {
	Swim()
}

//定义一个鸟类：其呼吸的方式四在陆地
type Bird struct {
	Name string
	Food string
	Kind string
}

func (b *Bird) Breath() {
	fmt.Println("鸟在呼吸")
}

func (b *Bird) Fly() {
	fmt.Printf("%s 在 飞\n", b.Name)
}

// 一定一个鱼类：其呼吸方式是在水下
type Fish struct {
	Name string
	Kind string
}

func (f *Fish) Breath() {
	fmt.Println("鱼 在 水下 呼吸")
}
func (f *Fish) Swim() {
	fmt.Printf("%s 在游泳\n", f.Name)
}

//断言

//一个普通函数，参数是动物接口
func Display(a BreathAble) {
	//调用接口的方法
	a.Breath()
	//调用实现类的成员:此时会报错
	// fmt.Println(a.Name)

	// 接口类型无法直接访问其具体实现类的成员，需要使用断言（type assertions），对接口的类型进行判断，类型断言格式：
	/*
		t := i.(T)			//不安全写法：如果i没有完全实现T接口的方法，这个语句将会触发宕机
		t, ok := i.(T)		// 安全写法：如果接口未实现接口，将会把ok掷为false，t掷为T类型的0值
		- i代表接口变量
		- T代表转换的目标类型
		- t代表转换后的变量
	*/
	//调用实现类的成员，此时会报错
	instance, ok := a.(*Bird) //注意：这里必须是 *Bird类型，因为*Bird实现了接口，不是Bird实现类接口

	if ok {
		// 得到具体的实现类，才能访问类的成员
		fmt.Println("该鸟类的名字是：", instance.Name)
	} else {
		fmt.Println("该动物不是鸟类")
	}
}

func TestInterface4(t *testing.T) {
	var b = &Bird{
		"斑鸠",
		"蚂蚱",
		"鸟类",
	}
	Display(b)
}
