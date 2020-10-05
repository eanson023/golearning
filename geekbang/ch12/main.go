package main

import (
	"fmt"
	"go_learning/ch12/mypack"
)

type Student mypack.Person

func (s *Student) Study() {
	fmt.Println("study")
}

func main() {
	s := &Student{}
	s.Study()
}
