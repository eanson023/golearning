package main

import (
	"fmt"

	"github.com/eanson023/golearning/geekbang/ch12/mypack"
)

type Student mypack.Person

func (s *Student) Study() {
	fmt.Println("study")
}

func main() {
	s := &Student{}
	s.Study()
}
