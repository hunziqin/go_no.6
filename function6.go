package main

import "fmt"

func Multiply(a int, b int, reply *int) {
	*reply = a * b
}

func main() {
	var n int
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply)
}
