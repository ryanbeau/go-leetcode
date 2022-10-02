package solution

func canFinish(numCourses int, prerequisites [][]int) bool {
	coursePrereqs := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		c, p := prerequisites[i][0], prerequisites[i][1]
		coursePrereqs[c] = append(coursePrereqs[c], p)
	}

	visited := make(map[int]struct{}, numCourses)
	var dfs func(c int) bool
	dfs = func(c int) bool {
		if _, has := visited[c]; has {
			return false //infinite loop
		}
		if len(coursePrereqs[c]) == 0 {
			return true
		}

		visited[c] = struct{}{}
		for _, p := range coursePrereqs[c] {
			if !dfs(p) {
				return false
			}
		}
		delete(visited, c)
		coursePrereqs[c] = nil
		return true
	}

	for c := 0; c < numCourses; c++ {
		if !dfs(c) {
			return false
		}
	}
	return true
}
