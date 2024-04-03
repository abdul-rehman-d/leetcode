package containsduplicate

func ContainsDuplicate(nums []int) bool {
	mapp := make(map[int]bool, len(nums))
	for _, num := range nums {
		_, has := mapp[num]
		if has {
			return true
		}
		mapp[num] = true
	}
	return false
}

