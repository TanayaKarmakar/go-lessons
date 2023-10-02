package main

import "fmt"

func main() {
	const name = "Tanaya"
	const subject = "Computer Science"
	fmt.Printf("%s's favourite subject %s\n", name, subject)

	fmt.Println(`The good way to learn any
	language is to read the "documentation". Trust me they
	are best source to know any technology`)

	a := 42

	fmt.Printf("Value of a in decimal %d\n", a)
	fmt.Printf("Value of a in binary %b\n", a)
	fmt.Printf("Value of a in hexadecimal %x\n", a)
}
