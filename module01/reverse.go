package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var reverse string

	for index := range word {
		position := len(word) - 1 - index
		reverse += string(word[position])
	}

	return reverse
}
