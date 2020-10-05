package ponym

import (
	"fmt"
	"testing"
)

//类型别名
type Code string

//多态

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (*GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

func (*JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

//调用方法时的多态可以哦
// 接口对应的是一个指针类型的实例
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T\t%v\n", p, p.WriteHelloWorld())
}

func TestPolymoriphism(t *testing.T) {
	//这种多态不行哦
	// var goP *Programmer = new(GoProgrammer)
	//错误写法 接口对应的是一个指针类型的实例
	// var goP GoProgrammer = GoProgrammer{}
	var goP *GoProgrammer = &GoProgrammer{}
	var javaP *JavaProgrammer = new(JavaProgrammer)
	writeFirstProgram(goP)
	writeFirstProgram(javaP)

}
