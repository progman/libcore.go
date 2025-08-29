//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"testing"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func TestPing(t *testing.T) {
	var want string
	var got  string


	want = "Ping"
	got  = Ping();
	if got != want {
		t.Errorf("Ping() = %q, want %q", got, want)
	}
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func TestPong(t *testing.T) {
	var want string
	var got  string


	want = "Pong"
	got  = Pong();
	if got != want {
		t.Errorf("Pong() = %q, want %q", got, want)
	}
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func TestIsFloatStrEqual(t *testing.T) {
	var err error
	var flagEqual bool


	flagEqual, err = IsFloatStrEqual("1", "")
	if err == nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", nil, err)
	}


	flagEqual, err = IsFloatStrEqual("1", "2")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != false {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, false)
	}


	flagEqual, err = IsFloatStrEqual("1", "1")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}


	flagEqual, err = IsFloatStrEqual("1", "1.")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}


	flagEqual, err = IsFloatStrEqual("1", "1.0")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}


	flagEqual, err = IsFloatStrEqual("1", "1.00")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}


	flagEqual, err = IsFloatStrEqual("0.1", "0.1")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}


	flagEqual, err = IsFloatStrEqual("0.1", ".1")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}


	flagEqual, err = IsFloatStrEqual("0.1", ".10")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func TestHello(t *testing.T) {
	TestPing(t)
	TestPong(t)
	TestIsFloatStrEqual(t)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
