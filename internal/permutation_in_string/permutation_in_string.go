package permutation_in_string

func CheckInclusion(s1, s2 string) bool {
    var freq [26]int

    for _, v := range s1 {
        freq[v-'a']++
    }

    check := func(s string, freq [26]int) bool {
        for _, raw := range s {
            ch := raw-'a'
            if freq[ch] == 0 {
                return false
            }
            freq[ch]--
        }
        return true
    }
    
    for left := 0; left <= len(s2) - len(s1); left++ {
        right := left + len(s1) // exclusive
        if check(s2[left:right], freq) {
            return true
        }
    }
    return false
}
