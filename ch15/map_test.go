package ch15

import (
	"sync"
	"testing"
)

// 1.1 map的创建
// Go内置了map类型，map是一个无序键值对集合（也有一些书籍翻译为字典）。

type Person struct {
	ID   string
	Name string
}

func TestMapCreate(t *testing.T) {
	// 普通创建：

	// 声明一个map类型，[]内的类型指任意可以进行比较的类型 int指值类型
	m := map[string]int{"eanson": 18, "wdh": 28}
	t.Log(m)
	t.Log("--------------------")

	// make方式创建map：
	map2 := make(map[string]Person)
	map2["123"] = Person{ID: "qq", Name: "aaa"}
	map2["234"] = *new(Person)
	person, isFind := map2["123"]
	t.Log(isFind)
	t.Log(person)
	person, isFind = map2["13"]
	t.Log(isFind)
	person.Name = "aaa"
	t.Log(person)
}

// 1.2 map的使用
func TestMapUse(t *testing.T) {
	// 通过key操作元素：
	var numbers map[string]int
	numbers = make(map[string]int)
	//put
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	//delete
	delete(numbers, "ten") //删除key为ten的元素

}

/*
注意：

- map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取；
- map的长度是不固定的，也就是和slice一样，也是一种引用类型
- 内置的len函数同样适用于map，返回map拥有的key的数量
- go没有提供清空元素的方法，可以重新make一个新的map，不用担心垃圾回收的效率，因为go中并行垃圾回收效率比写一个清空函数高效很多
- map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
*/

// 1.3 并发安全的map
func TestConcurentMap(t *testing.T) {
	m := make(map[int]int)

	// 编译会有错误提示：`fatal error: concurrent map read and map write`，即出现了并发读写，
	// 因为用两个并发程序不断的对map进行读和写，产生了竞态问题。map内部会对这种错误进行检查并提前发现。

	// Go内置的map只有读是线程安全的，读写是线程不安全的。

	// 需要并发读写时，一般都是加锁，但是这样做性能不高，在go1.9版本中提供了更高效并发安全的sync.Map。
	go func() {
		for { //无限写入
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1] //无限读取
		}
	}()

	//无限循环，让并发程序在后台运行
	for {
	}
}

func TestSyncMap(t *testing.T) {
	//同步Map
	var scence sync.Map

	//保存键值对
	scence.Store("id", 1)
	scence.Store(1, 2)

	//根据键取值
	t.Log(scence.Load("id"))

	//遍历
	scence.Range(func(k, v interface{}) bool {
		t.Log(k, v)
		return true
	})

	// 注意：map没有提供获取map数量的方法，可以在遍历时手动计算。sync.Map为了并发安全。损失了一定的性能。
	go func() {
		for { //无限写入
			scence.Store(1, 1)
		}
	}()

	go func() {
		for {
			_, _ = scence.Load(1)
		}
	}()

	//无限循环，让并发程序在后台运行
	for {
	}

}
