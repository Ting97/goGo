package operator_test

import (
	"testing"
)

const ( // 连续常量可以做到简洁定义
	Writeable = 1 << iota
	Readable
	ReadWriteable
)

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 5, 7}
	// c:=[...]int{1,2,3,4,5}
	// 长度和数据都一直的时候才是正确的
	t.Log(a == b)
	// t.Log(a == c) 长度不同的数组不可以比较
	k := 7
	//  按位清零 &^
	t.Log(k &^ Readable)
}
