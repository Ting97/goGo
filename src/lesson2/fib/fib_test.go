package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	var (
		test  = 1
		test1 = 2
	)

	test1, test = test, test1

	var a = 1
	var b = 1

	for i := 1; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a // 数据交换
	t.Log(a, b)
}
