package strategy

import (
	"fmt"
	"testing"
)

func TestRubberDuck(t *testing.T) {
	// 类无法实现多态 5555~
	var rubberDuck = NewRubberDuck()
	rubberDuck.display()
	rubberDuck.performFly()
	rubberDuck.performQuack()
}

func TestRedHeadDuck(t *testing.T) {
	var redHeadDuck = NewRedHeadDuck()
	redHeadDuck.display()
	redHeadDuck.performFly()
	redHeadDuck.performQuack()
}

// 使用者想自己实现个策略
type FlyWithRocket struct {
}

func (fwr *FlyWithRocket) fly() {
	fmt.Println("我使用火箭飞~")
}
func TestGreenRocketDuck(t *testing.T) {
	// 自定义鸭子
	greenRocketDuck := NewDuck("我系绿头鸭", &FlyWithRocket{}, &Quack{})
	greenRocketDuck.display()
	greenRocketDuck.performFly()
	greenRocketDuck.performQuack()
	// 动态更改行为
	greenRocketDuck.SetFlyBehavior(&FlyWithWings{})
	greenRocketDuck.performFly()
}
