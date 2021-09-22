package main

import "fmt"

func fib() func() int {
	var a int = 0
	var b int = 1
	return func() int {
		c := a    //c=0
		a = b     //a=1
		b = a + c //b=1
		return c  //c=0
	}
}

func main() {
	f := fib()
	for i := 0; i < 40; i++ {
		fmt.Println(f())
	}
}
