package interface_test

import (
	"fmt"
	"testing"
)

//duck type式接口
type Programer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (*GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Helll,World\")"
}

func TestClient(t *testing.T) {
	var p Programer
	//多态
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

type IList interface {
	AppendList(l *List)
	AppendLi(li *Li)
}

type List struct {
	// 头节点
	head *Li
	//尾节点
	tail *Li
}

type Sort interface {
	sort()
}

func (list *List) AppendList(l *List) {

}

func (list *List) AppendLi(li *Li) {

}

func NewOl() IList {
	return &Ol{}
}

func NewUl() IList {
	return &Ul{}
}

type Li struct {
	//孩子列表
	children *List
	// 下一个
	next *Li
}

type Ol struct {
	List
}

type Ul struct {
	List
}

func (ol *Ol) sort() {
	fmt.Printf("ol形式排序")
}

func (ul *Ul) sort() {
	fmt.Printf("ul形式排序")
}

func TestC(t *testing.T) {
	// ol := NewOl()
	// ul := NewUl()
	// ol.AppendList(ul)
}
