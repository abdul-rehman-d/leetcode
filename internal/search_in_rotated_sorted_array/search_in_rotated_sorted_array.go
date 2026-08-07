package search_in_rotated_sorted_array

func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		panic("invalid input: empty array")
	}

	s := 0             // startIdx
	e := len(nums) - 1 // endIdx

	for s <= e {
		m := s + ((e - s) / 2) // middleIdx

		if nums[m] == target {
			return m
		}

		if target > nums[m] {
			if target <= nums[e] || nums[m] > nums[e] {
				// find in second half
				s = m + 1
			} else {
				// find in first half
				e = m - 1
			}
		} else {
			if target >= nums[s] || nums[s] > nums[m] {
				// find in first half
				e = m - 1
			} else {
				// find in second half
				s = m + 1
			}
		}
	}

	return -1
}
