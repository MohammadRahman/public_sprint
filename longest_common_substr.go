package sprint

func LongestCommonSubstr(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	if len1 == 0 || len2 == 0 {
		return ""
	}

	lcs := make([][]int, len1+1)
	for i := range lcs {
		lcs[i] = make([]int, len2+1)
	}

	longestLen := 0
	endPos := 0
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if str1[i-1] == str2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
				if lcs[i][j] > longestLen {
					longestLen = lcs[i][j]
					endPos = i
				}
			} else {
				lcs[i][j] = 0
			}
		}
	}

	if longestLen == 0 {
		return ""
	}

	startPos := endPos - longestLen
	return str1[startPos:endPos]
}
