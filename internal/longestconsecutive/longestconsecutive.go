package longestconsecutive

func LongestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))

	for _, num := range nums {
		set[num] = true
	}

	highest := 0
	for _, num := range nums {
		if !set[num - 1] {
			count := 1
			for set[num+count] {
				count += 1
			}
			if highest < count {
				highest = count
			}
		}
	}

	return highest
}

