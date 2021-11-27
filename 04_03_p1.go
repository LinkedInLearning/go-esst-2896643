package main

import (
	"fmt"
)


func main() {
	y := 10
	fmt.Println(point(&y))
	y = 11
	fmt.Println(point(&y))
}

func point(x *int) int {
	return *x * *x
}