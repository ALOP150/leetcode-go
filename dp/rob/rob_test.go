package rob

import "testing"

func Test_rob(t *testing.T) {
	data := []int{1, 2, 3, 1}
	expected := 4
	real := rob(data)

	t.Logf("expected:%d,real:%d", expected, real)
	if expected == real {
		t.Log("success")
	} else {
		t.Fatalf("fail")
	}

	data = []int{2, 7, 9, 3, 1}
	expected = 12
	real = rob(data)

	t.Logf("expected:%d,real:%d", expected, real)
	if expected == real {
		t.Log("success")
	} else {
		t.Fatalf("fail")
	}

	data = []int{1, 2}
	expected = 2
	real = rob(data)

	t.Logf("expected:%d,real:%d", expected, real)
	if expected == real {
		t.Log("success")
	} else {
		t.Fatalf("fail")
	}

	data = []int{2, 1, 1, 2}
	expected = 4
	real = rob(data)

	t.Logf("expected:%d,real:%d", expected, real)
	if expected == real {
		t.Log("success")
	} else {
		t.Fatalf("fail")
	}
}
