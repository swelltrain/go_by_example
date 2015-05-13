package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	// you can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// the type can be inferred from the value
	var d = true
	fmt.Println(d)

	// variables without an initialization are 0 valued
	var e int
	fmt.Println(e)

	//  := is shorthand for var f string = "short"
	// only available in a function
	f := "short"
	fmt.Println(f)
}
