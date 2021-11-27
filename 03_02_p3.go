package main

import (
	"fmt"
)

func main() {
	
	villes := map [ int ] string {
		1 : "Montréal" ,
		14 : "Paris",
	}
	
	for c := range villes {
		fmt.Println(c)
	}
}
