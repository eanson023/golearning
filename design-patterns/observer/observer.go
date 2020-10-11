package observer

import (
	"fmt"
)

// Observer 所有观察者实现这个接口
type Observer interface {
	update(subject Subject, args ...interface{})
}

// Display 显示器需要显示必须调用实现这个接口
type Display interface {
	display()
}

// GeneralDisplay 标准显示
type GeneralDisplay struct {
	Data *Data
}

// 实现Observer接口
func (g *GeneralDisplay) update(subject Subject, args ...interface{}) {
	if len(args) > 0 {
		g.Data = args[0].(*Data)
		g.display()
	}
}

// 实现Display接口
func (g *GeneralDisplay) display() {
	fmt.Printf("Current conditions: %.2fF degrees and %.2f humidity\n", g.Data.Pressure, g.Data.Humidity)
}
