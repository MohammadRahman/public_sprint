package sprint

import "fmt"

func Abacus(a int, b int) int {

	if b == 0 {
		return 0
	}
	c := a / b
	return c
}

func main() {

	c := Abacus(4, 8)
	fmt.Println("c", c)
}
