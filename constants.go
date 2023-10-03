package main

import "fmt"

var a int

func main() {
	const c1 = iota
	fmt.Println(c1)

	a := 42
	fmt.Printf("decimal value of a is %d\n", a)
	fmt.Printf("binary value of a is %b\n", a)
	fmt.Printf("Hexadecimal value of a is %x\n", a)
}
