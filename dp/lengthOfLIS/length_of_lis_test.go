package lengthoflis

import (
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	input := []int{10, 9, 2, 5, 3, 7, 101, 18}
	expected := 4
	output := lengthOfLIS(input)
	t.Logf("expected:%d,output:%d", expected, output)
	if expected != output {
		t.Fail()
	}

	input = []int{0, 1, 0, 3, 2, 3}
	expected = 4
	output = lengthOfLIS(input)
	t.Logf("expected:%d,output:%d", expected, output)
	if expected != output {
		t.Fail()
	}
}
