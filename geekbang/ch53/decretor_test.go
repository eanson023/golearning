package ch53

import (
	"fmt"
	"testing"
)

func hello() {
	fmt.Println("hello world")
}

type Func func()

func Decrator(inner Func) Func {
	return func() {
		fmt.Println("hello decrator")
		inner()
		fmt.Println("end decrator")
	}
}

func TestD(t *testing.T) {
	bibao := Decrator(hello)
	bibao()
}
