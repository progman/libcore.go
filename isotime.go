//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"strconv"
	"fmt"
	"time"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetUnixnanotimeUtcInt64() (unixnanotimeUtc int64) {
	unixnanotimeUtc = time.Now().UTC().UnixNano()
	return
}

func GetUnixnanotimeUtc() (unixnanotimeUtc string) {
	unixnanotimeUtc = fmt.Sprintf("%d", GetUnixnanotimeUtcInt64())
	return
}

func GetUnixmicrotimeUtcInt64() (unixmicrotimeUtc int64) {
	unixmicrotimeUtc = time.Now().UTC().UnixNano() / int64(time.Microsecond)
	return
}

func GetUnixmicrotimeUtc() (unixmicrotimeUtc string) {
	unixmicrotimeUtc = fmt.Sprintf("%d", GetUnixmicrotimeUtcInt64())
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// 1624120658256315911 -> "2021-06-19 19:37:38.256315911"
func Unixnanotime2iso(unixnanotime int64) (iso string) {
	t := time.Unix(0, unixnanotime)

	var year int
	year = t.Year()

	var month int
	month = int(t.Month())

	var day int
	day = t.Day()

	var hour int
	hour = t.Hour()

	var minute int
	minute = t.Minute()

	var second int
	second = t.Second()

	var nanosecond int
	nanosecond = t.Nanosecond()

	iso = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%09d", year, month, day, hour, minute, second, nanosecond)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// "1624120658256315911" -> "2021-06-19 19:37:38.256315911"
func Unixnanotime_str2iso(unixnanotime_str string) (iso string) {

	var unixnanotime int64 = 0

	tmp, err := strconv.ParseUint(unixnanotime_str, 10, 64)
	if err == nil {
		unixnanotime = int64(tmp)
	}

	return Unixnanotime2iso(unixnanotime)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// 1624120658256315 -> "2021-06-19 19:37:38.256315000"
func Unixmicrotime2iso(unixmicrotime int64) (iso string) {
	return Unixnanotime2iso(unixmicrotime * 1000)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// "1624120658256315" -> "2021-06-19 19:37:38.256315000"
func Unixmicrotime_str2iso(unixmicrotime_str string) (iso string) {

	var unixmicrotime int64 = 0

	tmp, err := strconv.ParseUint(unixmicrotime_str, 10, 64)
	if err == nil {
		unixmicrotime = int64(tmp)
	}

	return Unixnanotime2iso(unixmicrotime * 1000)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// 1624120658 -> "2021-06-19 19:37:38.000000000"
func Unixtime2iso(unixtime int64) (iso string) {
	return Unixnanotime2iso(unixtime * 1000000000)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// "1624120658" -> "2021-06-19 19:37:38.000000000"
func Unixtime_str2iso(unixtime_str string) (iso string) {

	var unixtime int64 = 0

	tmp, err := strconv.ParseUint(unixtime_str, 10, 64)
	if err == nil {
		unixtime = int64(tmp)
	}

	return Unixnanotime2iso(unixtime * 1000000000)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func isotime_test() bool {

	if Unixnanotime2iso(1624120658256315911) != "2021-06-19 19:37:38.256315911" {
		fmt.Printf("step001\n")
		return false
	}

	if Unixnanotime_str2iso("1624120658256315911") != "2021-06-19 19:37:38.256315911" {
		fmt.Printf("step002\n")
		return false
	}

	if Unixmicrotime2iso(1624120658256315) != "2021-06-19 19:37:38.256315000" {
		fmt.Printf("step003\n")
		return false
	}

	if Unixmicrotime_str2iso("1624120658256315") != "2021-06-19 19:37:38.256315000" {
		fmt.Printf("step004\n")
		return false
	}

	if Unixtime2iso(1624120658) != "2021-06-19 19:37:38.000000000" {
		fmt.Printf("step005\n")
		return false
	}

	if Unixtime_str2iso("1624120658") != "2021-06-19 19:37:38.000000000" {
		fmt.Printf("step006\n")
		return false
	}

	fmt.Printf("ok\n")
	return true
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func main() {
	isotime_test()
}
*/
