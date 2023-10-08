package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is the init function")
}

func main() {
	x := rand.Intn(400)
	fmt.Printf("Value of x is %v\n", x)

	if x <= 100 {
		fmt.Println("Less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("Between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("Between 201 and 250")
	} else {
		fmt.Println("This was more than 250")
	}

	x = rand.Intn(400)
	switch {
	case x <= 100:
		fmt.Println("Less than 100")
	case x >= 101 && x <= 200:
		fmt.Println("Between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("This was more than 250")
	}

	//y := rand.Intn(3)
	fmt.Println("y = ", rand.Intn(3))
	fmt.Println("y = ", rand.Intn(3))
	fmt.Println("y = ", rand.Intn(3))
	fmt.Println("y = ", rand.Intn(3))
	fmt.Println("y = ", rand.Intn(3))

}
