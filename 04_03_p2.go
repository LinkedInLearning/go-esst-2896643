package main

import (
	"fmt"
)


func main() {
	doub := retDoublep(7)
	fmt.Println("Valeur double:", *doub)
	fmt.Println("Adresse mémoire:", doub)
}

func retDoublep(x int) *int {
	y := x * 2
	return &y
}