package test

import "testing"

// 测试文件一定是 _test结尾
// 大写代表包外可访问
func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}
