package rules

func unquote(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	quote := s[0]
	if quote != s[n-1] {
		return ""
	}

	return s[1 : n-1]
}
