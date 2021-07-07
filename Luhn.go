package libcore

import "errors"

func LuhnMake(str string) (a string, flagOk bool) {
	val := str + "0"
	val_size := len(val)
	flag_parity_str := libcore.IsParity(uint64(val_size))


	sum := 0
	for i := 0; i < val_size; i++ {
		if (val[i] < byte('0')) || (val[i] > byte('9')) {
			return
		}
		ch := val[i] - byte('0')


		flag_parity_idx := libcore.IsParity(uint64(i))
		if flag_parity_str == flag_parity_idx {
			tmp := int(ch * 2)
			if (tmp > 9) {
				sum += (tmp - 9)
			} else {
				sum += tmp
			}
		} else {
			sum += int(ch)
		}
	}


	if (int(int(sum / 10) * 10) == sum) {
		a = val
		flagOk = true
		return
	}


	luhn_sum := ((int(int(sum / 10)) + 1) * 10) - sum
	a = str + fmt.Sprintf("%d", luhn_sum)
	flagOk = true
	return
}


func LuhnCheck(val string) (flagOk bool) {
	val_size := len(val)
	flag_parity_str := libcore.IsParity(uint64(val_size))


	sum := 0
	for i := 0; i < val_size; i++ {
		if (val[i] < byte('0')) || (val[i] > byte('9')) {
			return
		}
		ch := val[i] - byte('0')


		flag_parity_idx := libcore.IsParity(uint64(i))
		if flag_parity_str == flag_parity_idx {
			tmp := int(ch * 2)
			if (tmp > 9) {
				sum += (tmp - 9)
			} else {
				sum += tmp
			}
		} else {
			sum += int(ch)
		}
	}


	if (int(int(sum / 10) * 10) != sum) {
		return
	}


	flagOk = true
	return
}


func luhnTest(x string, y string) bool {


	a, flagOk := LuhnMake(x)
	if flagOk == false {
		return false
	}

	fmt.Printf("===========================\n")
	fmt.Printf("x: %s\n", x)
	fmt.Printf("y: %s\n", y)
	fmt.Printf("a: %s\n", a)


	if a != y {
		return false
	}

	if LuhnCheck(a) == false {
		return false
	}


	return true
}


func LuhnTest() (err error) {
	if luhnTest("637683000000110", "6376830000001100") == false {
		err = errors.New("test01")
		return
	}

	if luhnTest("637683000000100", "6376830000001001") == false {
		err = errors.New("test02")
		return
	}

	if luhnTest("637683000000109", "6376830000001092") == false {
		err = errors.New("test03")
		return
	}

	if luhnTest("637683000000118", "6376830000001183") == false {
		err = errors.New("test04")
		return
	}

	if luhnTest("637683000000108", "6376830000001084") == false {
		err = errors.New("test05")
		return
	}

	if luhnTest("637683000000117", "6376830000001175") == false {
		err = errors.New("test06")
		return
	}

	if luhnTest("637683000000107", "6376830000001076") == false {
		err = errors.New("test07")
		return
	}

	if luhnTest("637683000000102", "6376830000001027") == false {
		err = errors.New("test08")
		return
	}

	if luhnTest("637683000000111", "6376830000001118") == false {
		err = errors.New("test09")
		return
	}

	if luhnTest("637683000000101", "6376830000001019") == false {
		err = errors.New("test10")
		return
	}

	return
}
