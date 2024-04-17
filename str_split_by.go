package sprint

func StrSplitBy(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	var substrings []string
	substring := ""

	for i := 0; i < len(s); i++ {
		match := true
		for j := 0; j < len(sep); j++ {
			if i+j >= len(s) || s[i+j] != sep[j] {
				match = false
				break
			}
		}

		if match {
			substrings = append(substrings, substring)
			substring = ""
			i += len(sep) - 1
		} else {
			substring += string(s[i])
		}
	}

	if len(substring) > 0 {
		substrings = append(substrings, substring)
	}

	return substrings
}
