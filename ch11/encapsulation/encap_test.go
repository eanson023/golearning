package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
结构体定义 没有,号
*/
type Employee struct {
	Id   string
	Name string
	Age  int
}

//func (e Employee) String() string {
//	fmt.Printf("String() Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

func (e *Employee) String() string {
	fmt.Printf("String() Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	//创建和初始化实例
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 20}
	//返回指针
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Log(*e2)
	//Employee类型
	t.Logf("e is %T", e)
	//如果加上取址符 就可以获得其指针类型
	t.Logf("e is %T", &e)
	//Employee指针类型
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bib", 20}
	/*
		使用带*号的String方法 两个指针是一样的 说明并没有对象复制产生
	*/
	fmt.Printf("TestStructOperations() Address is %x\n", unsafe.Pointer(&e.Name))
	//e := &Employee{"0", "Bib", 20}
	//不管通过指针访问还是实例访问field 因为Go语言会自动判断
	t.Log(e.String())
}

func TestIsInit(t *testing.T) {
	var b bool
	t.Log(b)
}

func hello() *Employee {
	e := new(Employee)
	e.Name = "eanson"
	return e
}
