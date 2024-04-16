package sprint

func StrReverse(s string) string {
	strBytes := []byte(s)
	length := len(strBytes)
	for i := 0; i < length/2; i++ {
		strBytes[i], strBytes[length-i-1] = strBytes[length-i-1], strBytes[i]
	}
	return string(strBytes)
}
