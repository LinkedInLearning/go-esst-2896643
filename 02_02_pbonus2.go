package main

import (
	"fmt"
)

func main() {
	x_1 := 3 + 4i
	x_2 := complex(5, 6)
	x_3 := x_1 + x_2
	fmt.Printf("Type de x_1: %T\n", x_1)
	fmt.Printf("Type de x_2: %T\n", x_2)
	
	fmt.Println("Somme: x_1 +x_2 =", x_3)
}