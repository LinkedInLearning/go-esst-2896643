package main

import (
	"fmt"
)

func main() {

	fmt.Println("Inline:", foncFonc(func(i int) int {return i * 5}, 10))
	fmt.Println("Inline:", foncFonc(func(i int) int {return i * i}, 10))
}

func foncFonc(f func(int) int, y int) int {
	return f(y)
}