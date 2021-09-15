package main

import (
	"fmt"
	"math"
)

var a float64 = 2
var b float64

func main() {
	b = MySqrt(a)
	if b < 0 {
		fmt.Println("error")
	} else {
		fmt.Println(b)
	}
}

/*func MySqrt(input float64) float64 {
	c := math.Sqrt(float64(a))
	return c
}*/

func MySqrt(input float64) (x1 float64) {
	x1 = math.Sqrt(float64(a))
	return
}
