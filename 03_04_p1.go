package main

import (
	"fmt"
)

func main() {
	fmt.Println(retValeurs(2))
	p1, p2 := retValeurs(3)
	fmt.Println(p1, p2)

}

func retValeurs(x int) (int, int) {
	return x * x, x * x * x
}