package permutation_in_string


func CheckInclusion(s1, s2 string) bool {
    var s1freq [26]int
    var s2freq [26]int
    var matches int

    if len(s1) > len(s2) {
        return false
    }

    for i := 0; i < len(s1); i++ {
        s1freq[s1[i]-'a']++
        s2freq[s2[i]-'a']++
    }

    for i := 0; i < 26; i++ {
        if s1freq[i] == s2freq[i] {
            matches++
        }
    }
    
    left := 0
    for right := len(s1); right < len(s2); right++ {
        leftCh := s2[left]-'a'
        rightCh := s2[right]-'a'

        if matches == 26 {
            return true
        }

        s2freq[rightCh]++
        if s1freq[rightCh] == s2freq[rightCh] {
            matches++
        } else if s1freq[rightCh] == s2freq[rightCh]-1 {
            matches--
        }

        s2freq[leftCh]--
        if s1freq[leftCh] == s2freq[leftCh] {
            matches++
        } else if s1freq[leftCh] == s2freq[leftCh]+1 {
            matches--
        }

        left++
    }
    return matches == 26
}
