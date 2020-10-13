package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq()
}

func wordfreq() {
	var counts = make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner reading error:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%10v\t%10v\n","单词","次数")
	for k, v := range counts {
		fmt.Printf("%10v\t%10v\n", k, v)
	}
}
