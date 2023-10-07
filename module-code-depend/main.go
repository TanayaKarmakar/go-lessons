package main

import (
	"fmt"

	"github.com/TanayaKarmakar/puppy"
)

func main() {
	str1 := puppy.Bark()

	str2 := puppy.Barks()

	fmt.Println(str1)

	fmt.Println(str2)
}
