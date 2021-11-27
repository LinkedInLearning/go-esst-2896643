package main

import (
	"fmt"
)

func main() {
	
	villes := [3]string{}
	villes[ 0 ] = "Rabat"
	villes[ 1 ] = "Paris"
	villes[ 2 ] = "Montréal"


	fmt.Println(villes)
	fmt.Println( "\nMise à jour des villes" )
	maj(villes)
	fmt.Println()
	fmt.Println(villes)
  	
}

func maj(villes [ 3 ] string ) {
	villes[ 0 ] = "Rabat"
	villes[ 1 ] = "Paris"
	villes[ 2 ] = "Madrid"

}