package ch24

import (
	"errors"
	"testing"
)

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

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-10); err == nil {
		t.Log(v)
	} else {
		if err == LessThanError {

		} else if err == GreaterThanError {

		}
		t.Log(err)
	}
}
