package mincostclimbingstairs

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	input := []int{10, 15, 20}
	expected := 15
	output := minCostClimbingStairs(input)
	t.Logf("expected:%d,output:%d\n", expected, output)
	if expected == output {
		t.Log("success")
	} else {
		t.Fail()
	}

	input = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	expected = 6
	output = minCostClimbingStairs(input)
	t.Logf("expected:%d,output:%d\n", expected, output)
	if expected == output {
		t.Log("success")
	} else {
		t.Fail()
	}
}
