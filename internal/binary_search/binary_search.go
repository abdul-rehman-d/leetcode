package binary_search

func BinarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		middleIdx := start + ((end - start) / 2)
		middle := nums[middleIdx]

		if middle == target {
			return middleIdx
		} else if middle > target {
			end = middleIdx - 1
		} else {
			start = middleIdx + 1
		}
	}

	return -1
}
