package main

import (
	"fmt"
)

func main() {
	s1 := "asas"
	s2 := "asas"
	fmt.Println(isAnagram(s1, s2))
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var m = make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
		m[s2[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
