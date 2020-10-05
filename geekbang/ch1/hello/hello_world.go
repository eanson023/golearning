package main

import (
	"fmt"
	"os"
)

func main() {
	//以数组形式输出
	// fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
}
