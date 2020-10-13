package strategy

import (
	"fmt"
)

// Display 显示接口 所有继承Duck类的可以实现这个接口
type Display interface {
	display()
}

// Duck 普通鸭子父类
type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
	description   string
}

// SetFlyBehavior 设置飞行行为
func (d *Duck) SetFlyBehavior(flyBehavior FlyBehavior) {
	d.flyBehavior = flyBehavior
}

// 委托
func (d *Duck) performQuack() {
	d.quackBehavior.quack()
}

// 委托
func (d *Duck) performFly() {
	d.flyBehavior.fly()
}

// 实现Display接口
func (d *Duck) display() {
	fmt.Println(d.description)
}

// NewDuck 根据策略创建鸭子
func NewDuck(desc string, flyBehavior FlyBehavior, quackBehavior QuackBehavior) *Duck {
	return &Duck{
		description:   desc,
		flyBehavior:   flyBehavior,
		quackBehavior: quackBehavior,
	}
}

// NewRubberDuck 创建一个橡皮鸭
func NewRubberDuck() *Duck {
	return &Duck{
		description:   "我系橡皮鸭",
		flyBehavior:   &FlyNoWay{},
		quackBehavior: &Squeak{},
	}
}

// NewRedHeadDuck 创建一个红头鸭
func NewRedHeadDuck() *Duck {
	return &Duck{
		description:   "我系红头鸭",
		flyBehavior:   &FlyWithWings{},
		quackBehavior: &Quack{},
	}
}
