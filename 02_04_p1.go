package main

import (
	"fmt"
)

func main() {	
	i:=0
	for i <10 {
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++
	}
}