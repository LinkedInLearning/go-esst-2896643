package main

import (
	"fmt"
)

func main() {
	i := 1
	dec(i)

	fmt.Println("i=",i)
}

func dec(x int){
	x--
}