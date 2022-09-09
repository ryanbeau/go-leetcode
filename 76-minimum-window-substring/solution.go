package solution

func minWindow(s string, t string) string {
	const TCOUNT, WCOUNT = 0, 1

	count := make(map[byte][2]int, len(t))
	for i := 0; i < len(t); i++ {
		if v, ok := count[t[i]]; ok {
			v[TCOUNT]++
		} else {
			count[t[i]] = [2]int{1, 0}
		}
	}

	total, need := 0, len(t)
	resultL, resultR := -1, -1

	for l, r := 0, 0; r < len(s); r++ {
		if v, ok := count[s[r]]; ok {
			v[WCOUNT]++
			if v[WCOUNT] <= v[TCOUNT] {
				total++
			}
		}

		for total == need {
			if resultR == -1 || r-l < resultR-resultL {
				resultL = l
				resultR = r
			}
			if v, ok := count[s[l]]; ok {
				v[WCOUNT]--
				if v[WCOUNT] < v[TCOUNT] {
					total--
				}
			}
			l++
		}
	}

	if resultR >= 0 {
		return s[resultL : resultR+1]
	}
	return ""
}
