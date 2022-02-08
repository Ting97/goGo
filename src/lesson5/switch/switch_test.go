package _switch

import "testing"

func TestSwitch(t *testing.T) {
	// 1. 条件表达式不限制为常量或者整数
	// 2. 单个case中可以出现多个结果集，使用逗号分割
	// 3. 与C语言规则相反，Go 语言不需要用break来明确退出一个case
	// 4. 可以不设定switch 之后的条件表达式，在此种情况下，整个switch结构与if else 逻辑作用等同

	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}
