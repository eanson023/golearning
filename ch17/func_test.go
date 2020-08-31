package ch17

import (
	"bytes"
	"fmt"
	"testing"
)

/*
注意：

- 函数名首字母小写为私有，大写为公有；
- 参数列表可以有0-多个，多参数使用逗号分隔，不支持默认参数；
- 返回值列表返回值类型可以不用写变量名
- 如果只有一个返回值且不声明类型，可以省略返回值列表与括号
- 如果有返回值，函数内必须有return
*/

// Go中函数常见写法：

//无返回值，默认返回0，所以也可以写 func fn() int {}
func fn() {}

//Go推荐给函数返回值起一个变量名
func fn1() (result int) {
	return 1
}

//第二种返回值写法
func fn2() (result int) {
	result = 1
	return
}

// 多返回值情
func fn3() (int, int, int) {
	return 1, 2, 3
}

//Go返回值推荐多返回值写法：
func fn4() (a int, b int, c int) {
	a, b, c = 1, 2, 3
	return
}

func fn5() (a, b, c int) { //多个参数类型如果相同，可以简写为： a,b,c int
	return 1, 2, 3
}

//值传递和引用传递
// 不管是值传递还是引用传递，传递给函数的都是变量的副本，不同的是，值传递的是值的拷贝，引用传递的是地址的拷贝，
// 一般来说，地址拷贝效率高，因为数据量小，而值拷贝决定拷贝的 数据大小，数据越大，效率越低。

// 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。

//-------------------------------------

// 可变参数

// 可变参数变量是一个包含所有参数的切片。如果要在多个可变参数中传递参数 ，
// 可以在传递时在可变参数变量中默认添 加“ ...”，将切片中的元素进行传递，而不是传递可变参数变量本身。

// 对可变参数列表进行遍历
func joinStrings(slist ...string) (res string) {
	var buf bytes.Buffer
	for _, s := range slist {
		buf.WriteString(s)
	}
	return buf.String()
}

func joinStringsPointer(slist ...*string) (res string) {
	var buf bytes.Buffer
	for _, s := range slist {
		buf.WriteString(*s)
	}
	return buf.String()
}

func TestKebianArgs(t *testing.T) {
	var sl1 []string = []string{"11", "22", "33"}
	res := joinStrings(sl1...)
	t.Log(res)
}

//参数传递
func rawPrint(rawList ...interface{}) {
	for index, a := range rawList {
		fmt.Println(index, a)
	}
}

//封装打印函数
func print(slist ...interface{}) {
	//将slist可变参数切片完整传递下一个函数
	rawPrint(slist...)
}

func TestArgTrans(t *testing.T) {
	print(1, "hello", 3)
}

// 匿名函数

// 匿名函数可以看做函数字面量，所有直接使用函数类型变量的地方都可以由匿名函数代替。
// 匿名函数可以直接赋值给函数变量，可以当做实参，也可以作为返回值使用，还可以直接被调用。

func TestAnoFunc(t *testing.T) {
	a := 3
	f1 := func(num int) { //f1即为匿名函数
		t.Log(num)
	}
	f1(a)

	//匿名函数自调
	func() {
		t.Log(a) //匿名函数访问外部变量
		//这里有个() ()代表发出调用函数的信号
	}()

	//匿名函数实战：取最大值,最小值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)

	t.Log(x, y)
}

// Go函数特性总结

/*
- 支持有名称的返回值；
- 不支持默认值参数；
- 不支持重载；
- 不支持命名函数嵌套，匿名函数可以嵌套；
- Go函数从实参到形参的传递永远是值拷贝，有时函数调用后实参指向的值发生了变化，是因为参数传递的是指针的拷贝，实参是一个指针变量，传递给形参的是这个指针变量的副本，实质上仍然是值拷贝；
- Go函数支持不定参数；
*/

//----------------------

// 两个特殊函数

/*
####  init函数

Go语言中，除了可以在全局声明中初始化实体，也可以在init函数中初始化。init函数是一个特殊的函数，
它会在包完成初始化后自动执行，执行优先级高于main函数，并且不能手动调用init函数，
每一个文件有且仅有一个init函数，初始化过程会根据包的以来关系顺序单线程执行。
*/

func init() {
	fmt.Println("init....")
}

func TestMain(t *testing.T) {
	t.Log("main......")
}

/*
#### new函数

new函数可以用来创建变量。表达式`new(T)`将创建一个T类型的匿名变量，
初始化为T类型的零值，然后返回变量地址，返回的指针类型为`*T`：
*/

func TestNewFunc(t *testing.T) {
	p := new(int) //p 为*int类型，只想匿名的int变量
	t.Log(*p)     //0
	*p = 2        //设置int匿名变量值为2
	t.Log(*p)
}
