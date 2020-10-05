package main

import (
	"fmt"
)

const (
	KiB = 1024 << (10 * iota)
	MiB
	GiB
	TiB
	PiB
	EiB
)

func main() {
	fmt.Println(KiB)
	fmt.Println(MiB)

}
