package libcore

func IsParity(val uint64) bool {
	if (val & 1) == 0 {
		return true
	}
	return false
}