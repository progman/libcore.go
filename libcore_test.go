package libcore

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	got := Hello();
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
