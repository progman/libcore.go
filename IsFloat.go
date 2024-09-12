//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
//import (
//	"errors"
//)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func IsFloat(str string) (flagOk bool) {
	strSize := len(str)

	if strSize == 0 {
//		err = errors.New("str is empty")
		return
	}


	var flagDot bool
	var flagDig bool
	for i := 0; i < strSize; i++ {

		if str[i] == byte('.') {
			if flagDot == false {
				flagDot = true
				continue
			}
//			err = errors.New("str has double dot")
			return
		}


		if (str[i] < byte('0')) || (str[i] > byte('9')) {
//			err = errors.New("str has strange char")
			return
		}

		flagDig = true
	}


	if flagDig == false {
//		err = errors.New("str has not dig char")
		return
	}


	flagOk = true
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//