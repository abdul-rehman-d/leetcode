package is_palindrome

import (
	"unicode"
)

func isNotAlphanumeric(char byte) bool {
	return !unicode.IsLetter(rune(char)) && !unicode.IsDigit(rune(char))
}

func isEqual(a, b byte) bool {
	aLower := unicode.ToLower(rune(a))
	bLower := unicode.ToLower(rune(b))

	return aLower == bLower
}

func IsPalindrome(s string) bool {
	j := len(s) - 1
	i := 0
	for i < j {
		if isNotAlphanumeric(s[i]) {
			i++
			continue
		}
		if isNotAlphanumeric(s[j]) {
			j--
			continue
		}
		if !isEqual(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
