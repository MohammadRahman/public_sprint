package sprint

func Abacus(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}
