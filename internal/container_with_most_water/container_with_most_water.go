package container_with_most_water

func MaxArea(height []int) int {
	s := 0
	e := len(height) - 1
	maxArea := 0

	for s < e {
		w := e - s
		h := min(height[s], height[e])

		maxArea = max(w*h, maxArea)

		if height[s] > height[e] {
			e--
		} else {
			s++
		}

	}

	return maxArea
}
