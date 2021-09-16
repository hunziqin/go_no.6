package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	function(slice...)
}

func function(slice ...int) {
	fmt.Println(slice)
}
