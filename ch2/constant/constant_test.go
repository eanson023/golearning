package constant

import (
	"testing"
)

const (
	Monday  = iota + 1
	Tuesday = 3
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

const (
	a = iota
	b
	c
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}
func TestConstantBit(t *testing.T) {
	a := 7 //0111
	t.Log(Readable, Writeable, Executable)
	t.Log(a&Readable == Readable, a&Writeable, a&Executable)
}

func TestConstant2(t *testing.T) {
	t.Log(a, b, c)
}
