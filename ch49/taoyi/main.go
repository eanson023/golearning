package main

func toHeap() *int {
	var x int
	return &x
}

func toStack() int {
	x := new(int)
	*x = 1
	return *x
}

// 由于栈的性能相对较高，变量是分配到了栈，还是堆中，对程序的性能和安全有较大影响。
// 逃逸分析是一种确定指针动态范围的方法，用来分析程序的哪些地方可以访问到指针。当一个变量或对象在子程序中分配内存时，一个指向变量或对象的指针可能逃逸到其他执行线程中，甚至去调用子程序。

// 指针逃逸：一个对象的指针在任何一个地方都可以访问到。

// 逃逸分析的结果可以用来保证指针的声明周期只在当前进程或线程中。

//  go run -gcflags '-m -l' main.go

// # command-line-arguments
// .\main.go:4:6: moved to heap: x
// .\main.go:9:10: new(int) does not escape
// .\main.go:29:18: new(int) escapes to heap
// .\main.go:34:10: new(int) does not escape
func main() {

}

var global *int

func f() {
	var x int
	x = 1
	global = &x
}
func g() {
	y := new(int)
	*y = 1
}
