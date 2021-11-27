package main

import (
	"fmt"
)

func main() {
	
	villes := []string{}
	villes = append (villes, "Rabat" )
	villes = append (villes, "Paris" )
	villes = append (villes, "Montréal")


	fmt.Println(villes)
	fmt.Println( "\nMise à jour des villes" )
	maj(villes)
	fmt.Println()
	fmt.Println(villes)
  	
}

func maj(villes [ ] string ) {
	villes[ 0 ] = "Rabat"
	villes[ 1 ] = "Paris"
	villes[ 2 ] = "Madrid"

}