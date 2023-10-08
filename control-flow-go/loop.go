package main

import "fmt"

func main() {
	fmt.Println("First loop")
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is ", i)
	}

	fmt.Println("Second loop")
	y := 0
	for y < 10 {
		fmt.Println("value of y is ", y)
		y++
	}

	fmt.Println("Third loop continue demo")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Current even number ", i)
	}

	fmt.Println("Nested loop demo")
	numRow := 3
	numCol := 3
	for currentRow := 0; currentRow < numRow; currentRow++ {
		for currentCol := 0; currentCol < numCol; currentCol++ {
			fmt.Printf("Row: %v, Col: %v|\t", currentRow, currentCol)
		}
		fmt.Println()
	}
}
