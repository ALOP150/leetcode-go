package issubsequence

import "testing"

func Test_isSubsequence(t *testing.T) {
	inputS := "abc"
	inputT := "ahbgdc"
	expected := true
	output := isSubsequence(inputS, inputT)
	t.Logf("expected:%t,real:%t", expected, output)
	if output != expected {
		t.Fatal("fail")
	} else {
		t.Log("success")
	}

	inputS = ""
	inputT = "ahbgdc"
	expected = true
	output = isSubsequence(inputS, inputT)
	t.Logf("expected:%t,real:%t", expected, output)
	if output != expected {
		t.Fatal("fail")
	} else {
		t.Log("success")
	}

	inputS = "z"
	inputT = ""
	expected = false
	output = isSubsequence(inputS, inputT)
	t.Logf("expected:%t,real:%t", expected, output)
	if output != expected {
		t.Fatal("fail")
	} else {
		t.Log("success")
	}

	inputS = "aaaaa"
	inputT = "bbaaa"
	expected = false
	output = isSubsequence(inputS, inputT)
	t.Logf("expected:%t,real:%t", expected, output)
	if output != expected {
		t.Fatal("fail")
	} else {
		t.Log("success")
	}
}
