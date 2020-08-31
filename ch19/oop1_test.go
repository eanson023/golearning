package ch19

import (
	"fmt"
	"testing"
)

// 模拟构造函数
// Go和传统的面向对象语言如Java有着很大区别。结构体没有构造函数初始化功能
type Person struct {
	Name string
	Age  int
}

func NewPersonByName(name string) *Person {
	return &Person{
		Name: name,
	}
}
func NewPersonByAge(age int) *Person {
	return &Person{
		Age: age,
	}
}

// 贴士：因为Go没有函数重载，为了避免函数名字冲突，
// 使用了`NewPersonByName`和`NewPersonByAge`两个不同的函数表示不同的`Person`构造过程。

func TestConstructor(t *testing.T) {
	p := NewPersonByName("eanson")
	t.Log(p)
	t.Logf("%T\n", p)
	t.Log(p.Age)
}

// 父子关系结构体初始化
// Person可以看做父类，Student是子类，子类需要继承父类的成员
type Student struct {
	Person
	ClassName string
}

//构造父类
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

//构造子类
func NewStudent(className string) *Student {
	p := &Student{}
	p.ClassName = className
	return p
}

func TestConstructor2(t *testing.T) {
	s := NewStudent("三班")
	t.Log(s)
}

// Go中的面向对象初识
// 在Go中，可以给任意类型(出了指针)添加相应方法:

type Integer int

//类似C#扩展方法
func (i Integer) Less(j Integer) bool {
	return i < j
}

func TestExtensionMethod(t *testing.T) {
	var i Integer = 1
	t.Log(i.Less(20))
}

// 二 方法
// Golang 中的方法是作用在指定的数据类型上的(即:和指定的数据类型绑定)，
// 因此自定义类型，都可以有方法，而不仅仅是 struct。

// 方法的声明和调用：
/*
func (receiver type) methodName(参数列表) (返回值列表){
	//方法体
	return 返回值
}
*/

// 方法与函数的示例
//一个run函数
func run(p *Person, name string) {
	p.Name = name
	fmt.Println("函数 run...", p.Name)
}

//一个run方法
func (p *Person) run(name string) {
	p.Name = name
	fmt.Println("方法 run...", p.Name)
}

func TestMethodAndFunc(t *testing.T) {
	//实例化一个对象(结构体)
	p1 := NewPerson("eanson", 18)

	//执行普通方法
	run(p1, "张三") //函数 run... 张三

	//执行方法
	p1.run("严胜") //方法 run... 严胜
}

/*
#### Go方法本质

Go的方法是一种作用于特定类型变量的函数，这种特定类型的变量叫做接收器（Receiver）。如果特定类型理解为结构体或者“类”时，接收器就类似于其他语言的this或者self。

在Go中，接收器可以是任何类型，不仅仅是结构体，依此我们看出，Go中的方法和其他语言的方法类似，但是Go语言的接收器强调方法的作用对象是实例。

方法与函数的区别就是：函数没有作用对象。

上述Person案例中，接收器类型是`*Person`，属于指针类型，非常接近Java中的`this`，由于指针的特性，调用方法时，
修改接收器指针的任意长远变量，在方法结束后，修改都是有效的。

当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效
*/

func (p Person) noPointerRun(name string) {
	p.Name = name
	fmt.Println("方法 noPointerRun...", p.Name)
}

func noPointerRun(name string, p Person) {
	p.Name = name
	fmt.Println("方法 noPointerRun...", p.Name)
}
func TestNoPointer(t *testing.T) {
	//实例化一个对象(结构体)
	p1 := NewPerson("eanson", 18)
	p1.noPointerRun("严胜")
	//未改变
	t.Log(p1.Name)
	noPointerRun("严胜", *p1)
	//还是未改变
	t.Log(p1.Name)
}

// 一般情况下，小对象由于复制时速度较快，适合使用非指针接收器，大对象因为复制性能较低，
// 适合使用指针接收器，此时再接收器和参数之间传递时不进行复制，只传递指针。
