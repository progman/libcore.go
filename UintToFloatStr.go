//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"fmt"
	"testing"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
"666797225", "6" -> "666.797225"
"255528067780000000000", "18" -> "255.528067780000000000"
"0", "3" -> "0.000"
*/
func UintToFloatStr(source string, decimalSize string) (target string, err error) {

	if IsUint(source) == false {
		err = fmt.Errorf("source is invalid")
		return
	}


	var decimalSizeInt64 int64
	decimalSizeInt64, err = StrToSint(decimalSize)
	if err != nil {
		err = fmt.Errorf("decimalSize is invalid")
		return
	}


	for {
		if len(source) >= int(decimalSizeInt64) {
			break
		}
		source = "0" + source
	}


	var part1 string = source[0 : len(source) - int(decimalSizeInt64)]
	var part2 string = source[len(source) - int(decimalSizeInt64) : len(source)]


	target = part1 + "." + part2


// convert ".x" to "0.x" if need
	if rune(target[0]) == rune(byte('.')) {
		target = "0" + target
	}


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func UintToFloatStr_test(t *testing.T) {
	var err error
	var target string
	var expect string


	expect = "666.797225"
	target, err = UintToFloatStr("666797225", "6")
	if err != nil {
		t.Errorf("UintToFloatStr() = %v, want %v", err, nil)
	}
	if target != expect {
		t.Errorf("UintToFloatStr() = %s, want %s", target, expect)
	}


	expect = "255.528067780000000000"
	target, err = UintToFloatStr("255528067780000000000", "18")
	if err != nil {
		t.Errorf("UintToFloatStr() = %v, want %v", err, nil)
	}
	if target != expect {
		t.Errorf("UintToFloatStr() = %s, want %s", target, expect)
	}


	expect = "0.000"
	target, err = UintToFloatStr("0", "3")
	if err != nil {
		t.Errorf("UintToFloatStr() = %v, want %v", err, nil)
	}
	if target != expect {
		t.Errorf("UintToFloatStr() = %s, want %s", target, expect)
	}


	fmt.Printf("UintToFloatStr(): PASS\n")
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
