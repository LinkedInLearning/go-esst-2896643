package main

import (
	"fmt"
)

func main() {
	
	villes := make ( map [ int ] string )
	fmt.Println( "Map vide: " , villes)
	villes[ 1 ] = "Montréal"
	villes[ 14 ] = "Paris"
	fmt.Println( "Map avec deux éléments : " , villes)
	fmt.Println( "Nombre d'éléments : " , len(villes))
}

