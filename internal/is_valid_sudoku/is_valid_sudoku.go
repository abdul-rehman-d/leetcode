package is_valid_sudoku

func IsValidSudoku(board [][]byte) bool {
	var rowfreq [9][9]bool
	var colfreq [9][9]bool
	var freq3x3 [9][9]bool

	for y, row := range board {
		for x, cell := range row {
			if cell == '.' {
				continue
			}

			cell = cell - '0' - 1

			// row check
			if rowfreq[y][cell] {
				return false
			}

			// col check
			rowfreq[y][cell] = true
			if colfreq[x][cell] {
				return false
			}
			colfreq[x][cell] = true

			// block check
			key := (y/3 * 3) + x/3
			if freq3x3[key][cell] {
				return false
			}
			freq3x3[key][cell] = true
		}
	}
	return true
}

