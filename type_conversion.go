package main

import "fmt"

func main() {
	a := 42

	fmt.Printf("%v is of type %T\n", a, a)

	b := 45.78

	fmt.Printf("%v is of type %T\n", b, b)

	var c float32 = 67.8

	fmt.Printf("%v is of type %T\n", c, c)
}
