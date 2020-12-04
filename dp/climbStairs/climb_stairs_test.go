package climbstairs

import "testing"

func Test_climbStairs(t *testing.T) {
	n := 3
	expected := 3
	real := climbStairs(n)
	t.Logf("expected:%d,real:%d", expected, real)
	if expected != real {
		t.Fatal("fail")
	} else {
		t.Log("success")
	}

}
