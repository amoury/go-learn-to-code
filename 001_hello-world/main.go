package main

import "fmt"

func main() {
	// short declaration operator
	name := "Ansar"
	fmt.Println("Hello ", name)

	name = "Something else"
	fmt.Println("Hello ", name)

	// using var to declare variable
	var x bool // initializes to its 0 value
	fmt.Println(x)

	// String can also use back ticks to preserve multiline etc
	greeting := `Hello "You" 
This is the super awesome language GO
	`

	fmt.Println(greeting)

}
