package main

import "fmt"

func main() {
	MakeAddSuffix()
}

func MakeAddSuffix() {
	var file string = ".bmp"
	addBmp := func(file string) { fmt.Printf("%s", file) }
	addBmp(file)
	/*addJpeg := func(file string) { fmt.Printf("%s", file) }
	addJpeg(file)*/
}

//第六章练习6.10未完全版
