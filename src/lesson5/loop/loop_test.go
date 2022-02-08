package loop

import "testing"

func TestFor(t *testing.T) {
	// for
	for i := 1; i < 5; i++ {
		t.Log(i)
	}
	// 类似 while
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
	// 类似 while(true)
	for {
		if n > 5 {
			break
		}
	}
}
