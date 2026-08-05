package search_a_2d_matrix

func SearchA2dMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		panic("invalid input: 0 length array")
	}

	rows := len(matrix)
	cols := len(matrix[0])

	start := 0
	end := (rows * cols) - 1

	for start <= end {
		middleIdx := start + ((end - start) / 2) // flat
		ci := middleIdx % cols
		ri := middleIdx / cols

		middle := matrix[ri][ci]

		if middle == target {
			return true
		} else if middle > target {
			end = middleIdx - 1
		} else {
			start = middleIdx + 1
		}
	}

	return false
}
