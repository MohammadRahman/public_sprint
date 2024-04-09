package main

import "fmt"

// data types in go
func Abacus(a int, b int) int {
	return a % b
}

func main() {
	c := Abacus(8, 3)

	fmt.Println(c)
}
