package solution

func combinationSum(candidates []int, target int) [][]int {
	results := [][]int{}

	var dfs func([]int, int, int)
	dfs = func(curr []int, i int, remaining int) {
		if i < len(candidates) {
			next := remaining - candidates[i]
			if next == 0 {
				result := make([]int, len(curr))
				copy(result, curr) //copy to new to prevent reslicing of results
				results = append(results, append(result, candidates[i]))
			} else if next > 0 {
				dfs(append(curr, candidates[i]), i, next)
			}
			dfs(curr, i+1, remaining)
		}
	}
	dfs([]int{}, 0, target)
	return results
}
