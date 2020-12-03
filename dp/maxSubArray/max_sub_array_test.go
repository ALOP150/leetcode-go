package maxsubarray

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	real := maxSubArray(nums)
	t.Logf("expected:%d,real:%d\n", expected, real)
	if expected != real {
		t.Fatal("fail")
	} else {
		t.Logf("success")
	}
}
