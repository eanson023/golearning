package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

type Dog struct {
	Pet //匿名字段
}

func (*Pet) Speak() {
	fmt.Println("....")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(name)
}

func (*Dog) Speak() {
	fmt.Println("Wangle！")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Eanson")
}

func TestSubClassAccess(t *testing.T) {

}

func makePetSpeak(p *Pet) {
	p.Speak()
	fmt.Println("\nPet spoke")
}

func TestLSP(t *testing.T) {

	//无法支持LSP原则 子类无法转为父类

	// var dog *Pet = new(Dog)
	// makePetSpeak(dog)
}
