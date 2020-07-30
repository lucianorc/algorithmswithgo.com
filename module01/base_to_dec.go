package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	dict := "0123456789ABCDEF"

	var result int
	multiplier := 1

	getNum := func(ch rune) int {
		for index, c := range dict[0:base] {
			if c == ch {
				return index
			}
		}
		return 0
	}

	for _, char := range Reverse(value) {
		result += getNum(char) * multiplier
		multiplier *= base
	}

	return result
}
