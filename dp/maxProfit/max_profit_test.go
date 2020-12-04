package maxprofit

import "testing"

func Test_maxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	real := maxProfit(nums)
	expected := 5

	t.Logf("expected:%d,real:%d", expected, real)
	if real != expected {
		t.Fatal("fail")
	}

	nums = []int{1, 2}
	real = maxProfit(nums)
	expected = 1

	t.Logf("expected:%d,real:%d", expected, real)
	if real != expected {
		t.Fatal("fail")
	}
}
