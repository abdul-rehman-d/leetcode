package longest_substring_without_repeating_characters

func LengthOfLongestSubstring(s string) int {
    l := 0
    longest := 0

    mapp := make(map[byte]bool)

    for r := 0; r < len(s); r++ {
        for {
            if mapp[s[r]] {
                delete(mapp, s[l])
                l++
            } else {
                break
            }
        }
        mapp[s[r]] = true
        longest = max(longest, r - l + 1)
    }

    return longest
}

