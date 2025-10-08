package longest_repeating_character_replacement

import "slices"

func CharacterReplacement(s string, k int) int {
	l, r := 0, 0
	longest := 0
	maxx := 0

	freq := make([]int, 26, 26)

	for l <= r && r < len(s) {
		freq[s[r]-65]++

		length := r - l + 1

		maxx = max(maxx, freq[s[r]-65])

		ourK := length - maxx

		if ourK > k {
			freq[s[l]-65]--
			freq[s[r]-65]--
			l++
			maxx = slices.Max(freq)
		} else {
			longest = max(length, longest)
			r++
		}
	}

	return longest
}
