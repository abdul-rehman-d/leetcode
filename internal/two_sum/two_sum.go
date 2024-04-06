package two_sum

func TwoSum(nums []int, target int) []int {
	mapp := make(map[int]int, len(nums))
	for idx, num := range nums {
		val, has := mapp[num]
		if has {
			return []int{val, idx}
		}
		mapp[target-num] = idx
	}
	return []int{}
}

