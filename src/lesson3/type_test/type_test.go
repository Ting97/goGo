package main

import "testing"

func TestType(t *testing.T) {
	// Go语言 基本类型
	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8
	// rune // alias for int32, represents a Unicode code point
	// float32 float64
	// complex64 complex128

	// 不支持隐式类型转换

	// 类型预定义的值
	// math.MaxInt64
	// math.MaxFloat64

	// 指针类型
	// 与其他编程语言差异
	// 不支持指针运算
	// string 是值类型，默认的初始化值微空字符串，不是nil
}

type LL int64 // 可以将类型进行别名转化

func TestImplicit(t *testing.T) {
	var a LL = 1
	t.Log(a)
}

func TestPoint(t *testing.T) {
	a := 1
	// 获取a的地址，但是不能用其去做运算
	aPtr := &a

	t.Log(a, aPtr)
	// %T可以获取变量类型
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("s == ''")
	}
}
