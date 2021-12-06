//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"errors"
	"strconv"
	"fmt"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func HexSignToDec(hexSign byte) (dec byte) {
	table := []byte{
//		0x00  0x01  0x02  0x03  0x04  0x05  0x06  0x07  0x08  0x09  0x0A  0x0B  0x0C  0x0D  0x0E  0x0F
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x00
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x10
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x20
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x30
		0xff, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x40
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x50
		0xff, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x60
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x70
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x80
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x90
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xA0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xB0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xC0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xD0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xE0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff} // 0xF0

	dec = table[hexSign]
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func DecSignToDec(decSign byte) (dec byte) {
	table := []byte{
//		0x00  0x01  0x02  0x03  0x04  0x05  0x06  0x07  0x08  0x09  0x0A  0x0B  0x0C  0x0D  0x0E  0x0F
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x00
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x10
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x20
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x30
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x40
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x50
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x60
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x70
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x80
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0x90
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xA0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xB0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xC0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xD0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 0xE0
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff} // 0xF0

	dec = table[decSign]
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func IsHexSign(val byte) (isHex bool) {
	dec := HexSignToDec(val)
	if dec != 0xff {
		isHex = true
	}
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func IsDecSign(val byte) (isDec bool) {
	dec := HexSignToDec(val)
	if dec != 0xff {
		isDec = true
	}
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func IsHex(str string) (isHex bool) {
	str_size := len(str)

	for i := 0; i < str_size; i++ {
		if IsHexSign(str[i]) == false {
			return
		}
	}

	isHex = true
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func IsDec(str string) (isDec bool) {
	str_size := len(str)

	for i := 0; i < str_size; i++ {
		if IsDecSign(str[i]) == false {
			return
		}
	}

	isDec = true
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func HexParity(str string) (string) {
	size := len(str)

	if IsParity(uint64(size)) == false {
		size++
		str = HexExpand(str, size)
	}

	return str
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func HexCrop(str string) (string) {
	str = HexParity(str)

	for {
		if len(str) <= 2 {
			break
		}


		if (str[0] == '0') && (str[1] == '0') {
			str = str[2:]
		} else {
			break
		}
	}

	return str
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func DecCrop(str string) (string) {

	for {
		if len(str) <= 1 {
			break
		}


		if str[0] == '0' {
			str = str[1:]
		} else {
			break
		}
	}

	return str
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func HexExpand(str string, size int) (string) {
	if len(str) > size {
		return str
	}

	for {
		if len(str) == size {
			break
		}
		str = "0" + str
	}

	return str
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func DecExpand(str string, size int) (string) {
	return HexExpand(str, size)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func SetByteInString(str string, index int, val byte) (string) {
	tmp := []byte(str)
	tmp[index] = val
	return string(tmp)
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func HexAdd(source string, add string) (target string, err error) {
	if IsHex(source) == false {
		err = errors.New("first argument is not hex")
		return
	}
	if IsHex(add) == false {
		err = errors.New("second argument is not hex")
		return
	}


	source = HexCrop(source)
	add = HexCrop(add)


	source_size := len(source)
	add_size    := len(add)

	size := source_size
	if size < add_size {
		size = add_size
	}


	source = HexExpand(source, size)
	add    = HexExpand(add, size)


	var a byte
	var b byte
	var x byte
	var carry byte
	for i := 0; i < size; i++ {
		x = HexSignToDec(source[size - i - 1]) + HexSignToDec(add[size - i - 1])

		x += carry
		carry = 0

		if (x > 15) {
			a = x / 16
			b = x - (a * 16)

			carry = a
			x = b
		}

		source = SetByteInString(source, size - i - 1, byte(strconv.FormatInt(int64(x), 16)[0]))

		if (((i + 1) == size) && (carry != 0)) {
			size++
			source = HexExpand(source, size)
			add    = HexExpand(add,    size)
		}
	}


	source = HexParity(source)
	target = source


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func DecAdd(source string, add string) (target string, err error) {
	if IsDec(source) == false {
		err = errors.New("first argument is not dec")
		return
	}
	if IsDec(add) == false {
		err = errors.New("second argument is not dec")
		return
	}


	source = DecCrop(source)
	add = DecCrop(add)


	source_size := len(source)
	add_size    := len(add)

	size := source_size
	if size < add_size {
		size = add_size
	}


	source = DecExpand(source, size)
	add    = DecExpand(add, size)


	var a byte
	var b byte
	var x byte
	var carry byte
	for i := 0; i < size; i++ {
		x = DecSignToDec(source[size - i - 1]) + DecSignToDec(add[size - i - 1])

		x += carry
		carry = 0

		if (x > 9) {
			a = x / 10
			b = x - (a * 10)

			carry = a
			x = b
		}

		source = SetByteInString(source, size - i - 1, byte(strconv.FormatInt(int64(x), 10)[0]))

		if (((i + 1) == size) && (carry != 0)) {
			size++
			source = DecExpand(source, size)
			add    = DecExpand(add,    size)
		}
	}


	target = source


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
       HexCmp() returns an integer indicating the result of the comparison, as follows:

       • 0, if the s1 and s2 are equal;

       • a negative value if s1 is less than s2;

       • a positive value if s1 is greater than s2.
*/
func HexCmp(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}

	if len(s1) < len(s2) {
		return -1
	}

	if len(s1) > len(s2) {
		return 1
	}

	for i := 0; i < len(s1); i++ {
		a := HexSignToDec(s1[i])
		b := HexSignToDec(s2[i])

		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
	}

	return 0 // paranoid mode
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
       DecCmp() returns an integer indicating the result of the comparison, as follows:

       • 0, if the s1 and s2 are equal;

       • a negative value if s1 is less than s2;

       • a positive value if s1 is greater than s2.
*/
func DecCmp(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}

	if len(s1) < len(s2) {
		return -1
	}

	if len(s1) > len(s2) {
		return 1
	}

	for i := 0; i < len(s1); i++ {
		a := DecSignToDec(s1[i])
		b := DecSignToDec(s2[i])

		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
	}

	return 0 // paranoid mode
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func HexTest() (err error) {

	var x string
	var y string
	var z string


	x = "1"
	y = "2"
	z, err = HexAdd(x, y)
	if err != nil {
		return
	}
	if z != "03" {
		err = errors.New("HexText() step001")
		return
	}


	x = "fe"
	y = "1"
	z, err = HexAdd(x, y)
	if err != nil {
		return
	}
	if z != "ff" {
		err = errors.New("HexText() step002")
		return
	}


	x = "ff"
	y = "1"
	z, err = HexAdd(x, y)
	if err != nil {
		return
	}
	if z != "0100" {
		err = errors.New("HexText() step003")
		return
	}


	x = "ff"
	y = "ff"
	z, err = HexAdd(x, y)
	if err != nil {
		return
	}
	if z != "01fe" {
		err = errors.New("HexText() step004")
		return
	}


	x = "d5104dc76695721d"
	y = "b80704bb7b4d7c03"
	z, err = HexAdd(x, y)
	if err != nil {
		return
	}
	if z != "018d175282e1e2ee20" {
		err = errors.New("HexText() step005")
		return
	}


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func DecTest() (err error) {

	var x string
	var y string
	var z string


	x = "01"
	y = "02"
	z, err = DecAdd(x, y)
	if err != nil {
		return
	}
	if z != "3" {
		err = errors.New(fmt.Sprintf("DecTest() step001, z:%s", z))
		return
	}


	x = "8"
	y = "1"
	z, err = DecAdd(x, y)
	if err != nil {
		return
	}
	if z != "9" {
		err = errors.New(fmt.Sprintf("DecTest() step002, z:%s", z))
		return
	}


	x = "9"
	y = "1"
	z, err = DecAdd(x, y)
	if err != nil {
		return
	}
	if z != "10" {
		err = errors.New(fmt.Sprintf("DecTest() step003, z:%s", z))
		return
	}


	x = "99"
	y = "99"
	z, err = DecAdd(x, y)
	if err != nil {
		return
	}
	if z != "198" {
		err = errors.New(fmt.Sprintf("DecTest() step004, z:%s", z))
		return
	}


	x = "1239223372036854775807"
	y = "1239223372036854775807"
	z, err = DecAdd(x, y)
	if err != nil {
		return
	}
	if z != "2478446744073709551614" {
		err = errors.New(fmt.Sprintf("DecTest() step005, z:%s", z))
		return
	}


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func main() {
	var err error

	err = HexTest()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("ok\n")
	}


	err = DecTest()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("ok\n")
	}
}
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
