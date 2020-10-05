package ch20

import (
	"fmt"
	"testing"
)

/*
接口（interface）是调用方和实现方均需要遵守的一种约束，约束开发者按照统一的方法命名、参数类型、数量来处理具体业务。
实际上，接口就是一组没有实现的方法声明，到某个自定义类型要使用该方法时，根据具体情况把这些方法实现出来。

接口语法:
type 接口类型名 interface {
	方法名1(参数列表) 返回值列表
	方法名2(参数列表) 返回值列表
	...
}
*/

//运输方式
type Trasnsporter interface {
	BicycleTran()
	CarTran()
}

type Writer interface {
	// 参数列表和返回值列表中的变量名可以被忽略
	Write([]byte) error
}

//驾驶员
type Driver struct {
	Name string
	Age  int
}

//实现运输方式接口
func (d *Driver) BicycleTran() {
	fmt.Println("使用自行车运输")
}

func (d *Driver) CarTran() {
	fmt.Println("使用小汽车运输")
}

//只要实现了Transporter接口的类型都可以作为参数
func trans(t Trasnsporter) {
	t.BicycleTran()
}

func TestInterface(t *testing.T) {
	driver := &Driver{
		"张三",
		27,
	}
	trans(driver)
	driver.CarTran()
}

/*
注意：

- Go语言的接口在命名时，一般会在单词后面添加er，如写操作的接口叫做Writer
- 当方法名首字母大写，且实现的接口首字母也是大写，则该方法可以被接口所在包之外的代码访问
- 方法与接口中的方法签名一致（方法名、参数列表、返回列表都必须一致）
- 参数列表和返回值列表中的变量名可以被忽略，如：type writer interfae{ Write([]byte) error}
- 接口中所有的方法都必须被实现
- 如果编译时发现实现接口的方法签名不一致，则会报错：` does not implement `。
*/

/*
## Go接口的特点

在上述示例中，Go无须像Java那样显式声明实现了哪个接口，即为非侵入式，接口编写者无需知道接口被哪些类型实现，接口实现者只需要知道实现的是什么样子的接口，
但无需指明实现了哪个接口。编译器知道最终编译时使用哪个类型实现哪个接口，或者接口应该由谁来实现。

类型和接口之间有一对多和多对一的关系，即：

- 一个类型可以实现多个接口，接口间是彼此独立的，互相不知道对方的实现
- 多个类型也可以实现相同的接口。
*/

type Service interface {
	Start()
	Log(string)
}

//日志器
type Logger struct {
}

//日志输出方法
func (g *Logger) Log(s string) {
	fmt.Println("日志:", s)
}

//游戏服务
type GameService struct {
	Logger
}

//实现游戏服务的Start方法
func (g *GameService) Start() {
	fmt.Println("游戏服务启动")
}

// 即使没有接口也能运行，但是当存在接口时，会隐式实现接口，让接口给类提供约束。
// 使用接口调用了结构体中的方法，也可以理解为实现了面向对象中的多态。
func TestInterface2(t *testing.T) {
	gameService := new(GameService)
	gameService.Start()
	gameService.Log("hello")
}

// 空接口
// 空接口是接口的特殊形式，没有任何方法，因此任何具体的类型都可以认为实现了空接口。
func TestNullInterface(t *testing.T) {
	var any interface{}

	any = 1
	t.Log(any)

	any = "Hello"
	t.Log(any)
}

// 空接口作为函数参数：
func Tt(i interface{}) {
	fmt.Printf("%T\n", i)
}

func TestInterface3(t *testing.T) {
	Tt(3)
	Tt("Hello")
	Tt(new(Logger))
	// 利用空接口，可以实现任意类型的存储：
	m := make(map[string]interface{})
	m["name"] = "eanson"
	m["age"] = 18
	map2 := map[interface{}]interface{}{}
	map2["xining"] = 2
}
