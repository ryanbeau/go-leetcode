package solution

import "strconv"

func encode(s []string) string {
	var b []byte
	for _, s := range s {
		b = strconv.AppendInt(b, int64(len(s)), 10)
		b = append(b, ')')
		b = append(b, s...)
	}
	return string(b)
}

func decode(s string) []string {
	l := 0
	v := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			c, err := strconv.Atoi(s[l : l+i-l])
			if err != nil {
				return nil
			}
			l = i + c + 1
			v = append(v, s[i+1:l])
			i += c
		}
	}
	return v
}
