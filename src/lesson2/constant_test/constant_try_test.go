package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const ( // 连续常量可以做到简洁定义
	Writeable = 1 << iota
	Readable
	ReadWriteable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Writeable, Readable, ReadWriteable)
}
