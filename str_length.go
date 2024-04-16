package sprint

func StrLength(s string) []int {
	runesCount := 0
	bytesCount := 0

	for i := 0; i < len(s); {
		if s[i]&0xC0 != 0x80 {
			runesCount++
		}
		bytesCount++
		i++
	}

	return []int{runesCount, bytesCount}
}
