package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will repeat until break or
	// a return from the enclosing function
	for {
		fmt.Println("break")
		break
	}
}
