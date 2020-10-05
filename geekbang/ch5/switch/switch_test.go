package _switch

import (
	"runtime"
	"testing"
)

func TestSwitch1(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X.")
	case "linux":
		t.Log("Linux.")
	default:
		t.Logf("%s", os)
	}
}
func TestSwitch2(t *testing.T) {
	n := -3
	a := 2
	switch {
	case 0 <= n, a <= 3:
		t.Log("0-3")
	case 4 <= n && n <= 6:
		t.Log("4-6")
	}
}

//case多条件
func TestSwitch3(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

//向if else一样的switch
func TestSwitch4(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1:
			t.Log("odd")
		}
	}
}
