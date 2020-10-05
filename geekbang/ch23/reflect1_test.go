package ch23

import (
	"fmt"
	"reflect"
	"testing"
)

type Value reflect.Value

type Student struct {
	Age  int
	Name string
}

type SQL interface {
	CreateSQL() string
}

func (s *Student) CreateSQL() string {
	sql := fmt.Sprintf("insert into student values(%d,%s)", s.Age, s.Name)
	return sql
}

func Test11(t *testing.T) {
	o := Student{
		Age:  20,
		Name: "Eanson",
	}
	fmt.Println(o.CreateSQL())
}

func Test22(t *testing.T) {
	var num float64 = 1.2345
	t.Log("type: ", reflect.TypeOf(num))
	t.Log("value", reflect.ValueOf(num))
	// 通过Kind类型,可以方便的验证反射的类型是否相同。
	t.Log("Kind is float64:", reflect.TypeOf(num).Kind() == reflect.Float64)
}

/*
反射转换为接口
Value中的Interface() 方法以空接口的形式返回reflect.Value中的值。
如果要进一步获取空接口的的真实值,则可以通过接口的转换语法对接口进行转换。
*/

func Test23(t *testing.T) {
	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	t.Log(convertPointer)
	t.Log(convertValue)

	t.Log(&num == convertPointer)
	t.Log(num == convertValue)
}

/*
除了直接转换外,reflect.Value 提供了一些转换到具体类型的方法。这些方法经过了特殊的处理,不管内部是int8,int16,int32，通过Int()方法最后都转换为int64。另外这些特殊的方法可以加速转换的速度。
*/
func Test24(t *testing.T) {
	var a int8 = 20
	var x int64 = reflect.ValueOf(a).Int()
	t.Log(x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	t.Logf("type:%T value:%v\n", y, y)
	//不一样
	t.Log(&b, &y)

	x = reflect.ValueOf(&a).Int()
}

/*
反射与间接访问Elem()
reflect.Value的Elem方法可以解决上面的问题,
如果reflect.Value内部存储值为指针或接口,则Elem方法返回指针或接口指向的数据。
*/

func Test25(t *testing.T) {
	aa := 56
	// Elem()等于取值符* 获取值
	x := reflect.ValueOf(&aa).Elem().Int()
	t.Logf("%T,%v\n", x, x)
	// 但是,如果Value存储的不是指针或接口,则仍然在运行时出错,因此在使用时要小心
	x = reflect.ValueOf(aa).Elem().Int()
}

/*
Elem()方法是非常必要的,因为在介绍接口时了解过,接口中存储的是指针,那么我们究竟要获取/修改的是指针本身还是指针指向的数据？为了更好的理解,看一个特殊的例子,反射值包含了接口的指针时
*/

func Test26(t *testing.T) {
	var z = 123
	var y = &z
	t.Logf("%T", y)
	//存储z的指针
	var inter interface{} = y
	t.Logf("%T,%v", inter, inter)

	//获取空接口的指针Value类型(接口中存的指针z)
	interPtrValue := reflect.ValueOf(&inter)

	//取出指针value中的值interface{}(这里是存的指针) 返回inetrValue
	var interValue reflect.Value = interPtrValue.Elem()
	// 接口类型
	t.Logf("%T,%v,%v", interValue, interValue, interValue.Kind())
	t.Logf("%T,%v", interValue.Interface().(*int), interValue.Interface().(*int))
	t.Log("&z==interValue.Interface()??", &z == interValue.Interface().(*int))
}

/*
修改反射的值
有多种方式可以修改反射中存储的值,例如reflect.value类型的Set方法

func (v Value) Set(x Value)
该方法的参数仍然是reflect.value类型。但是要求反射中存储的具体类型必须是指针。
*/
func Test27(t *testing.T) {
	var num float32 = 1.2345
	pointe := reflect.ValueOf(num)
	// reflect: reflect.Value.Set using unaddressable value [recovered]
	pointe.Set(reflect.ValueOf(789))
}

/*
只有反射中存储的实际值是指针,才能够赋值。否则是没有意义的。正如在接口中看到的。如果接口中存储的值实际是一个副本,对他进行指针方法调用会带来混淆。
*/
func Test28(t *testing.T) {
	var num float32 = 1.2345
	pointe := reflect.ValueOf(&num)
	t.Log(pointe.Kind())
	t.Log("settability of pointer:", pointe.CanSet())
	newValue := pointe.Elem()
	t.Log("settability of pointer:", newValue.CanSet())
	t.Log(newValue.Kind())
	newValue.SetFloat(123)
	t.Log(num)
}

/*
结构体与反射
在应用反射的案例中,大部分情况都涉及到结构体。假设现在有User结构体以及方法ReflectCallFunc。
*/
type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) ReflectCallFunc() {
	fmt.Println("jsonon RefelectCallFunc")
}

func TestRefStruct(t *testing.T) {
	var user User = User{
		Id:   111,
		Name: "Eanson",
		Age:  18,
	}
	getType := reflect.TypeOf(user)
	t.Log("get Type is:", getType.Name())
	getVale := reflect.ValueOf(user)
	t.Log("get all fields is:", getVale)
	t.Log("-----------------------------")
	/*
			反射遍历结构体字段
		如果希望遍历获取结构体中字段的名字以及方法,需要通过NumField() 函数获取结构体中字段的个数。
	*/
	for i := 0; i < getVale.NumField(); i++ {
		field := getType.Field(i)
		value := getVale.Field(i).Interface()
		t.Logf("%s:%v=%v\n", field.Name, field.Type, value)
	}
}

// 反射修改结构体字段
// type Point struct {
// 	x int
// 	y float64
// }

func TestStructChange(t *testing.T) {
	// point := Point{
	// 	x: 11,
	// 	y: 12.98,
	// }
	var s struct {
		//反射获取字段必须是公共的
		X int
		y float64
	}
	vp := reflect.ValueOf(s)
	t.Logf("%T,%v", vp, vp.Kind())
	vx, vy := vp.Field(0), vp.Field(1)
	t.Log(vx.Kind())
	t.Log(vy)
	// vb := reflect.ValueOf(123)
	t.Log(vx.CanSet())
	// vx.Set(vb)
	a := vx.Interface().(int)
	t.Log(&a, &s.X, a, s.X)
}

type user struct {
	Name string
	Age  int `json:"age" id:"100"` // 结构体标签
}

func TestRef222(t *testing.T) {
	s := user{
		Name: "zs",
		Age:  1,
	}

	typeOfUser := reflect.TypeOf(s)

	// 字段用法
	for i := 0; i < typeOfUser.NumField(); i++ { // NumField 当前结构体有多少个字段
		fieldType := typeOfUser.Field(i) // 获取每个字段
		fmt.Println(fieldType.Name, fieldType.Tag)
	}
}
