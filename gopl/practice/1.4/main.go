// 练习1.4：修改dup2程序，输出出现重复行的文件的名称
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			defer f.Close()
			countLines(f, counts)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	files := make(map[string]bool)
	for input.Scan() {
		s := input.Text()
		counts[s]++
		if counts[s] > 1 {
			fileInfo, _ := f.Stat()
			files[fileInfo.Name()] = true
		}
	}
	for name := range files {
		fmt.Println(name)
	}
}
