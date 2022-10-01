package solution

func pacificAtlantic(heights [][]int) [][]int {
	const PAC, ATL = 1, 2
	rows, cols := len(heights), len(heights[0])
	visited, results := make(map[int]int), [][]int{}

	var dfs func(r, c int, prev int, ocean int)
	dfs = func(r, c int, prev int, ocean int) {
		height := heights[r][c]
		if height < prev {
			return
		}
		i := r*cols + c
		if v, has := visited[i]; !has || v > 0 && v != ocean {
			if !has {
				visited[i] = ocean
			} else {
				visited[i] = -1
				results = append(results, []int{r, c})
			}

			if r > 0 {
				dfs(r-1, c, height, ocean) //north
			}
			if c > 0 {
				dfs(r, c-1, height, ocean) //west
			}
			if r < rows-1 {
				dfs(r+1, c, height, ocean) //south
			}
			if c < cols-1 {
				dfs(r, c+1, height, ocean) //east
			}
		}
	}

	for c := 0; c < cols; c++ {
		dfs(0, c, heights[0][c], PAC)           //north edge
		dfs(rows-1, c, heights[rows-1][c], ATL) //south edge
	}
	for r := 0; r < rows; r++ {
		dfs(r, 0, heights[r][0], PAC)           //west edge
		dfs(r, cols-1, heights[r][cols-1], ATL) //east edge
	}

	return results
}
