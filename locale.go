package word

// CheckLocale ...
func CheckLocale(s string) bool {
	return true
}

func isKorean(r rune) bool {
	return r >= 0xAC00 && r <= 0xD7A3
}
func isJapanese(r rune) bool {
	// Chinese
	if (r >= 0x4E00 && r <= 0x9FC3) ||
		(r >= 0x3001 && r <= 0x303D) ||
		(r >= 0x3220 && r <= 0x3240) ||
		(r >= 0x3250 && r <= 0x325F) ||
		(r >= 0x3280 && r <= 0x32FE) ||
		(r >= 0x3300 && r <= 0x33FF) ||
		(r >= 0x3400 && r <= 0x4DB5) ||
		(r >= 0xF900 && r <= 0xFAD9) {
		return true
	}

	// Japanese
	if (r >= 0x3040 && r <= 0x309F) ||
		(r >= 0x30A0 && r <= 0x30FF) ||
		(r >= 0x31F0 && r <= 0x31FF) ||
		(r >= 0x00A1 && r <= 0x00FF) ||
		(r >= 0x1D00 && r <= 0x1D7F) ||
		(r >= 0x2010 && r <= 0x205E) ||
		(r >= 0x2070 && r <= 0x20B5) ||
		(r >= 0x2100 && r <= 0x22FF) ||
		(r >= 0xFE50 && r <= 0xFE6F) ||
		(r >= 0xFF01 && r <= 0xFF9F) {
		return true
	}
	return false
}
func isChinese(r rune) bool {
	return r >= 0x4E00 && r <= 0x9FAF
}
func isThai(r rune) bool {
	if r >= 0x0E01 && r <= 0x0E7F {
		if r != 0x0E3A &&
			r != 0x0E4D &&
			r != 0x0E4E &&
			r != 0x0E4F &&
			r != 0x0E5A &&
			r != 0x0E5B {
			return true
		}
	}
	return false
}
func isAlphabet(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}
func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
