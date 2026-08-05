package largest_rectangle_in_histogram

func LargestRectangleInHistogram(heights []int) int {
	type A struct {
		Index  int
		Height int
	}

	if len(heights) == 0 {
		panic("invalid input: array length of 0")
	}
	if len(heights) == 1 {
		return heights[0]
	}

	stack := []A{
		{0, heights[0]},
	}
	m := 0
	l := len(heights)

	for i, a := range heights {
		if i == 0 {
			continue
		}
		idx := i
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if a < top.Height {
				width := i - top.Index
				area := width * top.Height
				m = max(area, m)

				idx = top.Index

				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, A{idx, a})
	}

	for _, s := range stack {
		width := l - s.Index
		area := width * s.Height
		m = max(area, m)
	}

	return m
}
