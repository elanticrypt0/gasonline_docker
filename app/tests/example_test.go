package test

import "testing"

func TestExample(t *testing.T) {
	got := 1
	want := 1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Logf("got %q, wanted %q", got, want)
	}
}
