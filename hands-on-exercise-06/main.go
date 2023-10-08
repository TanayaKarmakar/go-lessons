package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("i = ", i)
	}

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("x and y are %v and %v\t\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("Both are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("both are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is 4 to 6 inclusive of both the numbers")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("None of the previous format")
	}

	for i := 0; i < 100; i++ {

		x = rand.Intn(10)
		y = rand.Intn(10)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both are less than 4")
		case x > 6 && y > 6:
			fmt.Println("both are greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is 4 to 6 inclusive of both the numbers")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("None of the previous format")
		}
	}

}
