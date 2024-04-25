package sprint

func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func LCM(a, b int) int {
	product := a * (b / GCD(a, b))
	return product
}
