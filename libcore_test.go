package libcore

import "testing"

func TestHello(t *testing.T) {
	var want string
	var got  string


	want = "Ping"
	got  = Ping();
	if got != want {
		t.Errorf("Ping() = %q, want %q", got, want)
	}


	want = "Pong"
	got  = Pong();
	if got != want {
		t.Errorf("Pong() = %q, want %q", got, want)
	}
}
