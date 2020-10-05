package overflow

import (
	"fmt"
	"testing"
)

func TestOverflow(t *testing.T) {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
}

func TestRune(t *testing.T) {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}
