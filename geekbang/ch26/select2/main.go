package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		for range time.Tick(time.Second * 2) {
			select {
			case <-stop:
				fmt.Println("stop the go func")
				return
			default:
				fmt.Println("do go func things")
			}
		}
	}()
	time.Sleep(time.Second * 5)
	stop <- true
	fmt.Println("stop the work")
}
