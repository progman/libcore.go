//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
//import (
//	"time"
//)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func checkIsoDatetime(dateStr string, template string) (flagOk bool) {
	size = len(dateStr)


	if size != len(template) {
		return
	}


	flagOk = true
	for i := 0; i < size; i++ {

		if template[i] == byte('0') {
			if (dateStr[i] < byte('0')) || (dateStr[i] > byte('9')) {
				flagOk = false
				break
			}
			continue
		}

		if dateStr[i] != template[i] {
			flagOk = false
			break
		}
	}


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func IsIsoDatetime(dateStr string) (flagOk bool) {

	for {
		if checkIsoDatetime(dateStr, "0000-00-00 00:00:00") == true {
			break
		}
		if checkIsoDatetime(dateStr, "0000-00-00 00:00:00.000000") == true {
			break
		}
		if checkIsoDatetime(dateStr, "0000-00-00 00:00:00.000000000") == true {
			break
		}

		if checkIsoDatetime(dateStr, "0000-00-00T00:00:00Z00:00") == true {
			break
		}
		if checkIsoDatetime(dateStr, "0000-00-00T00:00:00.000000000Z00:00") == true {
			break
		}


		return
	}


	flagOk = true
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
