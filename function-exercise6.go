package main

import (
	"fmt"
)

func main() {
	result := 0
	for i := 0; i <= 12; i++ {
		result = Rfunction(i)
		fmt.Printf("%d factorial is %d\n", i, result)
	}
}

func Rfunction(n int) (res int) {
	if n == 0 {
		res = 1
	} else {
		res = n * Rfunction(n-1)
	}
	return
}
