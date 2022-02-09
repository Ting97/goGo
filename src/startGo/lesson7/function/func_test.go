package function

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	// 1. 可以返回多个值
	// 2. 所有参数都是值传递：slice、map、channel会有传引用的错觉
	// 3. 函数可以作为变量的值
	// 4. 函数可以作为参数和返回值
	_, i := returnMultiValues()
	t.Log(i)

	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))

	t.Log(sum(1, 2, 3))

}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ans := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ans
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 多重数组
func sum(op ...int) int {
	ans := 0
	for _, i := range op {
		ans += i
	}
	return ans
}

func TestDefer(t *testing.T) {
	// 延迟函数 类似于 Java中的try catch finally 中的finally块
	defer func() {
		t.Log("Clear resource")
	}()
	t.Log("Started")
	//panic("Fin Error")
}
