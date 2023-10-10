package main

import "fmt"

func main() {
	i := 20

	for i >= 0 {
		fmt.Println(i)
		i--
	}

	//for loop with a break
	x := 20
	for {
		if x < 0 {
			break
		}
		fmt.Println("x = ", x)
		x--
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer loop %v\t Inner loop %v\n", i, j)
		}
		fmt.Println()
	}
}
