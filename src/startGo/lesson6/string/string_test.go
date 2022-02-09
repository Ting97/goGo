package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringInit(t *testing.T) {
	// string 是一个数据类型，不是引用或者指针类型
	// string 是只读的 byte slice，len函数可以获取它包含的byte数
	// string 的 byte 数组可以放任何数据

	var s string
	t.Log(s)

	s = "string"
	t.Log(s)

	s = "\xE4\xB8\xA5" // 可以存放任何二进制数据
	t.Log(s)           // 代表"严"

	s = "中"
	t.Log(len(s))

	// rune 获取自字符串的Unicode的存储数据
	c := []rune(s)
	t.Log(c)
	t.Log(c[0])
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
		//t.Log(string(c))
	}
}

func TestStringFun(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, c := range parts {
		t.Log(c)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str " + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
