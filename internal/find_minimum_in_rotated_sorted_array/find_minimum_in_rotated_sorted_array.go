package find_minimum_in_rotated_sorted_array

func FindMinimumInRotatedSortedArray(nums []int) int {
	if len(nums) == 0 {
		panic("invalid input: empty array")
	}
	if len(nums) == 1 {
		return nums[0]
	}

	start := 0
	end := len(nums) - 1

	for start < end {
		middleIdx := start + ((end - start) / 2)
		middle := nums[middleIdx]
		// [1, 2, 3, 4, 5, 6] // perfect sort
		// [3, 4, 5, 6, 1, 2] // find in second half
		// [6, 1, 2, 3, 4, 5] // find in first half
		if nums[start] <= middle && middle < nums[end] {
			// perfect sort
			return nums[start]
		} else if nums[start] <= middle {
			// find in second half
			start = middleIdx + 1
		} else {
			// find in first half
			end = middleIdx
		}
	}

	return nums[start]
}
