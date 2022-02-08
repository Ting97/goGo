package array

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	// 初始化
	var arr [3]int
	t.Log(arr[0], arr[1])
	// 用... 代替表示容量大小和定义的一致
	arr2 := [...]int{1, 2, 4, 5}
	// foreach 语句，有idx进行下标的获取
	for idx, e := range arr2 {
		t.Log(idx, e)
	}
	// 不关注下标元素可以使用 _代替
	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArraySplit(t *testing.T) {
	// 截取数组的一部分
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 其实位置到结束位置的前一个数据
	t.Log(arr[2:5])
	t.Log(arr[:])
}
