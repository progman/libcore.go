//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"strconv"
	"fmt"
	"time"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
type Time struct {
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
	WeekDay    int
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetUnixnanotimeUtcInt64() (unixnanotimeUtc int64) {
	unixnanotimeUtc = time.Now().UTC().UnixNano()
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetUnixnanotimeUtc() (unixnanotimeUtc string) {
	unixnanotimeUtc = SintToStr(GetUnixnanotimeUtcInt64())
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetUnixmicrotimeUtcInt64() (unixmicrotimeUtc int64) {
	unixmicrotimeUtc = GetUnixnanotimeUtcInt64() / int64(time.Microsecond)
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetUnixmicrotimeUtc() (unixmicrotimeUtc string) {
	unixmicrotimeUtc = SintToStr(GetUnixmicrotimeUtcInt64())
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func UnixnanotimeUnpack(unixnanotime int64, pTime *Time) {
	t := time.Unix(0, unixnanotime)
	t = t.UTC()

	pTime.Year       = t.Year()
	pTime.Month      = int(t.Month())
	pTime.Day        = t.Day()
	pTime.Hour       = t.Hour()
	pTime.Minute     = t.Minute()
	pTime.Second     = t.Second()
	pTime.Nanosecond = t.Nanosecond()
	pTime.WeekDay    = int(t.Weekday())

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func UnixnanotimePack(pTime *Time) (unixnanotime int64) {

	t := time.Date(pTime.Year, time.Month(pTime.Month), pTime.Day, pTime.Hour, pTime.Minute, pTime.Second, pTime.Nanosecond, time.UTC)
	unixnanotime = t.UnixNano()

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// 1624120658256315911 -> "2021-06-19 16:37:38.256315911"
// Obsolete: func Unixnanotime2iso(unixnanotime int64) (iso string) {
func UnixnanotimeToIso(unixnanotime int64, offset int64, flagNanoseconds bool) (iso string) {
	var t1 Time

	UnixnanotimeUnpack(unixnanotime + offset, &t1)

	if flagNanoseconds == false {
		iso = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d",      t1.Year, t1.Month, t1.Day, t1.Hour, t1.Minute, t1.Second)
	} else {
		iso = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%09d", t1.Year, t1.Month, t1.Day, t1.Hour, t1.Minute, t1.Second, t1.Nanosecond)
	}

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// "1624120658256315911" -> "2021-06-19 16:37:38.256315911"
// Obsolete: func Unixnanotime_str2iso(unixnanotime_str string) (iso string) {
func UnixnanotimeStrToIso(unixnanotimeStr string, offset int64, flagNanoseconds bool) (iso string) {
	var err error
	var unixnanotime int64

	unixnanotime, err = StrToSint(unixnanotimeStr)
	if err != nil {
		unixnanotime = 0
	}

	return UnixnanotimeToIso(unixnanotime, offset, flagNanoseconds)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// 1624120658256315 -> "2021-06-19 16:37:38.256315000"
// Obsolete: func Unixmicrotime2iso(unixmicrotime int64) (iso string) {
func UnixmicrotimeToIso(unixmicrotime int64, offset int64, flagNanoseconds bool) (iso string) {
	return UnixnanotimeToIso(unixmicrotime * 1000, offset, flagNanoseconds)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// "1624120658256315" -> "2021-06-19 16:37:38.256315000"
// Obsolete: func Unixmicrotime_str2iso(unixmicrotime_str string) (iso string) {
func UnixmicrotimeStrToIso(unixmicrotimeStr string, offset int64, flagNanoseconds bool) (iso string) {
	var err error
	var unixmicrotime int64

	unixmicrotime, err = StrToSint(unixmicrotimeStr)
	if err != nil {
		unixmicrotime = 0
	}

	return UnixnanotimeToIso(unixmicrotime * 1000, offset, flagNanoseconds)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// 1624120658 -> "2021-06-19 16:37:38.000000000"
// Obsolete: func Unixtime2iso(unixtime int64) (iso string) {
func UnixtimeToIso(unixtime int64, offset int64, flagNanoseconds bool) (iso string) {
	return UnixnanotimeToIso(unixtime * 1000000000, offset, flagNanoseconds)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// "1624120658" -> "2021-06-19 16:37:38.000000000"
// Obsolete: func Unixtime_str2iso(unixtime_str string) (iso string) {
func UnixtimeStrToIso(unixtimeStr string, offset int64, flagNanoseconds bool) (iso string) {
	var err error
	var unixtime int64

	unixtime, err = StrToSint(unixtimeStr)
	if err != nil {
		unixtime = 0
	}

	return UnixnanotimeToIso(unixtime * 1000000000, offset, flagNanoseconds)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func TimeCropPerDay(source time.Time) (target time.Time) {

	var year int
	year = source.Year()

	var month int
	month = int(source.Month())

	var day int
	day = source.Day()

	target = time.Date(year, time.Month(month), day, 0, 0, 0, 0, source.Location()) // time.UTC

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func UnixnanotimeCropPerDay(unixnanotimeSource int64) (unixnanotimeTarget int64) {

	var t1 Time
	UnixnanotimeUnpack(unixnanotimeSource, &t1)

	t1.Hour       = 0
	t1.Minute     = 0
	t1.Second     = 0
	t1.Nanosecond = 0

	unixnanotimeTarget = UnixnanotimePack(&t1)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func isotime_test() bool {
	var t1 Time

	UnixnanotimeUnpack(1624120658256315911, &t1)
	if (t1.Year != 2021) || (t1.Month != 06) || (t1.Day != 19) || (t1.Hour != 16) || (t1.Minute != 37) || (t1.Second != 38) || (t1.Nanosecond != 256315911) {
		fmt.Printf("step001\n")
	}

	if UnixnanotimePack(&t1) != 1624120658256315911 {
		fmt.Printf("step002\n")
		return false
	}

	if UnixnanotimeCropPerDay(1624120658256315911) != 1624060800000000000 {
		fmt.Printf("step003\n")
		return false
	}

	if UnixnanotimeToIso(UnixnanotimeCropPerDay(1624120658256315911), 0, true) != "2021-06-19 00:00:00.000000000" {
		fmt.Printf("step004\n")
		return false
	}

	if UnixnanotimeToIso(1624120658256315911, 0, true) != "2021-06-19 16:37:38.256315911" {
		fmt.Printf("step005\n")
		return false
	}

	if UnixnanotimeStrToIso("1624120658256315911", 0, true) != "2021-06-19 16:37:38.256315911" {
		fmt.Printf("step006\n")
		return false
	}

	if UnixmicrotimeToIso(1624120658256315, 0, true) != "2021-06-19 16:37:38.256315000" {
		fmt.Printf("step007\n")
		return false
	}

	if UnixmicrotimeStrToIso("1624120658256315", 0, true) != "2021-06-19 16:37:38.256315000" {
		fmt.Printf("step008\n")
		return false
	}

	if UnixtimeToIso(1624120658, 0, true) != "2021-06-19 16:37:38.000000000" {
		fmt.Printf("step009\n")
		return false
	}

	if UnixtimeStrToIso("1624120658", 0, true) != "2021-06-19 16:37:38.000000000" {
		fmt.Printf("step010\n")
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
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
