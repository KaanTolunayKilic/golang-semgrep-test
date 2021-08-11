package model

import "testing"

func TestContainsBody(t *testing.T) {
	const body string = "Hello Word"
	const trueSubstr string = "ello"
	const falseSubstr string = "xyz"

	comment := Comment{Body: body}

	if !comment.Contains(trueSubstr) {
		t.Errorf("%s contains substring %s; want true", body, trueSubstr)
	}

	if comment.Contains(falseSubstr) {
		t.Errorf("%s contains not substring %s; want false", body, falseSubstr)
	}
}
