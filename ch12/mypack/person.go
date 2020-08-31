package mypack

import "fmt"

type Person struct {
}

func (p *Person) Run() {
	fmt.Println("run")
}
