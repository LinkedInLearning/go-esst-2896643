package main

import (
	"fmt"
)

func main() {
	add := func(a int, b int) int{
			return a + b
	}
	for i := 0; i<3; i++ {
		fmt.Println( "result : " , add( 10, i))
	}
}