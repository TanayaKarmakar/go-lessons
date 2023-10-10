package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("index = %v\t value = %v\n", i, v)
	}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("key: %v\t value: %v\n", k, v)
	}

	// extracting value of particular key
	age := m["James"]
	fmt.Printf("Age of James is %v\n", age)

	for i := 0; i < 10; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is 3")
		}
	}

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
