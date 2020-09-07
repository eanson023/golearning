package ch34

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 100
		},
	}
	// 此时为空的 会创建一下
	v := pool.Get().(int)
	fmt.Printf("%p,%v\n", &v, v)
	// 放入私有对象
	pool.Put(3)
	// https://www.jianshu.com/p/ce0d8dd011d7
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Printf("%p,%v\n", &v1, v1)
	v1, _ = pool.Get().(int)
	fmt.Println(v1)
}

// 多协程的情况小
func TestSyncPoolInMultiGroutine(t *testing.T) {
	var pool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new project.")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Test(t *testing.T) {
	stu := &Stu{
		Name: "Eanson",
	}
	fmt.Printf("%p,%v\n", stu, stu)
	stu.f1()
	stu.f2()
	fmt.Printf("%p,%v\n", stu, stu)
}

type Stu struct {
	Name string
}

func (stu *Stu) f1() {
	fmt.Printf("f1()%v,%p\n", stu.Name, stu)
}

func (stu Stu) f2() {
	fmt.Printf("f2()%v,%p\n", stu.Name, &stu)
	stu.Name = "Hahahah"
}
