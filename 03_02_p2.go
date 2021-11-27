package main

import (
	"fmt"
)

func main() {
	
	villes := map [ int ] string {
		1 : "Montréal" ,
		14 : "Paris",
	}
	
	fmt.Println( "Map avec deux éléments : " , villes)
}