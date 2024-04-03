package isanagram

func IsAnagram(s string, t string) bool {
	mapp := make(map[rune]uint16, 26)

	for _, ch := range s {
		mapp[ch]++
	}
	for _, ch := range t {
		if mapp[ch] == 0 {
			return false
		}
		mapp[ch]--
	}
	for _, val := range mapp {
		if val != 0 {
			return false
		}
	}

	return true
}

