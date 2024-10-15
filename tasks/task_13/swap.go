package main

import "fmt"

func main() {
	var (
		a = 1
		b = 2
	)

	fmt.Printf("a = %d, b = %d\n", a, b)

	a, b = b, a // || а = b, b = a

	fmt.Printf("a = %d, b = %d\n", a, b)
}
