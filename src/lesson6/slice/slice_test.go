package slice

import "testing"

func TestSliceInit(t *testing.T) {
	// 切片 可变长 可以理解成结构体{*data, len, cap}
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	// make 代表构建一个数据方法， 3代表大小，5代表容量
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	for i := 1; i < 10; i++ {
		s2 = append(s2, i)
		t.Log(len(s2), cap(s2))
	}
	t.Log(s2[0], s2[3])
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:7]
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceArrayComparing(t *testing.T) {
	// 数组和切片的差异
	// 1. 容量是否可伸缩
	// 2. 是否可比较

	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}
	t.Log(s1, s2)
	//if s1 == s2 {

	//}
}
