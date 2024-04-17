package sprint

func StrConcatWith(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	var result []byte
	result = appendToString(result, strs[0])

	for i := 1; i < len(strs); i++ {
		result = appendToString(result, sep)     // Append separator
		result = appendToString(result, strs[i]) // Append the string
	}

	return string(result)
}
func appendToString(dest []byte, src string) []byte {
	for i := 0; i < len(src); i++ {
		dest = append(dest, src[i])
	}
	return dest
}
