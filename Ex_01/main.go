package main

import "fmt"

type ansar int

var x ansar = 25

var y = int(x)

func main() {
	var x2 ansar

	fmt.Println(x2)
	fmt.Printf("%T\n", x2)

	x2 = 42
	fmt.Println(x2)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
