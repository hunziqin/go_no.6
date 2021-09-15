package main

import "fmt"

func main() {
	fmt.Printf("M 2*5*6 = %d\n", M(2, 5, 6))
}
func M(a int, b int, c int) int {
	return a * b * c
}
