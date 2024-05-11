package minimum_window_substring

import "fmt"

func MinWindow(s string, t string) string {
    if len(t) > len(s) {
        return ""
    }

	solutionStart := -1
	solutionEnd := -1

	tFreq := make(map[byte]int, 26)
	winFreq := make(map[byte]int, 26)

	need := 0
	have := 0

	for idx := range t {
        if _, has := tFreq[t[idx]]; has {
            tFreq[t[idx]]++
            continue
        }
        tFreq[t[idx]] = 1
        need++
	}

	fmt.Println(tFreq)

	left := 0
	right := 0
	winFreq[s[right]]++
	if winFreq[s[right]] == tFreq[s[right]] {
		have++
	}
	for left <= right && right < len(s) {
		fmt.Println(winFreq)
		fmt.Printf("r=%d\tl=%d\tr=%v\tl=%v\t%s\tneed=%d\thave=%d\n", right, left, s[right], s[left], s[left:right+1], need, have)

		if need == have {
			// solution save
			if right-left+1 < solutionEnd-solutionStart+1 || solutionStart == -1 || solutionEnd == -1 {
				solutionEnd = right
				solutionStart = left
			}

			if winFreq[s[left]] == tFreq[s[left]] {
				have--
			}
			winFreq[s[left]]--
			left++
		} else {
			right++
			if right == len(s) {
				break
			}
			winFreq[s[right]]++
			if winFreq[s[right]] == tFreq[s[right]] {
				have++
			}
		}
	}

	if solutionStart < 0 || solutionEnd < 0 {
		return ""
	}
	return s[solutionStart : solutionEnd+1]
}
