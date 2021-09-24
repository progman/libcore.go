//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import "errors"
import "strconv"
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func HexSignToDec(hex byte) (dec byte) {
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

	dec = table[hex]
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func IsHexSign(hex byte) (isHex bool) {
	dec := HexSignToDec(hex)
	if dec != 0xff {
		isHex = true
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
