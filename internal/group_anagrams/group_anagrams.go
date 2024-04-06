package group_anagrams

type counter [26]int

func GroupAnagrams(strs []string) [][]string {
	var a rune = 'a'

	mapp := make(map[counter]uint32)
	var out [][]string

	for _, str := range strs {
		key := counter{}
		for _, ch := range str {
			key[ch - a]++
		}
		if _, has := mapp[key]; has {
			out[mapp[key]] = append(out[mapp[key]], str)
		} else {
			mapp[key] = uint32(len(out))
			out = append(out, []string{str})
		}
	}

	return out
}

