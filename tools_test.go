package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var tools Tools

	s := tools.RandomString(10)
	if len(s) != 10 {
		t.Errorf("expected string of length 10, got %d", len(s))
	}
}
