package waystostep

import "testing"

func Test_waysToStep(t *testing.T) {
	input := 5
	expected := 13
	output := waysToStep(input)
	t.Logf("expected:%d,output:%d", expected, output)
	if output == expected {
		t.Log("success")
	} else {
		t.Fatal("fail")
	}
}
