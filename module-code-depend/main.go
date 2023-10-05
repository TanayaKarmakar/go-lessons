package main

import (
	"fmt"
	"github.com/GoesToEleven/puppy"
)

func main() {
	str1 := puppy.bark()

	str2 := puppy.barks()

	fmt.Println(str1)

	fmt.Println(str2)
}
