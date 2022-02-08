package _map

import (
	"testing"
)

func TestMapInit(t *testing.T) {
	//  默认初始化
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m, len(m))
	// 复制初始化
	m1 := map[string]int{}
	m1["one"] = 1
	// 使用make 初始化容量大小
	m2 := make(map[string]int, 10)
	t.Log(m2, len(m2))
}

func TestAccessNotExistingKey(t *testing.T) {
	// 在访问不存在时，会返回0值，不能通过返回nil来判断元素是否存在
	m := map[int]int{}
	m[0] = 0
	if v, ok := m[1]; ok {
		t.Log("exist ", v)
	} else {
		t.Log("not exist")
	}
}

func TestMapRange(t *testing.T) {
	m := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m {
		t.Log(k, v)
	}
}

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	t.Log(mySet)
}
