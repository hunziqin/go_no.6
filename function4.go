package main

import "fmt"

func main() {
	var i int
	var f float32
	i, _, f = ThreeValues()
	fmt.Printf("The int: %d, the float: %f \n", i, f)
}

func ThreeValues() (int, int, float32) {
	return 1, 2, 3.0
}
