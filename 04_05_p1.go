package main

import (
	"fmt"
)

func main() {
	fmt.Println("function1:", foncFonc(fonc_1, 10))
	fmt.Println("function2:", foncFonc(fonc_2, 12))
}

func fonc_1(x int) int {
	return x * 2
}

func fonc_2(x int) int {
	return x * 3
}

func foncFonc(f func(int) int, y int) int {
	return f(y)
}