package solution

func minWindow(s string, t string) string {
	const TCOUNT, WCOUNT = 0, 1

	m := make(map[byte]*[2]int, len(t))
	for i := 0; i < len(t); i++ {
		if v, ok := m[t[i]]; ok {
			v[TCOUNT]++
		} else {
			m[t[i]] = &[2]int{1, 0}
		}
	}

	total := 0
	resultL, resultR := -1, -1

	for l, r := 0, 0; r < len(s) && l <= len(s)-len(t); {
		if total < len(t) {
			if v, ok := m[s[r]]; ok {
				v[WCOUNT]++
				if v[WCOUNT] <= v[TCOUNT] {
					total++
				}
			}
		} else {
			if v, ok := m[s[l]]; ok {
				v[WCOUNT]--
				if v[WCOUNT] < v[TCOUNT] {
					total--
				}
			}
			l++
		}

		if total < len(t) {
			r++
		} else if resultR == -1 || r-l < resultR-resultL {
			resultL = l
			resultR = r
		}
	}

	if resultR >= 0 {
		return s[resultL : resultR+1]
	}
	return ""
}
