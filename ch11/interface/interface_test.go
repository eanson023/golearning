package interface_test

import "testing"

//duck type式接口
type Programer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (*GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Helll,World\")"
}

func TestClient(t *testing.T) {
	var p Programer
	//多态
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
