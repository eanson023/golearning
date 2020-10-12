package main

import (
	"fmt"
)

func clear(s []string) []string {
	var i int
	for j := 0; j < len(s); j++ {
		if j+1 < len(s) && s[j] == s[j+1] {
			continue
		}
		s[i] = s[j]
		i++
	}
	return s[:i]
}

func main() {
	s := []string{"1", "43", "2", "2"}
	s = clear(s)
	fmt.Println(s)
}
