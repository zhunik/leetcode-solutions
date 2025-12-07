package solutions

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	if strs[0] == "" {
		return ""
	}
	p := strs[0]

	var l int = len(p)
	for i := 1; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}
		l = min(l, len(strs[i]))
		for l >= 0 && p[:l] != strs[i][:l] {
			l--
		}
		if l < 0 {
			return ""
		}
	}

	return p[:l]
}
