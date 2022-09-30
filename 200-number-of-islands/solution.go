package solution

func numIslands(grid [][]byte) int {
	rows, cols, islands := len(grid), len(grid[0]), 0
	visited := make([]bool, rows*cols)

	var visit func(int, int) bool
	visit = func(r, c int) bool {
		i := r*cols + c
		if !visited[i] {
			visited[i] = true

			if r < rows-1 && grid[r+1][c] == '1' { //down
				visit(r+1, c)
			}
			if c < cols-1 && grid[r][c+1] == '1' { //right
				visit(r, c+1)
			}
			if r > 0 && grid[r-1][c] == '1' { //up
				visit(r-1, c)
			}
			if c > 0 && grid[r][c-1] == '1' { //left
				visit(r, c-1)
			}
			return true
		}
		return false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' && visit(r, c) {
				islands++
			}
		}
	}
	return islands
}
