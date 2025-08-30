//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"fmt"
	"testing"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
	 1      err
	 1 2    false
	 1 1    true
	 1 1.   true
	 1 1.0  true
	 1 1.00 true
	.1 0.1  true
	.1 0.10 true
*/
func IsFloatStrEqual(floatStr1 string, floatStr2 string) (flagEqual bool, err error) {

// are args float?
	var flagFloat1 bool = IsFloat(floatStr1, true)
	var flagFloat2 bool = IsFloat(floatStr2, true)


// args float or integer
	if flagFloat1 == false {
		if IsUint(floatStr1) == false {
			err = fmt.Errorf("floatStr is invalid")
			return
		}
	}
	if flagFloat2 == false {
		if IsUint(floatStr2) == false {
			err = fmt.Errorf("floatStr is invalid")
			return
		}
	}


// if all args integers
	if (flagFloat1 == false) && (flagFloat2 == false) {
		if floatStr1 == floatStr2 {
			flagEqual = true
			return
		}
		return
	}


// mutate to float of need
	if flagFloat1 == false {
		floatStr1 += "."
	}
	if flagFloat2 == false {
		floatStr2 += "."
	}


// convert ".x" to "0.x" if need
	if rune(floatStr1[0]) == rune(byte('.')) {
		floatStr1 = "0" + floatStr1
	}
	if rune(floatStr2[0]) == rune(byte('.')) {
		floatStr2 = "0" + floatStr2
	}


// if lens are equal
	if len(floatStr1) == len(floatStr2) {
		if floatStr1 == floatStr2 {
			flagEqual = true
			return
		}
		return
	}


// get max len
	var max int = len(floatStr1)
	if max < len(floatStr2) {
		max = len(floatStr2)
	}


// expand args
	for {
		if len(floatStr1) == max {
			break
		}
		floatStr1 += "0"
	}
	for {
		if len(floatStr2) == max {
			break
		}
		floatStr2 += "0"
	}


// compare args
	if floatStr1 == floatStr2 {
		flagEqual = true
		return
	}


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func IsFloatStrEqual_test(t *testing.T) {
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


	flagEqual, err = IsFloatStrEqual("0", ".0")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}


	flagEqual, err = IsFloatStrEqual("0", ".00")
	if err != nil {
		t.Errorf("IsFloatStrEqual() = %v, want %v", err, nil)
	}
	if flagEqual != true {
		t.Errorf("IsFloatStrEqual() = %t, want %t", flagEqual, true)
	}


	fmt.Printf("IsFloatStrEqual(): PASS\n")
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
