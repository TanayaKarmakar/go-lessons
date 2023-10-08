package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//if else demo
	x := 42
	if x < 42 {
		fmt.Println("x is less than meaning of life")
	} else {
		fmt.Println("x is equal to or more than meaning of life")
	}

	y := 50
	if y < 50 {
		fmt.Println("y is less than 50")
	} else if y == 50 {
		fmt.Println("y is equal to 50")
	} else {
		fmt.Println("y is greater than 50")
	}

	a := 10
	b := 20
	if a >= 10 && b < 20 {
		fmt.Println("a is greater than or equal to 10 and b is less than 20")
	} else if a >= 10 && b >= 20 {
		fmt.Println("a is greater than or equal to 10 and b is greater than or equal to 20")
	} else if a < 10 && b < 20 {
		fmt.Println("a is less than 10 and b is less than 20")
	} else {
		fmt.Println("Donno whats happening")
	}

	c := rand.Int()

	fmt.Printf("Value of c is %v\n", c)

	// switch case demo
	z := 50
	switch z {
	case 10:
		fmt.Println("Value is 10")
		//fallthrough
	case 20:
		fmt.Println("Valus is 20")
		//fallthrough
	case 30:
		fmt.Println("Valus is 30")
	default:
		fmt.Println("Value is greater than 30")
	}
}
