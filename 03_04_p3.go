package main

import (
	"fmt"
)

func main() {
	fmt.Println(retValeurs(2))
	p1, p2 := retValeurs(3)
	fmt.Println(p1, p2)

	p1, p2 = p2, p1
	fmt.Println(p1, p2)

	x1, x2, x3, x4 := p2*p2, p2*p2*p2, p2+p2, p2*3

	fmt.Println(x1, x2, x3, x4)
}

func retValeurs(x int) (int, int) {
	return x * x, x * x * x
}