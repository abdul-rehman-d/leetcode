package sliding_window_maximum

func MaxSlidingWindow(nums []int, k int) []int {
	out := make([]int, 0, len(nums)-k+1)

	queue := make([]int, 0, len(nums))

	l, r := 0, 0

	for l <= r && r < len(nums) {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[r] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, r)

		if l > queue[0] {
			queue = queue[1:]
		}

		if r-l+1 == k {
			out = append(out, nums[queue[0]])
			l++
		}
		r++
	}

	return out
}
