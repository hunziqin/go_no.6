package main

import (
	"fmt"
	"reflect"
)

func main() {
	var h string = "Hello world"
	fv := func(h string) { fmt.Printf("Hello world") }
	fv(h)
	fmt.Println(reflect.TypeOf(fv))
}
