package isvalidsudoku

func IsValidSudoku(board [][]byte) bool {
	rowfreq := make([]map[byte]bool, len(board))
	colfreq := make([]map[byte]bool, len(board[0]))
	var freq3x3 [3][3]map[byte]bool
	for i := 0; i < 9; i++ {
		rowfreq[i] = make(map[byte]bool, 9)
		colfreq[i] = make(map[byte]bool, 9)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			freq3x3[i][j] = make(map[byte]bool, 9)
		}
	}
	for y, row := range board {
		for x, cell := range row {
			if cell == '.' {
				continue
			}
			if rowfreq[y][cell] {
				return false
			}
			rowfreq[y][cell] = true
			if colfreq[x][cell] {
				return false
			}
			colfreq[x][cell] = true
			if freq3x3[y/3][x/3][cell] {
				return false
			}
			freq3x3[y/3][x/3][cell] = true
		}
	}
	return true
}

