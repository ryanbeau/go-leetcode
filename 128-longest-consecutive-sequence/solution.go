package solution

func longestConsecutive(nums []int) int {
	longest := 0

	m := make(map[int]int, len(nums))
	for _, n := range nums {
		if _, ok := m[n]; ok {
			continue //skip if duplicate
		}

		leftCount, leftOk := m[n-1]
		rightCount, rightOk := m[n+1]
		seq := leftCount + rightCount + 1

		if leftOk {
			m[n-leftCount] = seq
		}
		if rightOk {
			m[n+rightCount] = seq
		}
		m[n] = seq

		if seq > longest {
			longest = seq
		}
	}
	return longest
}
