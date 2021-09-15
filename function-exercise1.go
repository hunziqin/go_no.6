package main

import "fmt"

var a, b int = 1, 2
var c, d, e, f int

func main() {
	c, d, e = SAP(f)
	fmt.Printf("a+b=%d,a*b=%d,a-b=%d\n", c, d, e)
}

/*func SAP(input int) (int, int, int) {
	return a + b, a * b, a - b
}*/

func SAP(input int) (x1 int, x2 int, x3 int) {
	x1 = a + b
	x2 = a * b
	x3 = a - b
	return
}
