package main

import (
	"time"
	// "bufio"
	"fmt"
	// "os"
	"sync"
)

func getText() chan string {
	var str string
	ch := make(chan string)
	fmt.Scanln(&str)
	ch <- str
	return ch
}

func main() {
	time.Sleep(time.Second * 5)
	// fmt.Println(<-getText())
	var str string
	fmt.Scanln(&str)
	fmt.Println(str)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				mutex.Lock()
				fmt.Println(i, j)
				mutex.Unlock()
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
