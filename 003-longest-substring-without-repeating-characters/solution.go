package solution

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	charPos := make(map[byte]int, 107) //chars+digits+symbols+space
	count, l := 1, 0
	charPos[s[l]] = l

	for r := 1; r < len(s); r++ {
		if pos, has := charPos[s[r]]; has && pos >= l {
			l = pos + 1
		} else if c := r - l + 1; count < c {
			count = c
		}
		charPos[s[r]] = r
	}
	return count
}
