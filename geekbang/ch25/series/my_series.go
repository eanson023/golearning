package series

import (
	"errors"
	"fmt"
)

/*
相同包源文件中可以有多个init方法
*/
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

var LessThanError = errors.New("n should be  not less than 2")
var GreaterThanError = errors.New("n should be  not greater than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, LessThanError
	}
	if n > 100 {
		return nil, GreaterThanError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
