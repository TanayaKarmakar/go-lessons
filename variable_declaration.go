package main

import "fmt"

func main() {
	a := 42 //short declaration operator
	fmt.Println(a)

	b, c, d, _, f := 2, 3, 4, 5, "Happiness"

	fmt.Println(b, c, d, f)

	var g int
	fmt.Printf("Before assignment %d\n", g)

	g = 100
	fmt.Printf("After assignment %d\n", g)
}
