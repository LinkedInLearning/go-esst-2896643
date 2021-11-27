package main

import (
	"fmt"
)

func main() {
	c := 2
	for i := 0; i<3; i++ {
		func(a int, b int){
			fmt.Println( a+b)	
		}(i,c)	
	}
}