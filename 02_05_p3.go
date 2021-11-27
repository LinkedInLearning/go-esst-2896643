package main

import (
	"fmt"
)

func main() {
	i := 20
	dec(&i)

	fmt.Println("i=",i)
	fmt.Println("&i=",&i)
}

func dec(x *int){
	*x--
}